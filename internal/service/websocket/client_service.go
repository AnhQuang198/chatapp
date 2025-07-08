package websocket

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/internal/service"
	"chatapp/internal/service/websocket/data"
	"chatapp/internal/service/websocket/enum"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coder/websocket"
	"log"
	"sync"
	"time"
)

const (
	readTimeout  = 5 * time.Minute
	pingInterval = 2 * time.Minute
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
		readCtx, cancel := context.WithTimeout(ctx, readTimeout)
		_, msg, err := c.conn.Read(readCtx)
		cancel()

		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				log.Printf("Client %d closed connection", c.userId)
			} else if errors.Is(err, context.DeadlineExceeded) {
				log.Printf("Client %d timeout", c.userId)
			} else {
				log.Printf("Read error from client %d: %v", c.userId, err)
			}
			break
		}

		c.HandleEvent(hub, db, ctx, msg)
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

func (c *Client) HandleEvent(hub *Hub, db *sql.DB, ctx context.Context, rawMsg []byte) {
	var evt data.WebSocketEvent
	if err := json.Unmarshal(rawMsg, &evt); err != nil {
		log.Println("invalid event:", err)
		return
	}
	switch evt.Type {
	case enum.EventJoinRoom:
		var payload data.RoomPayload
		if err := json.Unmarshal(evt.Data, &payload); err != nil {
			log.Println("invalid join_room payload:", err)
			return
		}
		log.Printf("User %d joining room %d", c.userId, payload.RoomId)

		if c.roomId != 0 { //Neu da join room truoc do thi leave di
			oldRoom := hub.GetRoom(c.roomId)
			oldRoom.Leave(c.userId)
		}

		c.roomId = payload.RoomId
		newRoom := hub.GetRoom(c.roomId)
		newRoom.Join(c)
	case enum.EventSendMessage:
		var incoming dto.MessageDTO
		if err := json.Unmarshal(evt.Data, &incoming); err != nil {
			log.Println("invalid message:", err)
			return
		}

		incoming.SenderId = c.userId
		room := hub.GetRoom(c.roomId)
		room.Broadcast(hub, db, ctx, incoming, c)
	case enum.EventTyping:
		var payload data.RoomPayload
		if err := json.Unmarshal(evt.Data, &payload); err != nil {
			log.Println("invalid typing payload:", err)
			return
		}
		room := hub.GetRoom(payload.RoomId)
		typingEvent := data.WebSocketEvent{
			Type: enum.EventTyping,
			Data: json.RawMessage(fmt.Sprintf(`{"user_id": %d}`, c.userId)),
		}
		room.BroadcastToRoomExceptSender(c, typingEvent)
	case enum.EventLeaveRoom:
		var payload data.RoomPayload
		if err := json.Unmarshal(evt.Data, &payload); err != nil {
			log.Println("invalid leave_room payload:", err)
			return
		}
		room := hub.GetRoom(payload.RoomId)
		room.Leave(c.userId)

		if c.roomId == payload.RoomId {
			c.roomId = 0
		}
	default:
		log.Println("invalid event:", evt.Type)
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
			r.Info.IsGroup = msg.IsSendGroup
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

		sender.roomId = roomId

		//Update hub
		hub.mutex.Lock()
		hub.rooms[roomDTO.Id] = r
		delete(hub.rooms, 0)
		hub.mutex.Unlock()
	}

	repo := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(repo)

	msg.RoomId = r.Info.Id
	if msg.IsSendGroup { //neu la gui tin nhan nhom thi receiverId = roomId
		msg.ReceiverId = r.Info.Id
	}
	createMsg := dto.MapMessageToCreateDTO(msg)

	saveMsgErr := messageService.SendMessage(ctx, sender.userId, createMsg)
	if saveMsgErr == nil {
		//TODO: save message history
	}

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

func (r *Room) BroadcastToRoomExceptSender(sender *Client, event data.WebSocketEvent) {
	eventData, err := json.Marshal(event)
	if err != nil {
		log.Println("broadcast marshal error:", err)
		return
	}

	for userId, client := range r.Clients {
		if sender != nil && userId == sender.userId {
			continue
		}
		select {
		case client.send <- string(eventData):
		default:
			log.Printf("user %d send buffer full", userId)
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
