package cyrpt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var expected = "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5jb20iLCJleHAiOjE2MDQ3MzIwMzJ9.5e-CeihchJl6HC3DXlDOafo3MxvHHuZ9AIhKA32j-tk; Expires=Sat, 07 Nov 2020 06:53:52 GMT"
var expired = time.Date(2019, 11, 7, 6, 53, 52, 00, time.UTC)
var expiry = time.Date(2020, 11, 7, 6, 53, 52, 00, time.UTC)
var email = "test@test.com"

func TestBuild(t *testing.T) {
	actual := New().Expiry(expiry).Email(email).Build()
	assert.Equal(t, expected, actual)
}

func TestExpired(t *testing.T) {
	cookie := New().Email(email).Expiry(expired).Build()
	_, err := Validate(cookie)
	if err != nil {
		assert.Error(t, err)
	}
}

func TestValidate(t *testing.T) {
	cookie := New().Email(email).Expiry(time.Time{}).Build()
	claims, err := Validate(cookie)
	if err != nil {
		t.Errorf("could not validate test token string. details: %v", err)
	}
	e := claims.Email
	assert.Equal(t, email, e)
}
