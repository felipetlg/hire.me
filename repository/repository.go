package repository

import (
	"log"

	"github.com/felipetlg/hire.me/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	databaseLocation = "repository/database.db"
)

type Repository interface {
	SetupDatabase() error
	InsertNewEntry(url *model.Url) error
	RetrieveLongUrl(alias string) (*model.Url, error)
	UpdateUrl(url *model.Url)
	MostVisited(n int) ([]model.Url, error)
}

type UrlRepository struct {
	Db *gorm.DB
}

func (repo *UrlRepository) SetupDatabase() error {

	// DEBUG database query
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log level
	// 		Colorful:      false,       // Disable color
	// 	},
	// )
	//db, err := gorm.Open(sqlite.Open(databaseLocation), &gorm.Config{Logger: newLogger})

	db, err := gorm.Open(sqlite.Open(databaseLocation), &gorm.Config{})

	if err != nil {
		return err
	}

	repo.Db = db
	return nil
}

func (repo *UrlRepository) InsertNewEntry(url *model.Url) error {
	result := repo.Db.Create(url)

	// Caso o alias já exista retorna erro e não insere nova entrada
	return result.Error
}

func (repo *UrlRepository) RetrieveLongUrl(alias string) (*model.Url, error) {
	var url model.Url
	result := repo.Db.Where("alias = ?", alias).First(&url)

	// Retorna erro em caso de não encontrar alias na base
	return &url, result.Error
}

func (repo *UrlRepository) UpdateUrl(url *model.Url) {
	err := repo.Db.Model(url).Where("alias = ?", url.Alias).Update("visits", url.Visits).Error
	if err != nil {
		log.Print(err)
	}
}

func (repo *UrlRepository) MostVisited(n int) ([]model.Url, error) {
	urls := make([]model.Url, 0)
	if err := repo.Db.
		Order("visits desc").
		Limit(n).
		Find(&urls).Error; err != nil {
		return nil, err
	}

	return urls, nil
}
