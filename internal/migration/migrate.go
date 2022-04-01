package migration

import "gorm.io/gorm"

// CreateSchema - Make migrate DB
func CreateSchema(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Click{},
	)

	if err != nil {
		return err
	}
	return nil
}
