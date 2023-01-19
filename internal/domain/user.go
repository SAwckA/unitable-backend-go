package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	Username        string             `bson:"username"`
	Password        string             `bson:"password"`
	Email           string             `bson:"email"`
	Verified        bool               `bson:"verified"`
	Activated       bool               `bson:"activated"`
	TwoFAcode       int                `bson:"2fa_code"`
	EmailVerifyCode string             `bson:"email_verify_code"`
	Profile         UserProfile        `bson:"profile"`
	CreatedAt       time.Time          `bson:"CreatedAt"`
	LastVisit       time.Time          `bson:"LastVisit"`
}

func (i *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

type UserProfile struct {
	FullName string           `bson:"full_name" json:"full_name"`
	Contacts []ProfileContact `bson:"contacts" json:"contacts"`
}

type ProfileContact struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name" json:"name"`
	Value string             `bson:"value" json:"value"`
}

func NewProfileContact(name, value string) *ProfileContact {
	return &ProfileContact{
		ID:    primitive.NewObjectID(),
		Name:  name,
		Value: value,
	}
}

func NewUser(username string, password string, email string) *User {
	return &User{
		ID:              primitive.NewObjectID(),
		Username:        username,
		Email:           email,
		Password:        password,
		Verified:        false,
		Activated:       false,
		TwoFAcode:       0,
		EmailVerifyCode: uuid.NewString(),
		CreatedAt:       time.Now(),
		LastVisit:       time.Now(),
	}
}

func (u *User) UpdateLastVisit() {
	u.LastVisit = time.Now()
}
func (u *User) Set2FACode(value int) {
	u.TwoFAcode = value
}
func (u *User) SetEmailVerifyCode(value string) {
	u.EmailVerifyCode = value
}
func (u *User) SetUserVerification(value bool) {
	u.Verified = value
}
func (u *User) SetUserActivation(value bool) {
	u.Activated = value
}
