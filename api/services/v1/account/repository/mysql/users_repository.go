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

func (repo *usersRepository) GetUserByID(id *domain.UUID, imagePrefix string) (*account.Users, error, domain.RepositoryErrorType) {

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

	if user.AvatarURL != nil {
		user.AvatarURL.SetPrefix(imagePrefix)
	}

	return user, nil, 0

}

func (repo *usersRepository) UpdateUser(user *account.Users) (error, domain.RepositoryErrorType) {

	var err error
	var queryString string
	var queryKey interface{} = nil

	queryString = `
	UPDATE users
	SET name=?, phone=?, national_id_number=?, avatar_url=?, first_login=?, locale=?, bank_id=?, bank_account_id=?, bank_account_owner_name=?, updated_at=NOW()
	WHERE
	`

	if user.ID != nil {
		queryString += " id=?"
		queryKey = user.ID
	} else if user.Email != nil {
		queryString += " email=?"
		queryKey = user.Email
	}

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	var res sql.Result

	res, err = statement.Exec(
		user.Name,
		user.Phone,
		user.NationalIDNumber,
		user.AvatarURL,
		user.FirstLogin,
		user.Locale,
		user.BankID,
		user.BankAccountID,
		user.BankAccountOwnerName,
		queryKey,
	)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	if rowAffected == 0 {
		return fmt.Errorf("Failed to Insert Update User Data"), domain.RepositoryUpdateDataFailed
	}

	return nil, 0

}
