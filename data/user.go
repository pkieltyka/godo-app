package data

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"

	"gopkg.in/mgo.v2/bson"
	"upper.io/db"
)

var UserCollection db.Collection

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"-"`
}

func NewUser() *User {
	return &User{}
}

func FindUser(cond db.Cond) (user *User, err error) {
	res := UserCollection.Find(cond)
	err = res.One(&user)
	return
}

func FindUserById(id bson.ObjectId) (user *User, err error) {
	return FindUser(db.Cond{"_id": id})
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username empty")
	}
	if u.Password == "" {
		return errors.New("password empty")
	}
	// TODO: Invalidate duplicate usernames. Leave it to the database
	// to maintain a unique index.
	return nil
}

func (u *User) SetPassword(password string) {
	u.Password = u.HashPassword(password)
}

func (u *User) HashPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	return hex.EncodeToString(h.Sum(nil))
}
