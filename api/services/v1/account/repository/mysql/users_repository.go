package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/account"
)

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) account.UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (repo *usersRepository) GetUserByID(id *domain.UUID) (*account.Users, error, domain.RepositoryErrorType) {

	var err error
	var queryString string = `
	SELECT 
		u.email,
		u.name,
		u.phone,
		u.national_id_number,
		u.avatar_url,
		u.first_login,
		u.locale,
		u.bank_id,
		u.bank_account_id,
		u.bank_account_owner_name
	FROM 
		users u
	WHERE
		u.id = ?
	`

	var queryResult *sql.Row
	var user *account.Users = &account.Users{
		ID: id,
	}

	queryResult = repo.db.QueryRow(queryString, id)

	err = queryResult.Scan(
		&user.Email,
		&user.Name,
		&user.Phone,
		&user.NationalIDNumber,
		&user.AvatarURL,
		&user.FirstLogin,
		&user.Locale,
		&user.BankID,
		&user.BankAccountID,
		&user.BankAccountOwnerName,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err, domain.RepositoryDataNotFound
		}

		log.Println(err)
		return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	return user, nil, 0

}
