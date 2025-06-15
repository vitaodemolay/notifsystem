package campaign

import (
	"errors"

	"github.com/rs/xid"
)

type Contact struct {
	ID         string `json:"id" gorm:"size:50,primaryKey"`
	Email      string `json:"email" validate:"required,email" gorm:"size:100"`
	CampaignID string `json:"campaign_id" gorm:"size:50"`
}

func newContact(email string) (*Contact, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	return &Contact{
		ID:    xid.New().String(),
		Email: email,
	}, nil
}
