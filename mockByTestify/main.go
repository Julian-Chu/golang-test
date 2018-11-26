package main

import (
	"fmt"
	"strconv"
)

type MessageService interface {
	SendMessage(string) error
}

type EmailService struct{}

type MyMessageService struct {
	messageService MessageService
}

func (es EmailService) SendMessage(content string) error {
	fmt.Println("Send Email")
	return nil
}

func (ms MyMessageService) Charge(val int) error {
	ms.messageService.SendMessage(strconv.Itoa(val))
	fmt.Printf("Charge val %d\n", val)
	return nil
}

func main() {
	emailService := EmailService{}
	myMessageService := MyMessageService{emailService}
	myMessageService.Charge(10)
}
