package userModel

type User struct {
	ID        string `json:"id"`
	NameUser  string `json:"nameuser"`
	EmailUser string `json:"emailuser"`
	Password  string `json:"password"`
}
