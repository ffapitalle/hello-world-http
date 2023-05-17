package db

import (
	"strconv"
)

func (db *Database) AddVisit(visitor string) (int, error) {

	visits := 0

	val, err := db.Client.Get(db.Client.Context(), visitor).Result()
	if err != nil {
		println(err.Error())
	} else {
		visits, _ = strconv.Atoi(val)
	}

	visits = visits + 1
	db.Client.Set(db.Client.Context(), visitor, strconv.Itoa(visits), 0)

	return visits, err
}
