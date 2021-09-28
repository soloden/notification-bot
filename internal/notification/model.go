package notification

type Notification struct {
	Id   int    `json:"id" bson:"_id"`
	Text string `json:"text" bson:"text,omitempty"`
	// ExecutionDate string `bson:"execution_date,omitempty"`
	CreatedAt int   `json:"created_at" bson:"created_at"`
	OwnerId   int64 `json:"owner_id" bson:"owner_id,omitempty"`
}

func NewNotification(dto NotificationDto) Notification {
	return Notification{
		Text: dto.Text,
		// ExecutionDate: dto.ExecutionDate,
		CreatedAt: dto.CreatedAt,
		OwnerId:   dto.OwnerId,
	}
}

type NotificationDto struct {
	Text string `bson:"text"`
	// ExecutionDate string `bson:"execution_date"`
	CreatedAt int   `bson:"created_at"`
	OwnerId   int64 `bson:"owner_id"`
}
