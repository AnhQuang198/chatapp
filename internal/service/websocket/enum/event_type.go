package enum

type EventType string

const (
	EventJoinRoom     EventType = "join_room"
	EventSendMessage  EventType = "send_message"
	EventTyping       EventType = "typing"
	EventLeaveRoom    EventType = "leave_room"
	EventPing         EventType = "ping"
	EventPong         EventType = "pong"
	EventNotification EventType = "new_message_notification"
)
