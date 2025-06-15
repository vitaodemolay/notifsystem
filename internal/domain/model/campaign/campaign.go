package campaign

import (
	"slices"
	"time"

	"github.com/rs/xid"
	internalerrors "github.com/vitaodemolay/notifsystem/pkg/internal-errors"
)

type Campaign struct {
	ID         string    `json:"id" validate:"required" gorm:"size:50,primaryKey"`
	Title      string    `json:"title" validate:"min=10,max=30" gorm:"size:30"`
	CreatedAt  time.Time `json:"created_at" validate:"required"`
	Content    string    `json:"content" validate:"min=10,max=1024" gorm:"size:1024"`
	Contacts   []Contact `json:"contacts" validate:"min=1,dive"`
	StatusList []CampaignStatus
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

	campain.setPendingStatus()

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

func (c *Campaign) setPendingStatus() {
	c.addStatus(pendingStatus)
}

func (c *Campaign) addStatus(status CampaignStatusType) error {
	newStatus, err := newCamapaignStatus(status)
	if err != nil {
		return err
	}
	c.StatusList = append(c.StatusList, *newStatus)
	return nil
}

func (c *Campaign) GetActualStatus() CampaignStatus {
	if len(c.StatusList) == 0 {
		c.setPendingStatus()
	}
	return slices.MaxFunc(c.StatusList, func(a, b CampaignStatus) int {
		return a.CreatedAt.Compare(b.CreatedAt)
	})
}
