package entity

type Room struct {
	BaseEntity
	Name string `json:"name"`
}

func (Room) TableName() string {
	return "room"
}
