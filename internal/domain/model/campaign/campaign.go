package campaign

import (
	"time"

	"github.com/rs/xid"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
)

type Campaign struct {
	ID        string    `json:"id" validate:"required"`
	Title     string    `json:"title" validate:"min=10,max=30"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	Content   string    `json:"content" validate:"min=10,max=1024"`
	Contacts  []Contact `json:"contacts" validate:"min=1,dive"`
}

func NewCampaign(title, content string, emails []string) (*Campaign, error) {
	contacts, err := createContacts(emails)
	if err != nil {
		return nil, err
	}

	campain := &Campaign{
		ID:        xid.New().String(),
		Title:     title,
		CreatedAt: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

	if err := internalerrors.ValidateStruct(campain); err != nil {
		return nil, err
	}
	return campain, nil
}

func createContacts(emails []string) ([]Contact, error) {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contact, err := newContact(email)
		if err != nil {
			return nil, err
		}
		contacts[i] = *contact
	}
	return contacts, nil
}
