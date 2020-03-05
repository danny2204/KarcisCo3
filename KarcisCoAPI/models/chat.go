package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Chat struct {
	Id int `gorm:primary_key`
	User1 int
	User2 int
	Content string `sql:longtext`
	CreatedAt time.Time
	UpdateAt time.Time
	Deleted *time.Time `sql: index`
}

func GetAllChat(id int)([]Chat, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var chatList []Chat

	db.Where("User1 = ? OR User2 = ?", id, id).Find(&chatList)
	return chatList, err
}

func CreateChat(user1 int, user2 int, content string) (*Chat, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	chat := &Chat {
		User1: user1,
		User2: user2,
		Content: content,
	}
	db.Save(chat)
	return chat, err
}

func UpdateChat(id int, user1 int, user2 int, content string)(*Chat, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var chat Chat
	db.Model(&chat).Where("id = ?", id).Updates(map[string]interface{}{"user1": user1, "user2": user2, "content": content})
	db.Where("id = ?", id).Find(&chat)
	return &chat, nil
}

func GetChatById(id int)(*Chat, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var chat Chat
	db.Where("id = ?", id).Find(&chat)
	return &chat, err

}
