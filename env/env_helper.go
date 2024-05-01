package env

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

func loadEnv(structEnvironments ...interface{}) {
	godotenv.Load(".env")

	for _, structEnvironment := range structEnvironments {
		elements := reflect.ValueOf(structEnvironment).Elem()
		environments := make(map[string]string)

		elementType := elements.Type()
		for i := 0; i < elementType.NumField(); i++ {
			field := elementType.Field(i)
			env := getEnvFromTagOrStructField(field)

			environments[field.Name] = env
		}

		parsed := parseMapToJSON(environments)
		json.Unmarshal([]byte(parsed), &structEnvironment)
	}
}

func getEnvFromTagOrStructField(field reflect.StructField) string {
	tagNameOnEnv := field.Tag.Get("env")

	if len(tagNameOnEnv) == 0 {
		return os.Getenv(field.Name)
	}
	return os.Getenv(tagNameOnEnv)
}

func parseMapToJSON(mp map[string]string) string {
	bytesOfMap, _ := json.Marshal(mp)
	return string(bytesOfMap)
}
