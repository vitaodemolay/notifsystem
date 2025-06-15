package campaign

import (
	"errors"

	"github.com/rs/xid"
)

type Contact struct {
	ID         string `json:"id"`
	Email      string `json:"email" validate:"required,email"`
	CampaignID string `json:"campaign_id"`
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
