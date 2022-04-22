package utils

import (
	"github.com/rogerscuall/gotm/adapters/db"
	"github.com/rogerscuall/gotm/ports"
)

func DbConnection(dbName string) (ports.DbPort, error) {
	var err error
	var dbAdapter ports.DbPort
	dbAdapter, err = db.NewAdapter(dbName)
	if err != nil {
		panic(err)
	}
	//dbAdapter.Sync()
	return dbAdapter, err
}
