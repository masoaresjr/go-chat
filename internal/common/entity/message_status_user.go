package entity

type MessageStatusUser struct {
	BaseEntity
	Message             Message
	MessageID           int `json:"message_id"`
	MessageStatusType   MessageStatusType
	MessageStatusTypeID int `json:"message_status_type_id"`
	User                User
	UserID              int `json:"user_id"`
}

func (MessageStatusUser) TableName() string {
	return "message_status_user"
}
