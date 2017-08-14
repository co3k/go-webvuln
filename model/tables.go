package model

type User struct {
	Id       int
	Username string
}

type Session struct {
	SessionId string
	UserId    int
}

type Activity struct {
	Id        int
	UserId    int
	Username  string
	Body      string
	CreatedAt string
}
