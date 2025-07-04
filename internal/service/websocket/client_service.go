package websocket

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/coder/websocket"
	"log"
	"sync"
)

type Client struct {
	conn   *websocket.Conn
	roomId int64
	userId int64
	send   chan string
}

// Tao moi 1 client
func NewClient(conn *websocket.Conn, roomId int64, userId int64) *Client {
	return &Client{
		conn:   conn,
		roomId: roomId,
		userId: userId,
		send:   make(chan string, 256), //buffer 256 message
	}
}

func (c *Client) ReadLoop(ctx context.Context, db *sql.DB, hub *Hub) {
	room := hub.GetRoom(c.roomId)
	defer func() {
		log.Printf("Client %d disconnected", c.userId)
		room.Leave(c.userId)
		c.conn.Close(websocket.StatusNormalClosure, "read finished")
	}()

	for {
		_, msg, err := c.conn.Read(ctx)
		if err != nil {
			log.Println("read error:", err)
			break
		}

		var incoming dto.MessageDTO
		if err := json.Unmarshal(msg, &incoming); err != nil {
			log.Println("invalid message:", err)
			continue
		}

		incoming.SenderId = c.userId

		room.Broadcast(hub, db, ctx, incoming, c)
	}
}

func (c *Client) WriteLoop(ctx context.Context) {
	defer func() {
		log.Printf("Client %d write loop ended", c.userId)
		c.conn.Close(websocket.StatusNormalClosure, "write finished")
	}()

	for {
		select {
		case msg := <-c.send:
			err := c.conn.Write(ctx, websocket.MessageText, []byte(msg))
			if err != nil {
				log.Println("write error:", err)
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

type Room struct {
	Info      dto.RoomDTO
	Clients   map[int64]*Client
	mutex     sync.Mutex
	Persisted bool
}

func (r *Room) Join(c *Client) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.Clients[c.userId] = c
}

func (r *Room) Broadcast(hub *Hub, db *sql.DB, ctx context.Context, msg dto.MessageDTO, sender *Client) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.Persisted {
		repo := repository.NewRoomRepository(db)
		roomService := service.NewRoomService(repo)

		if msg.ReceiverId != 0 {
			r.Info.UserIds = append(r.Info.UserIds, msg.ReceiverId)
		}
		roomId, err := roomService.CreateRoom(ctx, r.Info, sender.userId)
		if err != nil {
			log.Println("failed to create room:", err)
			return
		}
		roomDTO, errGet := roomService.GetRoomById(ctx, roomId)
		if errGet != nil {
			log.Println("failed to find room:", err)
			return
		}
		r.Info = roomDTO
		r.Persisted = true

		//Update hub
		hub.mutex.Lock()
		hub.rooms[roomDTO.Id] = r
		delete(hub.rooms, 0)
		hub.mutex.Unlock()
	}

	repo := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(repo)

	createMsg := dto.MapMessageToCreateDTO(msg)
	_ = messageService.SendMessage(ctx, sender.userId, createMsg)

	data, err := json.Marshal(msg)
	if err != nil {
		log.Println("failed to encode message:", err)
		return
	}

	for userId, client := range r.Clients {
		if userId != sender.userId {
			client.send <- string(data)
		}
	}
}

func (r *Room) GetOnlineUserIDs() []int64 {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	userIDs := make([]int64, 0, len(r.Clients))
	for uid := range r.Clients {
		userIDs = append(userIDs, uid)
	}
	return userIDs
}

func (r *Room) Leave(clientID int64) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.Clients, clientID)
}
