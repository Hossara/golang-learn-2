package main

import (
	"errors"
	"main/alias"
	"time"
)

func main() {
	var sendSMSAction alias.Action = SendSMS
	sendSMSAction.DoPanic()
}

func SendSMS() error {
	println("Sending SMS...")
	time.Sleep(time.Second * 3)
	return errors.New("cannot send SMS")
}
