package db

type Student struct {
	id         int
	first_name string
	last_name  string
	college    string
	course     string
	followers  uint32
	following  uint32
}

func (s *Student) AddFollower() {
	s.followers++
}

func (s *Student) AddFollowing() {
	s.following++
}
