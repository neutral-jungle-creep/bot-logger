package pgSQL

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/pkg/logs"
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func NewConnectToDataBase(config *configs.Configuration) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), config.LinkToDB)
	if err != nil {
		log.Println(logs.ErrConnDB, err)
		return conn, err
	}
	log.Println(logs.ConnCorrect)

	return conn, nil
}

type addUser struct {
	user     *domain.User
	conn     *pgx.Conn
	config   *configs.Configuration
	isActive bool
}

type editUser struct {
	user     *domain.User
	conn     *pgx.Conn
	config   *configs.Configuration
	isActive bool
}

func NewAddUser(user *domain.User, conn *pgx.Conn, config *configs.Configuration, isActive bool) *addUser {
	return &addUser{
		user:     user,
		conn:     conn,
		config:   config,
		isActive: isActive,
	}
}

func NewEditUser(user *domain.User, conn *pgx.Conn, config *configs.Configuration, isActive bool) *editUser {
	return &editUser{
		user:     user,
		conn:     conn,
		config:   config,
		isActive: isActive,
	}
}

func (a *addUser) DBWrite() error {
	_, err := a.conn.Query(context.Background(), a.config.Queries.AddUser,
		a.user.Username,
		a.user.Id,
		a.isActive,
	)

	if err != nil {
		log.Println(logs.ErrAddBDU, err)
		return err
	}
	log.Println(logs.AddDBU, *a.user)
	return nil
}

func (e *editUser) DBWrite() error {
	_, err := e.conn.Query(context.Background(), e.config.Queries.EditUser,
		e.isActive,
		e.user.Id,
	)

	if err != nil {
		log.Println(logs.ErrEditDBU, err)
		return err
	}
	log.Println(logs.EditBDU, *e.user)
	return nil
}

type addMessage struct {
	message *domain.Message
	conn    *pgx.Conn
	config  *configs.Configuration
	isEdit  bool
}

type editMessage struct {
	message *domain.Message
	conn    *pgx.Conn
	config  *configs.Configuration
	isEdit  bool
}

func NewAddMessage(message *domain.Message, conn *pgx.Conn, config *configs.Configuration, isEdit bool) *addMessage {
	return &addMessage{
		message: message,
		conn:    conn,
		config:  config,
		isEdit:  isEdit,
	}
}

func NewEditMessage(message *domain.Message, conn *pgx.Conn, config *configs.Configuration, isEdit bool) *editMessage {
	return &editMessage{
		message: message,
		conn:    conn,
		config:  config,
		isEdit:  isEdit,
	}
}

func (e *editMessage) DBWrite() error {
	_, err := e.conn.Query(context.Background(), e.config.Queries.EditMessage,
		e.message.Text,
		e.isEdit,
		e.message.Id,
	)

	if err != nil {
		log.Println(logs.ErrEditBDM, err)
		return err
	}
	log.Println(logs.EditDBM, *e.message)
	return nil
}

func (a *addMessage) DBWrite() error {
	_, err := a.conn.Query(context.Background(), a.config.Queries.AddMessage,
		a.message.Id,
		a.message.Date,
		a.message.Text,
		a.isEdit,
		a.message.SenderId,
	)

	if err != nil {
		log.Println(logs.ErrAddDBM, err)
		return err
	}
	log.Println(logs.AddDBM, *a.message)
	return nil
}
