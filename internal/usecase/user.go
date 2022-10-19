package usecase

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/internal/storage/pgSQL"
)

func RunUser(user *domain.User, config *configs.Configuration) *pgSQL.WriteDBResult {
	conn, err := pgSQL.NewConnectToDataBase(config)
	if err != nil {
		return &pgSQL.WriteDBResult{Err: err}
	}

	result := pgSQL.WriteUserToDB(user, conn)

	return &pgSQL.WriteDBResult{Err: result}
}

func NewUser(name string, id string, active bool) *domain.User {
	return &domain.User{
		Username: name,
		UserId:   id,
		IsActive: active,
	}
}
