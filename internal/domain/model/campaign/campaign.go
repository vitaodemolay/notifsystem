package campaign

import (
	"time"

	"github.com/rs/xid"
	model "github.com/vitaodemolay/notifsystem/internal/domain/model/contact"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
)

type CampaignStatus string

const (
	Pending CampaignStatus = "Pending"
	Started CampaignStatus = "Started"
	Done    CampaignStatus = "Done"
)

type Campaign struct {
	ID        string          `json:"id" validate:"required"`
	Title     string          `json:"title" validate:"min=10,max=30"`
	CreatedAt time.Time       `json:"created_at" validate:"required"`
	Content   string          `json:"content" validate:"min=10,max=2048"`
	Contacts  []model.Contact `json:"contacts" validate:"min=1,dive"`
	Status	  CampaignStatus
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
		Status:    Pending,
	}

	if err := internalerrors.ValidateStruct(campain); err != nil {
		return nil, err
	}
	return campain, nil
}

func createContacts(emails []string) ([]model.Contact, error) {
	contacts := make([]model.Contact, len(emails))
	for i, email := range emails {
		contact, err := model.NewContact(email)
		if err != nil {
			return nil, err
		}
		contacts[i] = *contact
	}
	return contacts, nil
}
