package cyrpt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPassword = []byte("testpassword")
var testPasswordHashAndSalt = "$2a$04$VeMm0sjfZyHstGBBncyQBupW4J6DZ9rdM/b9WAG8Bo7OTuCq885iS"

func TestNewRandomPassword(t *testing.T) {
	p := NewRandomPassword()
	fmt.Println(p)
}

func TestHashAndSalt1(t *testing.T) {
	str, err := HashAndSalt(testPassword)
	if err != nil {
		t.Errorf("could not hash and salt test password string. details: %v", err)
	}
	fmt.Println(str)
}

func TestComparePasswords(t *testing.T) {
	b := ComparePasswords(testPasswordHashAndSalt, []byte(testPassword))
	assert.True(t, b)
}
