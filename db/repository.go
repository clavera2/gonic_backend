package db

func CreateStudent(fname, lname, coll, cour string) *Student {
	return &Student{
		id:         0,
		first_name: fname,
		last_name:  lname,
		college:    coll,
		course:     cour,
		followers:  0,
		following:  1,
	}
}
