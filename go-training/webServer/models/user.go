package models

type User struct {
	id        int
	firstName string
	lastName  string
}

var (
	users  []*User
	nextId int
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	user.id = nextId
	nextId++
	users = append(users, &user)
	return user, nil
}
