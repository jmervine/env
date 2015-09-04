package fixtures

import (
	"fmt"
	"os"
	"time"
)

var Fixtures = make(map[string]interface{})

func init() {
	Fixtures["F_STRING"] = "string"
	Fixtures["F_BYTES"] = []byte{'b', 'y', 't', 'e', 's'}
	Fixtures["F_INT"] = int(1)
	Fixtures["F_INT32"] = int32(1)
	Fixtures["F_INT64"] = int64(1)
	Fixtures["F_FLOAT32"] = float32(1)
	Fixtures["F_FLOAT64"] = float64(1)
	Fixtures["F_BOOL"] = bool(true)

	d, _ := time.ParseDuration("1h")
	Fixtures["F_DURATION"] = d
}

func UnsetFixtures() {
	for key, _ := range Fixtures {
		os.Unsetenv(key)
	}

	// and others, sometimes
	os.Unsetenv("F_BLANK")
	os.Unsetenv("F_WHITESPACE")
}

func SetFixtures() {
	for key, val := range Fixtures {
		if key == "F_BYTES" {
			os.Setenv(key, string(val.([]byte)))
			continue
		}

		os.Setenv(key, fmt.Sprintf("%v", val))
	}
}

func ResetFixtures() {
	UnsetFixtures()
	SetFixtures()
}
