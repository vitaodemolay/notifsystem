package internalerrors

import (
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

type userDummy struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"min=3,max=15"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"-"`
}

type testSuite struct {
	StructDummy userDummy
	Faker       faker.Faker
}

func setup() *testSuite {
	return &testSuite{
		Faker: faker.New(),
		StructDummy: userDummy{
			ID:    "123",
			Name:  "John Doe",
			Email: "john@test.com",
			Age:   30,
		},
	}
}

func Test_ValidateStruct_Valid(t *testing.T) {
	assert := assert.New(t)
	suite := setup()

	err := ValidateStruct(suite.StructDummy)
	assert.NoError(err, "Expected no error for valid struct")
}

func Test_ValidateStruct_WithRequiredFieldError(t *testing.T) {
	assert := assert.New(t)
	suite := setup()

	suite.StructDummy.ID = ""

	err := ValidateStruct(suite.StructDummy)
	assert.Error(err, "Expected error for required field")
	assert.Equal("id is required", err.Error(), "Expected specific error message")
}

func Test_ValidateStruct_WithMaxFieldError(t *testing.T) {
	assert := assert.New(t)
	suite := setup()

	suite.StructDummy.Name = suite.Faker.Lorem().Text(30)

	err := ValidateStruct(suite.StructDummy)
	assert.Error(err, "Expected error for max field")
	assert.Equal("name is required with max 15", err.Error(), "Expected specific error message")
}

func Test_ValidateStruct_WithMinFieldError(t *testing.T) {
	assert := assert.New(t)
	suite := setup()

	suite.StructDummy.Name = ""

	err := ValidateStruct(suite.StructDummy)
	assert.Error(err, "Expected error for min field")
	assert.Equal("name is required with min 3", err.Error(), "Expected specific error message")
}

func Test_ValidateStruct_WithEmailFieldError(t *testing.T) {
	assert := assert.New(t)
	suite := setup()

	suite.StructDummy.Email = "invalid-email"

	err := ValidateStruct(suite.StructDummy)
	assert.Error(err, "Expected error for email field")
	assert.Equal("email is invalid", err.Error(), "Expected specific error message")
}

func Test_ValidateStruct_WithEmptyStruct(t *testing.T) {
	assert := assert.New(t)
	suite := setup()

	suite.StructDummy = userDummy{}

	err := ValidateStruct(suite.StructDummy)
	assert.Error(err, "Expected error for empty struct")
	assert.Equal("id is required;name is required with min 3;email is required", err.Error(), "Expected specific error message for empty struct")
}

func Test_ValidateStruct_WithMultipleErrors(t *testing.T) {
	assert := assert.New(t)
	suite := setup()

	suite.StructDummy.ID = ""
	suite.StructDummy.Name = suite.Faker.Lorem().Text(30)
	suite.StructDummy.Email = "invalid-email"

	err := ValidateStruct(suite.StructDummy)
	assert.Error(err, "Expected error for multiple fields")
	assert.Equal(err.Error(), "id is required;name is required with max 15;email is invalid", "Expected specific error message for multiple fields")
}
