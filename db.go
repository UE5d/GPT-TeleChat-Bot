package main

import (
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbFile = "chats.db"
	DB     *gorm.DB
)

type Message struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey" json:"id"`
	ChatID  string `json:"chatId,omitempty"`  // telegrams conversation id
	Role    string `json:"role,omitempty"`    // chatgpt role
	Content string `json:"content,omi