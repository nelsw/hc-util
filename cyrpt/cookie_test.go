package cyrpt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var expected = "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDQ3MzIwMzJ9.E9LGY5NnapOW4QsHClLBRhfz6-CvjL9mx0XgU66UtQM; Expires=Sat, 07 Nov 2020 06:53:52 GMT"
var expired = time.Date(2019, 11, 7, 6, 53, 52, 00, time.UTC)
var expiry = time.Date(2020, 11, 7, 6, 53, 52, 00, time.UTC)
var email = "test@test.com"

func TestBuild(t *testing.T) {
	actual := New().Expiry(expiry).Email(email).Build()
	assert.Equal(t, expected, actual)
}

func TestExpired(t *testing.T) {
	err := Validate(New().Email(email).Expiry(expired).Build(), &Claims{})
	if err != nil {
		assert.Error(t, err)
	}
}

func TestValidate(t *testing.T) {
	err := Validate(New().Email(email).Expiry(time.Time{}).Build(), &Claims{})
	if err != nil {
		t.Errorf("could not validate test token string. details: %v", err)
	}
}
