package main

import (
	"dsa-golang/Chapter1/StructuralPatterns/decorator/example"
	"fmt"
)

type IProcess interface {
	process()
}

type ProcessClass struct {
}

func (processCless *ProcessClass) process() {
	fmt.Println("Process class implementing process")
}

type ProcessDecorator struct {
	processInstance *ProcessClass
}

func (processDecorator *ProcessDecorator) process() {
	if processDecorator.processInstance == nil {
		fmt.Println("Process decorator implementing process")
	} else {
		fmt.Print("Process decorator implementing processing and ")
		processDecorator.processInstance.process()
	}
}

func main() {
	var process = &ProcessClass{}
	var processDecorator = &ProcessDecorator{
		processInstance: process,
	}
	process.process()
	processDecorator.process()

	smsNotification := &example.SmsNotification{
		DestinationID: "a lot",
		Message:       "how many niggars you shot?",
	}
	slackNotification := &example.SlackNotification{
		Notification: smsNotification,
	}
	slackNotification.Send()
}
