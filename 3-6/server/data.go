package main

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]User{
	"1": {
		Id:    "1",
		Name:  "11",
		Phone: "11111111",
	},
	"2": {
		Id:    "2",
		Name:  "22",
		Phone: "22222222",
	},
}
