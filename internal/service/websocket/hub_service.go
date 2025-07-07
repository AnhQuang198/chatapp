package websocket

import (
	"chatapp/internal/dto"
	"chatapp/internal/repository"
	"chatapp/internal/service"
	"context"
	"database/sql"
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

	repo := repository.NewRoomRepository(db)
	roomService := service.NewRoomService(repo)

	if !ok {
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
		if room.Info.Id != 0 {
			roomDTO := room.Info
			roomDTO.UserIds = append(roomDTO.UserIds, c.userId)
			//TODO: update userId for user_ids in room
			if err := roomService.UpdateRoomUserId(ctx, roomDTO); err != nil {
				log.Println("Error updating room:", roomDTO)
				return
			}
		}
	}

	room.Join(c)
	log.Printf("User %d joined room %d", c.userId, roomId)
}

func (h *Hub) GetRoom(roomId int64) *Room {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	return h.rooms[roomId]
}
