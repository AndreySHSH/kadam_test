package migration

type Click struct {
	Id           string  `gorm:"index:unique" json:"id,omitempty"`
	Ua           string  `json:"ua,omitempty"`
	Cpc          float64 `json:"cpc,omitempty"`
	RedirectUri  string  `json:"redirect_uri,omitempty"`
	IsSuspicious bool    `gorm:"default:false" json:"is_suspicious"`
}
