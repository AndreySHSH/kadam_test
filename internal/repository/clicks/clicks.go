package clicks

import (
	"gorm.io/gorm/clause"
	"kadam_test/internal/migration"
	homework "kadam_test/internal/proto"
)

func (c *Repository) CheckUniqueness(id string) (bool, error) {
	var clickModel []migration.Click

	if err := c.Gorm.Model(&clickModel).Select("id").Where("id = ?", id).Find(&clickModel).Error; err != nil {
		return true, err
	}

	if len(clickModel) == 0 {
		return true, nil
	}
	return false, nil
}

func (c *Repository) Create(click *homework.Click, IsSuspicious bool) error {
	var clickModel migration.Click

	clickModel.Id = click.Id
	clickModel.Ua = click.Ua
	clickModel.Cpc = click.Cpc
	clickModel.RedirectUri = click.RedirectUri
	clickModel.IsSuspicious = IsSuspicious

	if err := c.Gorm.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"is_suspicious",
		}),
	}).Create(&clickModel).Error; err != nil {
		return err
	}
	return nil
}
