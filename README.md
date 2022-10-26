# bot-logger

## назначение

Бот слушает события в указанном чате и записывает их в базу данных **postgresql**. События из других чатов бот игнорирует, отправлет в ответ «Нет доступа». Программа может быть полезна в случаях, где история переписки очень важна, как пример: рабочие чаты

## обработка событий

События, которые на данный момент обрабатывает бот:

- добавление нового участника или его самостоятельный вход в чат
- удаление участника или его самостоятельный выход из чата
- появление нового текстового сообщения
- редактирование текстового сообщения

## хранение данных

В базе данных информация будет храниться в двух таблицах: участники и их сообщения, таблицы должны иметь связь один ко многим.

**участники**

поля: id, user_name, tg_user_id, active_user

- при добавлении нового пользователя, создается запись о нем: его порядковый номер в таблице, никней в телеграм, телеграм айди, булево значение _true_ - это означает что пользователь является участновом чата.
- при удалении пользователя, в столбце **active_user** значение меняется на _false_.
- при добавлении в чат пользователя, который в нем уже когда-то состоял, в столбце **active_user** у этого пользователя значение меняется на _true_.

**Код для создания таблицы:**

```postgresql
CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    user_name character varying(100) COLLATE pg_catalog."default",
    tg_user_id integer,
    active_user boolean,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)
```
**сообщения**

поля: id, message_id, user_id, date, text, is_edit

- при появлении нового текстового сообщения в чате, создается запись о нем: его порядковый номер в таблице, номер сообщения в чате, порядковый номер отправителя из таблицы с пользователями, дата и время отправки сообщения, содержание сообщения, булево значение _false_ - это означает что сообщение не было отредактированно.
- при редактировании сообщения, в столбце **is_edit** значение меняется на _true_, в столбце **text** содержание меняется на актуальное. 

**Код для создания таблицы:**

```postgresql
CREATE TABLE IF NOT EXISTS public.messages
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    message_id integer,
    user_id integer,
    date timestamp without time zone,
    text character varying(100) COLLATE pg_catalog."default",
    is_edit boolean,
    CONSTRAINT messages_pkey PRIMARY KEY (id),
    CONSTRAINT messages_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
```

## инструкция для настройки

Для запуска бота требуется файл конфигурации **config.yml**, который имеет вид:

```yml
token: "*********************************"
logFile: "../logs.log"
unwrittenDataFile: "../unwritten_data.json"
accessChatID: -11111111
adminsTgChatID: 11111111
linkToDB: "postgres://username:password@localhost:5432/database_name"
queries:
  addUser: "INSERT INTO public.\"users\" (tg_user_id, user_name, active_user) VALUES ($1, $2, $3);"
  editUser: "UPDATE public.\"users\" SET active_user=$1 WHERE tg_user_id=$2;"
  addMessage: "INSERT INTO public.\"messages\" (message_id, date, text, is_edit, user_id) VALUES ($1, $2, $3, $4, (SELECT id FROM public.\"users\" WHERE tg_user_id = $5));"
  editMessage: "UPDATE public.\"messages\" SET text=$1, is_edit=$2 WHERE message_id=$3;"
```

### пояснения к полям:

- **token** - токен от телеграм бота который в чате имеет права администратора
- **logFile** - название файла для записи логов программы
- **unwrittenDataFile** - имя файла, который создется при не корректном завершении работы бота. [подробнее](#исключения)
- **accessChatID** - чат в котором бот не будет игнорировать сообщения.
- **adminsTgChatID** - пользователь, который будет получать в личные сообщения уведомления об остановке работы бота. [подробнее](#исключения)
- **linkToDB** - ссылка для подключения к базе данных
- **queries** - sql запросы с помощью которых будет осуществляться добавление и редактирование данных в таблицах.

## исключения

**ВАЖНО!** Для того, чтобы бот смог отправить сообщение, получатель должен его запустить.

Существует два вида исключений:

- Не критичное: в чат поступит два сообщения. При этом программой будет создан файл с именем, указанном в конфигурации по ключу **unwrittenDataFile**. В него будет записано событие чата, прочитанное, но не записанное ботом в базу данных.

![img.png](img/err.png)

- Критичное: в чат не  поступит сообщение о том, что прочитанные данные сохранены в файле, это значит что произошла непредвиденная ошибка. При возникновении подобных исключений можно обратиться в телеграм [@arrrrseniy](https://t.me/arrrrseniy)

![img_1.png](img/fatal_err.png)

При перезапуске бота событие чата, находящееся в файле с именем **unwrittenDataFile** 
будет обработано и записано в базу данных, файл будет автоматически удален.