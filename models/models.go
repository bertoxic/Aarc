package models

import (
	"time"
)

type User struct {
	Name string
	Email              string
	Password           string
	Verified           bool
	Verification       string
	VerificationExpiry time.Time
	Uuid   			   string
}

type MailData struct {
	To      string
	From    string
	Subject string
	Content string
	Template string
}

type ComputerData struct {
	ScreenSize int
	Cpu int
	StorageSize int
	RamSize int
	Gen int
	Resolution int
}