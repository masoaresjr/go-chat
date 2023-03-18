package entity

type RoomUser struct {
	BaseEntity
	Active bool
	Room   Room
	RoomID int `json:"room_id"`
	User   User
	UserID int `json:"user_id"`
}

func (RoomUser) TableName() string {
	return "room_user"
}
