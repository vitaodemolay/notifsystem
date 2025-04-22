package model

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
}

func NewCampaign(title, content string, emails []string) (*Campaign, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	} else if content == "" {
		return nil, errors.New("content cannot be empty")
	} else if len(emails) == 0 {
		return nil, errors.New("at least one email is required")
	}

	contacts, err := createContacts(emails)
	if err != nil {
		return nil, err
	}

	return &Campaign{
		ID:        xid.New().String(),
		Title:     title,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}

func createContacts(emails []string) ([]Contact, error) {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contact, err := NewContact(email)
		if err != nil {
			return nil, err
		}
		contacts[i] = *contact
	}
	return contacts, nil
}
