package fixtures

import (
	"fmt"
	"os"
)

var Fixtures = map[string]interface{}{
	"F_STRING":  "string",
	"F_INT":     int(1),
	"F_INT32":   int32(1),
	"F_INT64":   int64(1),
	"F_UINT":    uint(1),
	"F_UINT64":  uint64(1),
	"F_FLOAT32": float32(1),
	"F_FLOAT64": float64(1),
	"F_BOOL":    bool(true),
}

func UnSetFixtures() {
	for key, _ := range Fixtures {
		os.Unsetenv(key)
	}

	// and others, sometimes
	os.Unsetenv("F_BLANK")
	os.Unsetenv("F_WHITESPACE")
}

func SetFixtures() {
	for key, val := range Fixtures {
		os.Setenv(key, fmt.Sprintf("%v", val))
	}
}
