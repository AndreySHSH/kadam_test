package handlers

import (
	"kadam_test/internal/repository/clicks"
	"kadam_test/internal/service/clicker"
)

type Handlers struct {
	ClickerService   *clicker.Service
	ClicksRepository *clicks.Repository
}
