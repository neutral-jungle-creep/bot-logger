package telegram

import (
	"bot_logger/internal/service"
	"bot_logger/internal/storage"
	"github.com/jackc/pgx/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) Handler {
	return Handler{
		service: service,
	}
}

func HandlerComposite(conn *pgx.Conn) *Handler {
	stor := storage.NewStorage(conn)
	serv := service.NewService(stor)
	handler := NewHandler(serv)

	return &handler
}
