package main

import (
	taskTwentyOne "WBTechL1/21"
)

func main() {
	emailSvc := &taskTwentyOne.EmailService{From: "noreply@acme.local"}
	smsSvc := &taskTwentyOne.SMSService{SenderID: "ACME"}

	emailNotifier := taskTwentyOne.NewEmailAdapter(emailSvc, "ops@acme.local", "ALERT")
	smsNotifier := taskTwentyOne.NewSMSAdapter(smsSvc, "+1234567890")

	longMsg := "Service payment-processor is down\nerror: connection refused\nplease investigate immediately"

	// Подставляем разные адаптеры в один код клиента, не меняя legacy-логику
	taskTwentyOne.SendAlert(emailNotifier, longMsg)
	taskTwentyOne.SendAlert(smsNotifier, longMsg)
}
