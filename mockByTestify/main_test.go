package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type EmailServiceMock struct {
	mock.Mock
}

func (e *EmailServiceMock) SendMessage(content string) error {
	fmt.Println("Mock")
	args := e.Called(content)
	return args.Error(0)
}

func TestCharge(t *testing.T) {
	emailService := new(EmailServiceMock)
	emailService.On("SendMessage", "100").Return(nil)

	myService := MyMessageService{emailService}
	myService.Charge(100)

	emailService.AssertExpectations(t)
}
