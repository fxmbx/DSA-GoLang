package example

import "fmt"

type NotificationResponse struct {
	DestinationID string
	err           error
}
type Notification interface {
	Send() NotificationResponse
}

type SmsNotification struct {
	DestinationID string
	Message       string
}

func (sms *SmsNotification) Send() NotificationResponse {
	fmt.Println("Sending SMS notification")
	fmt.Println(sms.Message)
	if len(sms.DestinationID) > 0 {
		fmt.Println("Message delivered to: ", sms.DestinationID)
	}
	return simulateResponse(sms.DestinationID)
}

type SlackNotification struct {
	Notification Notification
}

func (slackNotification SlackNotification) Send() NotificationResponse {
	response := slackNotification.Notification.Send()
	if len(response.DestinationID) > 0 {
		slackUserID := "slack-user-01"
		fmt.Println("sending to slack: ", slackUserID)
	}
	return simulateResponse(response.DestinationID)
}
func simulateResponse(destinationID string) NotificationResponse {
	return NotificationResponse{destinationID, nil}
}
