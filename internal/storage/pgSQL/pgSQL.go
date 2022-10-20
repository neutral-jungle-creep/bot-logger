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
	user   *domain.User
	conn   *pgx.Conn
	config *configs.Configuration
}

type editUser struct {
	user   *domain.User
	conn   *pgx.Conn
	config *configs.Configuration
}

func NewAddUser(user *domain.User, conn *pgx.Conn, config *configs.Configuration) *addUser {
	return &addUser{
		user:   user,
		conn:   conn,
		config: config,
	}
}

func NewEditUser(user *domain.User, conn *pgx.Conn, config *configs.Configuration) *editUser {
	return &editUser{
		user:   user,
		conn:   conn,
		config: config,
	}
}

func (a *addUser) DBWrite() error {
	_, err := a.conn.Query(context.Background(), a.config.Queries.AddUser,
		a.user.Username,
		a.user.UserId,
		a.user.IsActive,
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
		e.user.IsActive,
		e.user.UserId,
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
}

type editMessage struct {
	message *domain.Message
	conn    *pgx.Conn
	config  *configs.Configuration
}

func NewAddMessage(message *domain.Message, conn *pgx.Conn, config *configs.Configuration) *addMessage {
	return &addMessage{
		message: message,
		conn:    conn,
		config:  config,
	}
}

func NewEditMessage(message *domain.Message, conn *pgx.Conn, config *configs.Configuration) *editMessage {
	return &editMessage{
		message: message,
		conn:    conn,
		config:  config,
	}
}

func (e *editMessage) DBWrite() error {
	_, err := e.conn.Query(context.Background(), e.config.Queries.EditMessage,
		e.message.MessageText.Query,
		e.message.MessageText.Problem,
		e.message.MessageText.Cause,
		e.message.MessageText.Solution,
		e.message.MessageText.Source,
		//message.V4Data,
		e.message.IsEdit,
		e.message.MessageId,
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
		a.message.MessageId,
		a.message.Date,
		a.message.MessageText.Query,
		a.message.MessageText.Problem,
		a.message.MessageText.Cause,
		a.message.MessageText.Solution,
		a.message.MessageText.Source,
		//message.V4Data,
		a.message.IsEdit,
		a.message.MessageSender.UserId,
	)

	if err != nil {
		log.Println(logs.ErrAddDBM, err)
		return err
	}
	log.Println(logs.AddDBM, *a.message)
	return nil
}
