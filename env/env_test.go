package env

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnvTestSuite struct {
	suite.Suite
}

func (suite *EnvTestSuite) SetupTest() {
	// Set up a sample .env file for testing
	os.Setenv("FOO", "TestFooValue")
	os.Setenv("BAR", "TestBarValue")
}

func (suite *EnvTestSuite) TestLoadEnv() {
	// Define a struct for testing
	type TestStruct struct {
		Foo string `env:"FOO"`
		Bar string `env:"BAR"`
	}

	// Create an instance of the struct
	testStruct := &TestStruct{}

	// Call the loadEnv function with the test struct
	loadEnv(testStruct)

	// Check if the values are loaded correctly
	expectedFoo := os.Getenv("FOO")
	expectedBar := os.Getenv("BAR")

	suite.Equal(expectedFoo, testStruct.Foo, "Foo values do not match")
	suite.Equal(expectedBar, testStruct.Bar, "Bar values do not match")
}

func (suite *EnvTestSuite) TestGetEnvFromTagOrStructField() {
	// Define a struct field for testing
	type TestStruct struct {
		Foo string `env:"FOO"`
	}

	field := reflect.TypeOf(TestStruct{}).Field(0)

	// Test when there is a tag
	envValue := getEnvFromTagOrStructField(field)
	expectedEnvValue := os.Getenv("FOO")

	suite.Equal(expectedEnvValue, envValue, "Env values do not match")

	// Test when there is no tag
	fieldWithoutTag := reflect.TypeOf(TestStruct{}).Field(0)
	fieldWithoutTag.Tag = reflect.StructTag("")

	envValueWithoutTag := getEnvFromTagOrStructField(fieldWithoutTag)
	expectedEnvValueWithoutTag := os.Getenv("Foo")

	suite.Equal(expectedEnvValueWithoutTag, envValueWithoutTag, "Env values without tag do not match")
}

func (suite *EnvTestSuite) TestParseMapToJSON() {
	// Test the parseMapToJSON function with a sample map
	testMap := map[string]string{"Key1": "Value1", "Key2": "Value2"}
	expectedJSON := `{"Key1":"Value1","Key2":"Value2"}`

	jsonResult := parseMapToJSON(testMap)

	suite.Equal(expectedJSON, jsonResult, "JSON values do not match")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}
