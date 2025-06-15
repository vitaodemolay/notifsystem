package campaign

import "errors"

type Contact struct {
	Email string `json:"email" validate:"required,email"`
}

func newContact(email string) (*Contact, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	return &Contact{
		Email: email,
	}, nil
}
