package envcfg

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	. "github.com/jmervine/GoT"
)

var fixtures = map[string]interface{}{
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

func init() {
	unsetFixtures()

	d, _ := time.ParseDuration("1h")
	fixtures["F_DURATION"] = d
}

func unsetFixtures() {
	for key, _ := range fixtures {
		os.Unsetenv(key)
	}
}

func setFixtures() {
	for key, val := range fixtures {
		os.Setenv(key, fmt.Sprintf("%v", val))
	}
}

func TestSet(T *testing.T) {
	defer unsetFixtures()

	for key, val := range fixtures {
		Set(key, val)

		Go(T).AssertEqual(os.Getenv(key), fmt.Sprintf("%v", val))
	}
}

func TestSetMap(T *testing.T) {
	defer unsetFixtures()
	SetMap(fixtures)

	for key, val := range fixtures {
		Go(T).AssertEqual(os.Getenv(key), fmt.Sprintf("%v", val))
	}
}

func TestIsSet(T *testing.T) {
	defer unsetFixtures()
	setFixtures()

	Go(T).Assert(IsSet("F_STRING"))
}

func TestGet(T *testing.T) {
	defer unsetFixtures()

	s := Get("F_STRING")
	Go(T).AssertNil(s)

	setFixtures()

	s = Get("F_STRING")
	Go(T).AssertEqual(*s, "string")

	s = GetString("F_STRING")
	Go(T).AssertEqual(*s, "string")
}

func TestRequire(T *testing.T) {
	defer unsetFixtures()

	s, e := Require("F_STRING")
	Go(T).AssertEqual(e.Error(), "missing required string from F_STRING")
	Go(T).AssertNil(s)

	setFixtures()

	s, e = Require("F_STRING")
	Go(T).AssertEqual(*s, "string")
	Go(T).AssertNil(e)

	s, e = RequireString("F_STRING")
	Go(T).AssertEqual(*s, "string")
	Go(T).AssertNil(e)
}

func TestGetOrSet(T *testing.T) {
	defer unsetFixtures()

	s := GetOrSet("F_STRING", "default")
	Go(T).AssertEqual(*s, "default")
	unsetFixtures()

	setFixtures()
	s = GetOrSet("F_STRING", "default")
	Go(T).AssertEqual(*s, "string")

	unsetFixtures()
	s = GetOrSetString("F_STRING", "default")
	Go(T).AssertEqual(*s, "default")
	unsetFixtures()

	setFixtures()
	s = GetOrSetString("F_STRING", "default")
	Go(T).AssertEqual(*s, "string")
}

func TestRequireDuration(T *testing.T) {
	defer unsetFixtures()

	d, e := RequireDuration("F_DURATION")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(d)

	setFixtures()
	d, e = RequireDuration("F_DURATION")
	Go(T).AssertEqual(*d, fixtures["F_DURATION"])
	Go(T).AssertNil(e)
}

func TestGetDuration(T *testing.T) {
	defer unsetFixtures()
	setFixtures()

	d := GetDuration("F_DURATION")
	Go(T).AssertEqual(*d, fixtures["F_DURATION"])
}

func TestGetOrSetDuration(T *testing.T) {
	defer unsetFixtures()

	def, _ := time.ParseDuration("1d")

	d := GetOrSetDuration("F_DURATION", def)
	Go(T).AssertEqual(*d, def)

	setFixtures()

	d = GetOrSetDuration("F_DURATION", def)
	Go(T).AssertEqual(*d, fixtures["F_DURATION"])
}

func TestGetInt(T *testing.T) {
	defer unsetFixtures()

	i := GetInt("F_INT")
	Go(T).AssertNil(i)

	setFixtures()

	i = GetInt("F_INT")
	Go(T).AssertEqual(*i, fixtures["F_INT"])
}

func TestGetOrSetInt(T *testing.T) {
	defer unsetFixtures()

	def := int(2)

	i := GetOrSetInt("F_INT", def)
	Go(T).AssertEqual(*i, int(2))

	setFixtures()

	i = GetInt("F_INT")
	Go(T).AssertEqual(*i, fixtures["F_INT"])
}

func TestRequireInt(T *testing.T) {
	defer unsetFixtures()

	i, e := RequireInt("F_INT")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	setFixtures()
	i, e = RequireInt("F_INT")
	Go(T).AssertEqual(*i, fixtures["F_INT"])
	Go(T).AssertNil(e)
}

func TestGetUint(T *testing.T) {
	defer unsetFixtures()

	i := GetUint("F_UINT")
	Go(T).AssertNil(i)

	setFixtures()

	i = GetUint("F_UINT")
	Go(T).AssertEqual(*i, fixtures["F_UINT"])
}

func TestGetOrSetUint(T *testing.T) {
	defer unsetFixtures()

	def := uint(2)

	i := GetOrSetUint("F_UINT", def)
	Go(T).AssertEqual(*i, uint(2))

	setFixtures()

	i = GetUint("F_UINT")
	Go(T).AssertEqual(*i, fixtures["F_UINT"])
}

func TestRequireUint(T *testing.T) {
	defer unsetFixtures()

	i, e := RequireUint("F_UINT")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	setFixtures()
	i, e = RequireUint("F_UINT")
	Go(T).AssertEqual(*i, fixtures["F_UINT"])
	Go(T).AssertNil(e)
}

func TestGetUint64(T *testing.T) {
	defer unsetFixtures()

	i := GetUint64("F_UINT64")
	Go(T).AssertNil(i)

	setFixtures()

	i = GetUint64("F_UINT64")
	Go(T).AssertEqual(*i, fixtures["F_UINT64"])
}

func TestGetOrSetUint64(T *testing.T) {
	defer unsetFixtures()

	def := uint64(2)

	i := GetOrSetUint64("F_UINT64", def)
	Go(T).AssertEqual(*i, uint64(2))

	setFixtures()

	i = GetUint64("F_UINT64")
	Go(T).AssertEqual(*i, fixtures["F_UINT64"])
}

func TestRequireUint64(T *testing.T) {
	defer unsetFixtures()

	i, e := RequireUint64("F_UINT64")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	setFixtures()
	i, e = RequireUint64("F_UINT64")
	Go(T).AssertEqual(*i, fixtures["F_UINT64"])
	Go(T).AssertNil(e)
}

func TestGetInt32(T *testing.T) {
	defer unsetFixtures()

	i := GetInt32("F_INT32")
	Go(T).AssertNil(i)

	setFixtures()

	i = GetInt32("F_INT32")
	Go(T).AssertEqual(*i, fixtures["F_INT32"])
}

func TestGetOrSetInt32(T *testing.T) {
	defer unsetFixtures()

	def := int32(2)

	i := GetOrSetInt32("F_INT32", def)
	Go(T).AssertEqual(*i, int32(2))

	setFixtures()

	i = GetInt32("F_INT32")
	Go(T).AssertEqual(*i, fixtures["F_INT32"])
}

func TestRequireInt32(T *testing.T) {
	defer unsetFixtures()

	i, e := RequireInt32("F_INT32")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	setFixtures()
	i, e = RequireInt32("F_INT32")
	Go(T).AssertEqual(*i, fixtures["F_INT32"])
	Go(T).AssertNil(e)
}

func TestGetInt64(T *testing.T) {
	defer unsetFixtures()

	i := GetInt64("F_INT64")
	Go(T).AssertNil(i)

	setFixtures()

	i = GetInt64("F_INT64")
	Go(T).AssertEqual(*i, fixtures["F_INT64"])
}

func TestGetOrSetInt64(T *testing.T) {
	defer unsetFixtures()

	def := int64(2)

	i := GetOrSetInt64("F_INT64", def)
	Go(T).AssertEqual(*i, int64(2))

	setFixtures()

	i = GetInt64("F_INT64")
	Go(T).AssertEqual(*i, fixtures["F_INT64"])
}

func TestRequireInt64(T *testing.T) {
	defer unsetFixtures()

	i, e := RequireInt64("F_INT64")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	setFixtures()
	i, e = RequireInt64("F_INT64")
	Go(T).AssertEqual(*i, fixtures["F_INT64"])
	Go(T).AssertNil(e)
}

func TestGetFloat32(T *testing.T) {
	defer unsetFixtures()

	i := GetFloat32("F_FLOAT32")
	Go(T).AssertNil(i)

	setFixtures()

	i = GetFloat32("F_FLOAT32")
	Go(T).AssertEqual(*i, fixtures["F_FLOAT32"])
}

func TestGetOrSetFloat32(T *testing.T) {
	defer unsetFixtures()

	def := float32(2)

	i := GetOrSetFloat32("F_FLOAT32", def)
	Go(T).AssertEqual(*i, float32(2))

	setFixtures()

	i = GetFloat32("F_FLOAT32")
	Go(T).AssertEqual(*i, fixtures["F_FLOAT32"])
}

func TestRequireFloat32(T *testing.T) {
	defer unsetFixtures()

	i, e := RequireFloat32("F_FLOAT32")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	setFixtures()
	i, e = RequireFloat32("F_FLOAT32")
	Go(T).AssertEqual(*i, fixtures["F_FLOAT32"])
	Go(T).AssertNil(e)
}

func TestGetFloat64(T *testing.T) {
	defer unsetFixtures()

	i := GetFloat64("F_FLOAT64")
	Go(T).AssertNil(i)

	setFixtures()

	i = GetFloat64("F_FLOAT64")
	Go(T).AssertEqual(*i, fixtures["F_FLOAT64"])
}

func TestGetOrSetFloat64(T *testing.T) {
	defer unsetFixtures()

	def := float64(2)

	i := GetOrSetFloat64("F_FLOAT64", def)
	Go(T).AssertEqual(*i, float64(2))

	setFixtures()

	i = GetFloat64("F_FLOAT64")
	Go(T).AssertEqual(*i, fixtures["F_FLOAT64"])
}

func TestRequireFloat64(T *testing.T) {
	defer unsetFixtures()

	i, e := RequireFloat64("F_FLOAT64")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	setFixtures()
	i, e = RequireFloat64("F_FLOAT64")
	Go(T).AssertEqual(*i, fixtures["F_FLOAT64"])
	Go(T).AssertNil(e)
}

func TestGetBool(T *testing.T) {
	defer unsetFixtures()

	i := GetBool("F_BOOL")
	Go(T).AssertNil(i)

	setFixtures()

	i = GetBool("F_BOOL")
	Go(T).AssertEqual(*i, fixtures["F_BOOL"])
}

func TestGetOrSetBool(T *testing.T) {
	defer unsetFixtures()

	def := bool(false)

	i := GetOrSetBool("F_BOOL", def)
	Go(T).AssertEqual(*i, bool(false))

	setFixtures()

	i = GetBool("F_BOOL")
	Go(T).AssertEqual(*i, fixtures["F_BOOL"])
}

func TestRequireBool(T *testing.T) {
	defer unsetFixtures()

	i, e := RequireBool("F_BOOL")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	setFixtures()
	i, e = RequireBool("F_BOOL")
	Go(T).AssertEqual(*i, fixtures["F_BOOL"])
	Go(T).AssertNil(e)
}

func Test_empty(T *testing.T) {
	s := ""
	Go(T).Assert(empty(&s))

	s = "not"
	Go(T).Refute(empty(&s))
}

func Test_get(T *testing.T) {
	defer unsetFixtures()
	setFixtures()

	Go(T).AssertEqual(get("F_STRING"), "string")
	Go(T).AssertEqual(get("F_INT"), "1")
}

func Test_toString(T *testing.T) {
	Go(T).AssertEqual(toString(9), "9")
}

func Test_onError(T *testing.T) {
	Go(T).AssertNil(onError(nil))
	Go(T).RefuteNil(onError(fmt.Errorf("error")))
}

func Example() {
	PanicOnRequire = true
	Set("EX_PORT", 3000)

	var (
		addr  *string
		port  *int
		debug *bool
		err   error
	)

	port, err = RequireInt("EX_PORT")
	if err != nil {
		log.Fatal(err)
	}

	addr = GetString("EX_ADDR")
	if addr == nil {
		s := ""
		addr = &s
	}

	debug = GetOrSetBool("EX_DEBUG", false)

	fmt.Printf("addr=%v port=%v debug=%v", *addr, *port, *debug)
	// Output: addr= port=3000 debug=false
}

func ExampleSet() {
	Set("SOME_INT", 1)
}

func ExampleGet() {
	// where: SOME_INT=1
	i := Get("SOME_INT")

	if *i == "1" {
		fmt.Printf("%v", i)
	}
}

func ExampleGetInt() {
	// where: SOME_INT=1
	i := GetInt("SOME_INT")

	if *i == 1 {
		fmt.Printf("%v", i)
	}
}
