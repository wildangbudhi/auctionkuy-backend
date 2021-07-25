package mysql

import (
	"database/sql"
	"fmt"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
)

type banksRepository struct {
	db *sql.DB
}

func NewBanksRepository(db *sql.DB) assets.BanksRepository {
	return &banksRepository{
		db: db,
	}
}

func (repo *banksRepository) FetchBanks() ([]assets.Banks, error, domain.RepositoryErrorType) {

	var err error
	var queryString string = `
	SELECT
		b.id,
		b.name,
		b.icon_url
	FROM
		banks b 
	`

	var queryResult *sql.Rows
	var banks []assets.Banks = make([]assets.Banks, 0)

	queryResult, err = repo.db.Query(queryString)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	defer queryResult.Close()

	for queryResult.Next() {

		var setting assets.Banks = assets.Banks{}

		err = queryResult.Scan(
			&setting.ID,
			&setting.Name,
			&setting.IconURL,
		)

		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
		}

		banks = append(banks, setting)

	}

	return banks, nil, 0

}
