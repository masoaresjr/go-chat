package entity

type Message struct {
	BaseEntity
	Payload string `json:"payload"`
	Room    Room
	RoomID  int `json:"room_id"`
	User    User
	UserID  int `json:"user_id"`
}

func (Message) TableName() string {
	return "message"
}
