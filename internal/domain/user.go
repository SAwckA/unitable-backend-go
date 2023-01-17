package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	_id             primitive.ObjectID `bson:"_id"`
	Username        string             `bson:"username"`
	Password        string             `bson:"password"`
	Email           string             `bson:"email"`
	Verified        bool               `bson:"verified"`
	Activated       bool               `bson:"activated"`
	TwoFAcode       int                `bson:"2fa_code"`
	EmailVerifyCode string             `bson:"email_verify_code"`
	Profile         UserProfile        `bson:"profile"`
}

type UserProfile struct {
	FullName string            `bson:"full_name"`
	Contacts map[string]string `bson:"contacts"`
}
