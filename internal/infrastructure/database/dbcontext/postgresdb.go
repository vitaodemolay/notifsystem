package dbcontext

import (
	"github.com/vitaodemolay/notifsystem/internal/domain/model/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPgDb(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{}, &campaign.CampaignStatus{})

	return db, nil
}
