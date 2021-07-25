package mysql

import (
	"database/sql"
	"fmt"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
)

type appSettingsRepository struct {
	db *sql.DB
}

func NewAppSettingsRepository(db *sql.DB) assets.AppSettingsRepository {
	return &appSettingsRepository{
		db: db,
	}
}

func (repo *appSettingsRepository) FetchAppSettings() ([]assets.AppSettings, error, domain.RepositoryErrorType) {

	var err error
	var queryString string = `
	SELECT 
		ass.key,
		ass.value  
	FROM 
		app_settings ass
	`

	var queryResult *sql.Rows
	var appSettings []assets.AppSettings = make([]assets.AppSettings, 0)

	queryResult, err = repo.db.Query(queryString)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	defer queryResult.Close()

	for queryResult.Next() {

		var setting assets.AppSettings = assets.AppSettings{}

		err = queryResult.Scan(
			&setting.Key,
			&setting.Value,
		)

		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
		}

		appSettings = append(appSettings, setting)

	}

	return appSettings, nil, 0

}
