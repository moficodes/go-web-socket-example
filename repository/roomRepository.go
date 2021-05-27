package repository

import (
	"database/sql"

	"github.com/moficodes/go-web-socket-example/models"
)

type Room struct {
	ID      string
	Name    string
	Private bool
}

func (room *Room) GetID() string {
	return room.ID
}

func (room *Room) GetName() string {
	return room.Name
}

func (room *Room) GetPrivate() bool {
	return room.Private
}

type RoomRepository struct {
	DB *sql.DB
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func (repo *RoomRepository) AddRoom(room models.Room) {
	stmt, err := repo.DB.Prepare("INSERT INTO room(id, name, private) values(?,?,?)")
	checkErr(err)
	_, err = stmt.Exec(room.GetID(), room.GetName(), room.GetPrivate())
	checkErr(err)
}

func (repo *RoomRepository) FindRoomByName(name string) models.Room {
	row := repo.DB.QueryRow("SELECT id, name, private FROM room where name = ? LIMIT 1", name)
	var room Room
	if err := row.Scan(&room.ID, &room.Name, &room.Private); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}

	return &room
}
