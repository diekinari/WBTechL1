package taskTwentyOne

import (
	"errors"
	"fmt"
	"time"
)

// Notifier - интерфейс уведомления, который ждет клиентский модуль
type Notifier interface {
	Notify(messag string) error
}

// --- Legacy логика, несовместимая с новым клиентским запросом ---

// EmailService - структура уведомления по email
type EmailService struct {
	From string
}

// Send отправляет уведомление по email
func (e *EmailService) Send(to, subject, body string) error {
	if to == "" {
		return errors.New("no recipient")
	}
	fmt.Printf("[EmailService] From:%s To:%s Subject:%s Body:%s\n", e.From, to, subject, body)
	return nil
}

// SMSService – эмуляция отдельной библиотеки со своей сигнатурой
type SMSService struct {
	SenderID string
}

// Push отпарвляет уведомление по СМС
func (s *SMSService) Push(to, shortMsg string) error {
	if to == "" {
		return errors.New("no phone")
	}
	fmt.Printf("[SMSService] From:%s To:%s Msg:%s\n", s.SenderID, to, shortMsg)
	return nil
}

// -- Адаптеры под требуемый интерфейс --

type EmailAdapter struct {
	Email *EmailService
	// Т.к. метод интерфейса ждет в аргументах только сообщение, будем задавать остальные параметры в полях структуры
	To      string
	Subject string
}

func NewEmailAdapter(email *EmailService, to, subject string) *EmailAdapter {
	return &EmailAdapter{Email: email, To: to, Subject: subject}
}

func (a *EmailAdapter) Notify(message string) error {
	// преобразуем сообщение под EmailService.Send
	body := fmt.Sprintf("Timestamp: %s\n\n%s", time.Now().Format(time.RFC3339), message)
	return a.Email.Send(a.To, a.Subject, body)
}

// SMSAdapter делает SMSService совместимым с Notifier
type SMSAdapter struct {
	SMS    *SMSService
	Number string
}

func NewSMSAdapter(sms *SMSService, number string) *SMSAdapter {
	return &SMSAdapter{SMS: sms, Number: number}
}

func (a *SMSAdapter) Notify(message string) error {
	// dummy logic
	if len(message) > 160 {
		message = message[:157] + "..."
	}
	return a.SMS.Push(a.Number, message)
}

// SendAlert – пример клиентского кода, работающего с Notifier (должен быть в отдельном пакете с интерфейсом)
func SendAlert(n Notifier, msg string) {
	if err := n.Notify(msg); err != nil {
		fmt.Printf("[SendAlert] failed: %s\n", err)
	}
}
