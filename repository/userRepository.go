package repository

import (
	"database/sql"
	"log"

	"github.com/moficodes/go-web-socket-example/models"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (user *User) GetID() string {
	return user.ID
}

func (user *User) GetName() string {
	return user.Name
}

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) AddUser(user models.User) {
	stmt, err := repo.DB.Prepare("INSERT INTO user(id, name) values(?,?)")
	checkErr(err)

	_, err = stmt.Exec(user.GetID(), user.GetName())
	checkErr(err)
}

func (repo *UserRepository) RemoveUser(user models.User) {
	stmt, err := repo.DB.Prepare("DELETE FROM user WHERE id = ?")
	checkErr(err)

	_, err = stmt.Exec(user.GetID())
	checkErr(err)
}

func (repo *UserRepository) FindUserByID(ID string) models.User {
	row := repo.DB.QueryRow("SELECT id, name FROM user WHERE id = ? LIMIT 1", ID)
	var user User

	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return &user
}

func (repo *UserRepository) GetAllUsers() []models.User {
	rows, err := repo.DB.Query("SELECT id, name FROM user")

	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, &user)
	}

	return users
}
