package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)
import "github.com/pkg/errors"

type User struct {
	Id        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func main() {
	user, err := UserDTO(2)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("User not found")
		}
	}
	log.Println(user)
}

func UserDTO(userId int64) (*User, error) {
	db, err := sqlx.Connect("mysql", "root:PXDN93VRKUm8TeE7@(localhost:33069)/jike_practice?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		log.Fatalln(err)
	}
	user := User{}

	err = db.Get(&user, "select * from user where id = ? limit 1", userId)
	if err != nil && err == sql.ErrNoRows {
		return nil, errors.Wrapf(err, "Record not found for user id = %d.", userId)
	}
	if err != nil {
		return nil, errors.Wrap(err, "Unknown error")
	}

	return &user, nil
}
