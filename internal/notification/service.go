package notification

import (
	"context"
	"fmt"

	"github.com/soloden/notificator-bot/pkg/logging"
)

type service struct {
	storage Storage
	logger  logging.Logger
}

func NewService(storage Storage, logger logging.Logger) *service {
	return &service{
		storage: storage,
		logger:  logger,
	}
}

type Service interface {
	Create(context.Context, NotificationDto) (string, error)
	GetOne(context.Context, int) (Notification, error)
	GetMany(context.Context, []int) ([]Notification, error)
	Update(context.Context, NotificationDto) error
	Delete(context.Context, int) error
}

func (s *service) Create(ctx context.Context, dto NotificationDto) (string, error) {

	nNtf := NewNotification(dto)

	ntfId, err := s.storage.Create(ctx, nNtf)
	if err != nil {
		s.logger.Error(fmt.Errorf("failed to create new notification due err - %s", err))
		return ntfId, fmt.Errorf("ошибка при создании уведомления")
	}

	return ntfId, nil
}

func (s *service) GetOne(ctx context.Context, id int) (ntf Notification, err error) {
	return
}

func (s *service) GetMany(ctx context.Context, ids []int) (ntfList []Notification, err error) {
	return
}

func (s *service) Update(ctx context.Context, dto NotificationDto) error {
	return nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return nil
}
