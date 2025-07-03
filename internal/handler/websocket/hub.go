package websocket

import (
	"chatapp/internal/dto"
	"log"
	"sync"
)

type Room struct {
	Info    dto.RoomDTO
	Clients map[*Client]bool
	mutex   sync.Mutex
}

type Hub struct {
	rooms map[int64]*Room
	mutex sync.Mutex
}

func NewRoom(dto dto.RoomDTO) *Room {
	return &Room{
		Info:    dto,
		Clients: make(map[*Client]bool),
	}
}

func NewHub(roomDTOs []dto.RoomDTO) *Hub {
	hub := Hub{
		rooms: make(map[int64]*Room),
	}

	for _, data := range roomDTOs {
		room := NewRoom(data)
		hub.rooms[data.Id] = room
	}
	return nil
}

// User join room
func (h *Hub) Join(roomId int64, c *Client) {
	h.mutex.Lock()
	room, ok := h.rooms[roomId]
	h.mutex.Unlock()

	if !ok {
		log.Printf("room %d not found in hub", roomId)
		return
	}

	room.mutex.Lock()
	defer room.mutex.Unlock()
	room.Clients[c] = true
}

// User leave room
func (h *Hub) Leave(roomId int64, c *Client) {
	h.mutex.Lock()
	room, ok := h.rooms[roomId]
	h.mutex.Unlock()

	if !ok {
		return
	}

	room.mutex.Lock()
	defer room.mutex.Unlock()

	delete(room.Clients, c)
}

// Broadcast: gửi tin nhắn tới tất cả client trong room, trừ sender
func (h *Hub) Broadcast(roomId int64, msg string, sender *Client) {
	h.mutex.Lock()
	room, ok := h.rooms[roomId]
	h.mutex.Unlock()

	if !ok {
		return
	}

	room.mutex.Lock()
	defer room.mutex.Unlock()

	for client := range room.Clients {
		if client.userId != sender.userId {
			client.Send(msg)
		}
	}
}
