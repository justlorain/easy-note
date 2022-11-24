package pack

import (
	"easy-note/cmd/user/dal/db"
	"easy-note/kitex_gen/demouser"
)

// User pack user info
func User(u *db.User) *demouser.User {
	if u == nil {
		return nil
	}

	return &demouser.User{UserId: int64(u.ID), Username: u.Username, Avatar: "test"}
}

// Users pack list of user info
func Users(us []*db.User) []*demouser.User {
	users := make([]*demouser.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
