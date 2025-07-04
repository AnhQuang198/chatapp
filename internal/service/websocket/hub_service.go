package websocket

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/internal/service"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"sync"
)

type Hub struct {
	rooms map[int64]*Room
	mutex sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[int64]*Room),
	}
}

func (h *Hub) Join(ctx context.Context, db *sql.DB, roomId int64, c *Client) {
	h.mutex.Lock()
	room, ok := h.rooms[roomId]
	h.mutex.Unlock()

	if !ok {
		repo := repository.NewRoomRepository(db)
		roomService := service.NewRoomService(repo)

		roomDTO, err := roomService.GetRoomById(ctx, roomId)
		persisted := false
		if err != nil {
			log.Println("Room not found, creating transient room:", roomId)
			roomDTO = dto.NewRoomDTO(roomId)
		} else {
			persisted = true
		}
		room = &Room{
			Info:      roomDTO,
			Clients:   make(map[int64]*Client),
			Persisted: persisted,
		}
		h.mutex.Lock()
		h.rooms[roomId] = room
		h.mutex.Unlock()
	} else {
		//TODO: update userId for user_ids in room
	}

	room.Join(c)

	if room.Persisted {
		repo := repository.NewMessageRepository(db)
		messageService := service.NewMessageService(repo)
		lstMessages, _ := messageService.GetMessageByRoomId(ctx, roomId)
		for _, m := range lstMessages {
			data, _ := json.Marshal(m)
			c.send <- string(data)
		}
	}
}

func (h *Hub) GetRoom(roomId int64) *Room {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	return h.rooms[roomId]
}
