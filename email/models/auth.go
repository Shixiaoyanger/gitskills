package models

type User struct {
	Username string `json:"username"`
	Password string `josn:"password"`
}

func CheckAuth(user User) bool {
	_, passwd, err := Getinfo(user.Username)
	if err == false {
		return err
	}
	if passwd != Getkey(user.Password) {
		return false
	}
	return true

}
