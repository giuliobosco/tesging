package mock

import "github.com/giuliobosco/tesging/model"

// userList, mock user list
var userList = []model.User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
	{Username: "user3", Password: "pass3"},
}

// GetUserList gets the mock user list
func GetUserList() *[]model.User {
	return &userList
}

// AppendUser appends a user to the mock user list
func AppendUser(u model.User) *model.User {
	userList = append(userList, u)

	return &u
}
