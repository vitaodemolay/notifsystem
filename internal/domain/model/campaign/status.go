package campaign

import (
	"errors"
	"time"
)

type CampaignStatusType string

const (
	pendingStatus  CampaignStatusType = "Pending"
	canceledStatus CampaignStatusType = "Canceled"
	deletedStatus  CampaignStatusType = "Deleted"
	startedStatus  CampaignStatusType = "Started"
	doneStatus     CampaignStatusType = "Done"
)

var ErrInvalidCampaignStatus error = errors.New("invalid campaign status")

type CampaignStatus struct {
	ID         uint               `json:"id" gorm:"primaryKey"`
	Value      CampaignStatusType `json:"value" gorm:"size:15"`
	CampaignID string             `json:"campaign_id" gorm:"size:50"`
	CreatedAt  time.Time          `json:"created_at"`
}

func newCamapaignStatus(status CampaignStatusType) (*CampaignStatus, error) {
	if err := statusValidate(status); err != nil {
		return nil, err
	}

	return &CampaignStatus{
		Value:     status,
		CreatedAt: time.Now(),
	}, nil
}

func statusValidate(status CampaignStatusType) error {
	if status == "" {
		return ErrInvalidCampaignStatus
	}
	switch status {
	case pendingStatus, canceledStatus, deletedStatus, startedStatus, doneStatus:
		return nil
	default:
		return ErrInvalidCampaignStatus
	}
}
