package models

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) ChangeAge(age int) {
	u.Age = age
}
