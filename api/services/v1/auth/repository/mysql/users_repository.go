package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
)

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) auth.UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (repo *usersRepository) GetUserByEmail(email *domain.Email) (*auth.Users, error, domain.RepositoryErrorType) {

	var err error
	var queryString string = `
		SELECT 
			id, locale, first_login
		FROM 
			users 
		WHERE 
			email=?
	`

	var queryResult *sql.Row
	var user *auth.Users = &auth.Users{
		Email: email,
	}

	queryResult = repo.db.QueryRow(queryString, email)

	err = queryResult.Scan(
		&user.ID,
		&user.Locale,
		&user.FirstLogin,
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

func (repo *usersRepository) CreateUser(user *auth.Users) (error, domain.RepositoryErrorType) {

	var err error
	var queryString string

	queryString = `
	INSERT INTO users( id, email, locale, first_login )
	VALUES( ?, ?, ?, ? )
	`

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	var res sql.Result

	res, err = statement.Exec(
		user.ID,
		user.Email,
		user.Locale,
		user.FirstLogin,
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
		return fmt.Errorf("Failed to Save User Data"), domain.RepositoryCreateDataFailed
	}

	return nil, 0

}

func (repo *usersRepository) UpdateUser(user *auth.Users) (error, domain.RepositoryErrorType) {

	var err error
	var queryString string

	queryString = `
	UPDATE users
	SET locale=?, first_login=?, updated_at=NOW()
	WHERE
	`

	if user.ID != nil {
		queryString += " id=?"
	} else if user.Email != nil {
		queryString += " email=?"
	}

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	var res sql.Result

	if user.ID != nil {
		res, err = statement.Exec(
			user.Locale,
			user.FirstLogin,
			user.ID,
		)
	} else if user.Email != nil {
		res, err = statement.Exec(
			user.Locale,
			user.FirstLogin,
			user.Email,
		)
	}

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
		return fmt.Errorf("Failed to Insert New User"), domain.RepositoryUpdateDataFailed
	}

	return nil, 0

}
