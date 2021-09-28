package notification

import (
	"context"
	"fmt"

	"github.com/soloden/notificator-bot/pkg/logging"
)

type Service struct {
	storage Storage
	logger  logging.Logger
}

func NewService(storage Storage, logger logging.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) Create(ctx context.Context, dto NotificationDto) (string, error) {

	nNtf := NewNotification(dto)

	ntfId, err := s.storage.Create(ctx, nNtf)
	if err != nil {
		s.logger.Error(fmt.Errorf("failed to create new notification due err - %s", err))
		return ntfId, fmt.Errorf("ошибка при создании уведомления")
	}

	return ntfId, nil
}

func (s *Service) GetOne(ctx context.Context, id int) (ntf Notification, err error) {
	return
}

func (s *Service) CreateMany(ctx context.Context, ids []int) (ntfList []Notification, err error) {
	return
}

func (s *Service) Update(ctx context.Context, dto NotificationDto) error {
	return nil
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return nil
}
