package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/bertoxic/aarc/models"
)

// getUserByEmail retrieves a user's information from the database by email address
func(m *postgresDBRepo)GetUserByEmail(email string) (models.User, error) {
	// Retrieve the user's information from the database
	// ...

	return models.User{}, nil
}

// updateUser saves the updated user information in the database
func(m *postgresDBRepo)UpdateUser(user models.User) error {
	// Save the updated user information in the database
	// ...

	return nil
}

func(m *postgresDBRepo)CreateUser(user models.User) error {
	// Save the updated user information in the database
	// ...
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `--sql
	insert into users (name, email, password, verification, verification_expiry, created_at, updated_at) values ($1,$2, $3,$4,$5,$6,$7)
	`
	_, err := m.DB.ExecContext(ctx, stmt,user.Name,user.Email,user.Password,user.Verification,user.VerificationExpiry,time.Now(),time.Now())


	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func(m *postgresDBRepo)GetVerifiedUserByEmail(email string) (models.User, error) {
	// Save the updated user information in the database
	// ...
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	log.Println("zzzzzzzzzzzzz..",email,)
	stmt := `--sql
	select name, email, password, verification, verified , verification_expiry  from users where email = $1
	`
	row := m.DB.QueryRowContext(ctx, stmt,email)
	

	err:= row.Scan(
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Verification,
		&u.Verified,
		&u.VerificationExpiry,
	
	)

	if err != nil {
		log.Println(err)
		return u, err
	}
	log.Println("xxxxxxxxxxxcvvvv..",u.Email,u.Verification)
	return u, nil
}



func(m *postgresDBRepo) IsEmailUsed(email string) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var databaseMail string
	query := `--sql
	select  email from users where email = $1
	`
	row:= m.DB.QueryRowContext(ctx,query,email)

	row.Scan(&databaseMail)
	if err:=row.Err();err !=nil {
		log.Println(err)
		return true , err
	}

	if email == databaseMail {
		log.Println(email ,"==",databaseMail)
		return true, nil
	}
		return false, nil


}

func(m *postgresDBRepo) CheckifTableExist(email string) error{

sqlStatement := `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	verification TEXT,
	verification_expiry TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

// Execute the SQL statement to create the table
_, err := m.DB.Exec(sqlStatement)
return err

}