package model

type Url struct {
	Alias    string `gorm:"column:alias" json:”alias”`
	LongUrl  string `gorm:"column:longUrl" json:”longUrl”`
	ShortUrl string `gorm:"column:shortUrl" json:”shortUrl”`
	Visits   int    `gorm:"column:visits" json:”visits”`
}
