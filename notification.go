package fyne

import "time"

type Notification struct {
	Title, Content string
}

func NewNotification(title, content string) *Notification { _ = "STUB: not implemented"; return nil }

type ScheduledNotification struct {
	*Notification

	DeliveryTime time.Time

	id string
}

func (s *ScheduledNotification) ID() string { _ = "STUB: not implemented"; return "" }

func (s *ScheduledNotification) Cancel() error { _ = "STUB: not implemented"; return nil }

func NewScheduledNotification(id string, n *Notification, deliverAt time.Time) *ScheduledNotification {
	_ = "STUB: not implemented"
	return nil
}
