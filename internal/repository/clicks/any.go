package clicks

import "gorm.io/gorm"

type Repository struct {
	Gorm *gorm.DB
}

func NewClickRepository(gorm *gorm.DB) *Repository {
	return &Repository{
		Gorm: gorm,
	}
}
