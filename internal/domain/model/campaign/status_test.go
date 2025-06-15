package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newCamapaignStatus_When_statusTypeIsEmpty(t *testing.T) {
	assert := assert.New(t)

	status, err := newCamapaignStatus("")

	assert.Nil(status)
	assert.ErrorIs(err, ErrInvalidCampaignStatus)
}

func Test_newCamapaignStatus_When_statusTypeIsInvalid(t *testing.T) {
	assert := assert.New(t)

	status, err := newCamapaignStatus("InvalidStatus")

	assert.Nil(status)
	assert.ErrorIs(err, ErrInvalidCampaignStatus)
}

func Test_newCamapaignStatus_When_statusTypeIsValid(t *testing.T) {
	assert := assert.New(t)

	type caseTest struct {
		title    string
		status   CampaignStatusType
		espected func(stat *CampaignStatus)
	}
	cases := []caseTest{
		{
			title:  "Pending status",
			status: pendingStatus,
			espected: func(stat *CampaignStatus) {
				assert.Equal(stat.Value, pendingStatus)
				assert.NotZero(stat.CreatedAt)
			},
		},
		{
			title:  "Canceled status",
			status: canceledStatus,
			espected: func(stat *CampaignStatus) {
				assert.Equal(stat.Value, canceledStatus)
				assert.NotZero(stat.CreatedAt)
			},
		},
		{
			title:  "Deleted status",
			status: deletedStatus,
			espected: func(stat *CampaignStatus) {
				assert.Equal(stat.Value, deletedStatus)
				assert.NotZero(stat.CreatedAt)
			},
		},
		{
			title:  "Started status",
			status: startedStatus,
			espected: func(stat *CampaignStatus) {
				assert.Equal(stat.Value, startedStatus)
				assert.NotZero(stat.CreatedAt)
			},
		},
		{
			title:  "Done status",
			status: doneStatus,
			espected: func(stat *CampaignStatus) {
				assert.Equal(stat.Value, doneStatus)
				assert.NotZero(stat.CreatedAt)
			},
		},
	}
	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			status, err := newCamapaignStatus(c.status)
			assert.NotNil(status)
			assert.NoError(err)
			c.espected(status)
		})
	}
}
