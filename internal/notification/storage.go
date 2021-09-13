package notification

import "context"

type Storage interface {
	Create(ctx context.Context, ntf Notification) (string, error)
	GetOne(ctx context.Context, id string) (Notification, error)
	GetNotificationByUserId(ctx context.Context, userId string) ([]Notification, error)
	Update(ctx context.Context, ntf Notification) error
	Delete(ctx context.Context, id string) error
}
