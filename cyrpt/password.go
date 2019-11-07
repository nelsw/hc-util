package cyrpt

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewRandomPassword() string {
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(charset))]
	}
	if p, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost); err != nil {
		return err.Error()
	} else {
		return string(p)
	}
}

func HashAndSalt(password []byte) (string, error) {
	if hashAndSalt, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost); err != nil {
		return "", err
	} else {
		return string(hashAndSalt), nil
	}
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), plainPwd) == nil
}
