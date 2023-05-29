package model

import (
	"errors"
	"html"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string  `json:"id"`
	UserName string  `json:"username"`
	Password string  `json:"password"`
	Entries  []Entry `json:"entries"`
}

var nameMap = make(map[string]User)
var idMap = make(map[string]*User)

func (user *User) Save() (*User, error) {
	userId := uuid.New().String()
	_, exist := nameMap[user.UserName]
	if exist {
		return &User{}, errors.New("user is existing")
	}

	newUser := User{ID: userId, UserName: user.UserName, Password: user.Password}
	nameMap[user.UserName] = newUser
	idMap[userId] = &newUser

	return &newUser, nil
}

func (user *User) BeforeSave() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.UserName = html.EscapeString(strings.TrimSpace(user.UserName))
	return nil
}

func GetUserByName(username string) (*User, error) {
	if username == "" {
		return &User{}, errors.New("must provide username")
	}
	u, ok := nameMap[username]
	if !ok {
		return &User{}, errors.New("cannot find user by username " + username)
	}
	return &u, nil
}

func GetUserById(id string) (*User, error) {
	if id == "" {
		return &User{}, errors.New("must provide user id")
	}
	u, ok := idMap[id]
	if !ok {
		return &User{}, errors.New("cannot find user by user id " + id)
	}
	return u, nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
