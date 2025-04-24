package contact

import "errors"

type Contact struct {
	Email string `json:"email"`
}

func NewContact(email string) (*Contact, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	return &Contact{
		Email: email,
	}, nil
}
