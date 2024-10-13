package db

import (
	"gorm.io/gorm"
)

func CreateStudent(fname, lname, coll, cour string) *Student {
	return &Student{
		first_name: fname,
		last_name:  lname,
		college:    coll,
		course:     cour,
		followers:  0,
		following:  1,
	}
}

var db, err = gorm.Open()

func MigrateDB() {
	if err != nil {
		panic("Failed to connect to te database :(")
	}

	db.AutoMigrate(&Student{})
}

// func GetStudent(fname string) *Student {

// }

// func GetAllStudents() []Student {

// }
