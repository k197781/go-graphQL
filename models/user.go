package models

type User struct {
	Id   int
	Name string `sql:"size:255"`
}
