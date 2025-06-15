package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewContact_CreateContact(t *testing.T) {
	assert := assert.New(t)
	email := "email@test.com"

	contact, _ := newContact(email)

	assert.Equal(contact.Email, email)
}

func Test_NewContact_EmptyEmail(t *testing.T) {
	assert := assert.New(t)

	contact, err := newContact("")

	assert.Nil(contact)
	assert.Equal(err.Error(), "email cannot be empty")
}
