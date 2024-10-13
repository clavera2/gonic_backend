package service

import (
	"github.com/clavera2/gonic_backend/db"
	// "github.com/clavera2/gonic_backend/api"
)

func CreateStudent(fname, lname, coll, cour string) {
	db.CreateStudent(fname, lname, coll, cour)
}
