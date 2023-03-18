package entity

type MessageStatusType struct {
	BaseEntity
	Type string
}

func (MessageStatusType) TableName() string {
	return "message_status_type"
}
