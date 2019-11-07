package aws

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var code = 200
var head = map[string]string{"Access-Control-Allow-Origin": "*"}
var text = string("bad command: test msg")
var data = map[string]string{"success": ""}
var toke = ""

func TestBuild(t *testing.T) {

	// read json from file
	inputJSON, err := ioutil.ReadFile("./testdata/response.json")
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	// de-serialize into Go object
	var expected events.APIGatewayProxyResponse
	if err := json.Unmarshal(inputJSON, &expected); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	// build actual test object
	actual, err := New().Head(head).Code(code).Text(text).Toke(toke).Data(data).Build()
	if err != nil {
		t.Errorf("could not build test response. details: %v", err)
	}

	assert.Equal(t, err, nil)
	assert.Equal(t, expected, actual)
}
