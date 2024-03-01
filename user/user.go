package user

type user struct {
	email     string
	name      string
	last_name string
	pswdhash  string
}

var userList = []user{}

func GetUserObject(email string) (user, bool) {
	// TODO use database
	for _, user := range userList {
		if user.email == email {
			return user, true
		}
	}
	return user{}, false
}

func (u *user) ValidatePasswordHash(pswdhash string) bool {
	return u.pswdhash == pswdhash
}

// func AddUserObjedct(user *user) bool {
// 	// declare new user object
// 	newUser := user{
// 		email:        user.email,
// 		passwordhash: user.pswdhash,
// 		name:         user.name,
// 		lastName:     user.last_name,
// 	}

// 	for _, ele := range userList {
// 		if ele.email == user.email {
// 			return false
// 		}
// 	}

// 	userList = append(userList, newUser)
// }
