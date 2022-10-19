package pgSQL

import (
	"bot_logger/configs"
	"bot_logger/internal/domain"
	"bot_logger/pkg/logs"
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

const (
	addMessage = `INSERT INTO public."test_logs" (message_id, date, query, problem, cause, solution, source,
                 is_edit, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 
                                                   (SELECT id FROM public."test_users"
	 												WHERE tg_user_id = $9)
													)`
	editMessage = `UPDATE public."test_logs" SET query=$1, problem=$2, cause=$3, solution=$4, source=$5, is_edit=$6
	               WHERE message_id=$7;`
	addUser  = `INSERT INTO public."test_users" (user_name, tg_user_id, active_employee) VALUES ($1, $2, $3)`
	editUser = `UPDATE public."test_users" SET active_employee=$1 WHERE tg_user_id=$2;`
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

func WriteUserToDB(user *domain.User, conn *pgx.Conn) error {
	if user.IsActive {
		return addUserToDB(user, conn)
	} else {
		return editUserInDB(user, conn)
	}
}

func addUserToDB(user *domain.User, conn *pgx.Conn) error {
	_, err := conn.Query(context.Background(), addUser,
		user.Username,
		user.UserId,
		user.IsActive,
	)

	if err != nil {
		log.Println(logs.ErrAddBDU, err)
		return err
	}
	log.Println(logs.AddDBU, user)
	return nil
}

func editUserInDB(user *domain.User, conn *pgx.Conn) error {
	_, err := conn.Query(context.Background(), editUser,
		user.IsActive,
		user.UserId,
	)

	if err != nil {
		log.Println(logs.ErrEditDBU, err)
		return err
	}
	log.Println(logs.EditBDU, user)
	return nil
}

func WriteMessageToDB(message *domain.Message, conn *pgx.Conn) error {
	if message.IsEdit {
		return editMessageInDB(message, conn)
	} else {
		return addMessageToDB(message, conn)
	}
}

func editMessageInDB(message *domain.Message, conn *pgx.Conn) error {
	_, err := conn.Query(context.Background(), editMessage,
		message.MessageText.Query,
		message.MessageText.Problem,
		message.MessageText.Cause,
		message.MessageText.Solution,
		message.MessageText.Source,
		//message.V4Data,
		message.IsEdit,
		message.MessageId,
	)

	if err != nil {
		log.Println(logs.ErrEditBDM, err)
		return err
	}
	log.Println(logs.EditDBM, message)
	return nil
}

func addMessageToDB(message *domain.Message, conn *pgx.Conn) error {
	_, err := conn.Query(context.Background(), addMessage,
		message.MessageId,
		message.Date,
		message.MessageText.Query,
		message.MessageText.Problem,
		message.MessageText.Cause,
		message.MessageText.Solution,
		message.MessageText.Source,
		//message.V4Data,
		message.IsEdit,
		message.MessageSender.UserId,
	)

	if err != nil {
		log.Println(logs.ErrAddDBM, err)
		return err
	}
	log.Println(logs.AddDBM, message)
	return nil
}
