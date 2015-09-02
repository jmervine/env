package envcfg

import (
	. "github.com/jmervine/envcfg/_fixtures"

	"fmt"
	"log"
	"os"
	"testing"
	"time"

	. "github.com/jmervine/GoT"
)

var env = "_fixtures/fixtures.env"

func init() {
	UnSetFixtures()

	d, _ := time.ParseDuration("1h")
	Fixtures["F_DURATION"] = d
}

func TestLoad(T *testing.T) {
	defer UnSetFixtures()

	os.Setenv("F_INT", "999")

	err := Load(env)
	Go(T).AssertNil(err)

	// ensure no clobber
	i := int(999)
	Go(T).AssertEqual(GetInt("F_INT"), &i)

	// specials
	Go(T).AssertEqual(get("F_BLANK"), "")
	Go(T).AssertEqual(get("F_WHITESPACE"), "whitespace")
	Go(T).AssertNil(Get("#F_COMMENT"))

	// standard
	for key, _ := range Fixtures {
		Go(T).RefuteNil(Get(key))
	}

	// check a few at random
	u := uint64(9)
	Go(T).AssertEqual(GetUint64("F_UINT64"), &u)

	d, _ := time.ParseDuration("1h1m1s")
	Go(T).AssertEqual(GetDuration("F_DURATION"), &d)

	b := false
	Go(T).AssertEqual(GetBool("F_BOOL"), &b)
}

func TestSet(T *testing.T) {
	defer UnSetFixtures()

	for key, val := range Fixtures {
		Set(key, val)

		Go(T).AssertEqual(os.Getenv(key), fmt.Sprintf("%v", val))
	}
}

func TestSetMap(T *testing.T) {
	defer UnSetFixtures()
	SetMap(Fixtures)

	for key, val := range Fixtures {
		Go(T).AssertEqual(os.Getenv(key), fmt.Sprintf("%v", val))
	}
}

func TestIsSet(T *testing.T) {
	defer UnSetFixtures()
	SetFixtures()

	Go(T).Assert(IsSet("F_STRING"))
}

func TestGet(T *testing.T) {
	defer UnSetFixtures()

	s := Get("F_STRING")
	Go(T).AssertNil(s)

	SetFixtures()

	s = Get("F_STRING")
	Go(T).AssertEqual(*s, "string")

	s = GetString("F_STRING")
	Go(T).AssertEqual(*s, "string")
}

func TestRequire(T *testing.T) {
	defer UnSetFixtures()

	s, e := Require("F_STRING")
	Go(T).AssertEqual(e.Error(), "missing required string from F_STRING")
	Go(T).AssertNil(s)

	SetFixtures()

	s, e = Require("F_STRING")
	Go(T).AssertEqual(*s, "string")
	Go(T).AssertNil(e)

	s, e = RequireString("F_STRING")
	Go(T).AssertEqual(*s, "string")
	Go(T).AssertNil(e)
}

func TestGetOrSet(T *testing.T) {
	defer UnSetFixtures()

	s := GetOrSet("F_STRING", "default")
	Go(T).AssertEqual(*s, "default")
	UnSetFixtures()

	SetFixtures()
	s = GetOrSet("F_STRING", "default")
	Go(T).AssertEqual(*s, "string")

	UnSetFixtures()
	s = GetOrSetString("F_STRING", "default")
	Go(T).AssertEqual(*s, "default")
	UnSetFixtures()

	SetFixtures()
	s = GetOrSetString("F_STRING", "default")
	Go(T).AssertEqual(*s, "string")
}

func TestRequireDuration(T *testing.T) {
	defer UnSetFixtures()

	d, e := RequireDuration("F_DURATION")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(d)

	SetFixtures()
	d, e = RequireDuration("F_DURATION")
	Go(T).AssertEqual(*d, Fixtures["F_DURATION"])
	Go(T).AssertNil(e)
}

func TestGetDuration(T *testing.T) {
	defer UnSetFixtures()
	SetFixtures()

	d := GetDuration("F_DURATION")
	Go(T).AssertEqual(*d, Fixtures["F_DURATION"])
}

func TestGetOrSetDuration(T *testing.T) {
	defer UnSetFixtures()

	def, _ := time.ParseDuration("1d")

	d := GetOrSetDuration("F_DURATION", def)
	Go(T).AssertEqual(*d, def)

	SetFixtures()

	d = GetOrSetDuration("F_DURATION", def)
	Go(T).AssertEqual(*d, Fixtures["F_DURATION"])
}

func TestGetInt(T *testing.T) {
	defer UnSetFixtures()

	i := GetInt("F_INT")
	Go(T).AssertNil(i)

	SetFixtures()

	i = GetInt("F_INT")
	Go(T).AssertEqual(*i, Fixtures["F_INT"])
}

func TestGetOrSetInt(T *testing.T) {
	defer UnSetFixtures()

	def := int(2)

	i := GetOrSetInt("F_INT", def)
	Go(T).AssertEqual(*i, int(2))

	SetFixtures()

	i = GetInt("F_INT")
	Go(T).AssertEqual(*i, Fixtures["F_INT"])
}

func TestRequireInt(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireInt("F_INT")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	SetFixtures()
	i, e = RequireInt("F_INT")
	Go(T).AssertEqual(*i, Fixtures["F_INT"])
	Go(T).AssertNil(e)
}

func TestGetUint(T *testing.T) {
	defer UnSetFixtures()

	i := GetUint("F_UINT")
	Go(T).AssertNil(i)

	SetFixtures()

	i = GetUint("F_UINT")
	Go(T).AssertEqual(*i, Fixtures["F_UINT"])
}

func TestGetOrSetUint(T *testing.T) {
	defer UnSetFixtures()

	def := uint(2)

	i := GetOrSetUint("F_UINT", def)
	Go(T).AssertEqual(*i, uint(2))

	SetFixtures()

	i = GetUint("F_UINT")
	Go(T).AssertEqual(*i, Fixtures["F_UINT"])
}

func TestRequireUint(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireUint("F_UINT")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	SetFixtures()
	i, e = RequireUint("F_UINT")
	Go(T).AssertEqual(*i, Fixtures["F_UINT"])
	Go(T).AssertNil(e)
}

func TestGetUint64(T *testing.T) {
	defer UnSetFixtures()

	i := GetUint64("F_UINT64")
	Go(T).AssertNil(i)

	SetFixtures()

	i = GetUint64("F_UINT64")
	Go(T).AssertEqual(*i, Fixtures["F_UINT64"])
}

func TestGetOrSetUint64(T *testing.T) {
	defer UnSetFixtures()

	def := uint64(2)

	i := GetOrSetUint64("F_UINT64", def)
	Go(T).AssertEqual(*i, uint64(2))

	SetFixtures()

	i = GetUint64("F_UINT64")
	Go(T).AssertEqual(*i, Fixtures["F_UINT64"])
}

func TestRequireUint64(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireUint64("F_UINT64")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	SetFixtures()
	i, e = RequireUint64("F_UINT64")
	Go(T).AssertEqual(*i, Fixtures["F_UINT64"])
	Go(T).AssertNil(e)
}

func TestGetInt32(T *testing.T) {
	defer UnSetFixtures()

	i := GetInt32("F_INT32")
	Go(T).AssertNil(i)

	SetFixtures()

	i = GetInt32("F_INT32")
	Go(T).AssertEqual(*i, Fixtures["F_INT32"])
}

func TestGetOrSetInt32(T *testing.T) {
	defer UnSetFixtures()

	def := int32(2)

	i := GetOrSetInt32("F_INT32", def)
	Go(T).AssertEqual(*i, int32(2))

	SetFixtures()

	i = GetInt32("F_INT32")
	Go(T).AssertEqual(*i, Fixtures["F_INT32"])
}

func TestRequireInt32(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireInt32("F_INT32")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	SetFixtures()
	i, e = RequireInt32("F_INT32")
	Go(T).AssertEqual(*i, Fixtures["F_INT32"])
	Go(T).AssertNil(e)
}

func TestGetInt64(T *testing.T) {
	defer UnSetFixtures()

	i := GetInt64("F_INT64")
	Go(T).AssertNil(i)

	SetFixtures()

	i = GetInt64("F_INT64")
	Go(T).AssertEqual(*i, Fixtures["F_INT64"])
}

func TestGetOrSetInt64(T *testing.T) {
	defer UnSetFixtures()

	def := int64(2)

	i := GetOrSetInt64("F_INT64", def)
	Go(T).AssertEqual(*i, int64(2))

	SetFixtures()

	i = GetInt64("F_INT64")
	Go(T).AssertEqual(*i, Fixtures["F_INT64"])
}

func TestRequireInt64(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireInt64("F_INT64")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	SetFixtures()
	i, e = RequireInt64("F_INT64")
	Go(T).AssertEqual(*i, Fixtures["F_INT64"])
	Go(T).AssertNil(e)
}

func TestGetFloat32(T *testing.T) {
	defer UnSetFixtures()

	i := GetFloat32("F_FLOAT32")
	Go(T).AssertNil(i)

	SetFixtures()

	i = GetFloat32("F_FLOAT32")
	Go(T).AssertEqual(*i, Fixtures["F_FLOAT32"])
}

func TestGetOrSetFloat32(T *testing.T) {
	defer UnSetFixtures()

	def := float32(2)

	i := GetOrSetFloat32("F_FLOAT32", def)
	Go(T).AssertEqual(*i, float32(2))

	SetFixtures()

	i = GetFloat32("F_FLOAT32")
	Go(T).AssertEqual(*i, Fixtures["F_FLOAT32"])
}

func TestRequireFloat32(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireFloat32("F_FLOAT32")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	SetFixtures()
	i, e = RequireFloat32("F_FLOAT32")
	Go(T).AssertEqual(*i, Fixtures["F_FLOAT32"])
	Go(T).AssertNil(e)
}

func TestGetFloat64(T *testing.T) {
	defer UnSetFixtures()

	i := GetFloat64("F_FLOAT64")
	Go(T).AssertNil(i)

	SetFixtures()

	i = GetFloat64("F_FLOAT64")
	Go(T).AssertEqual(*i, Fixtures["F_FLOAT64"])
}

func TestGetOrSetFloat64(T *testing.T) {
	defer UnSetFixtures()

	def := float64(2)

	i := GetOrSetFloat64("F_FLOAT64", def)
	Go(T).AssertEqual(*i, float64(2))

	SetFixtures()

	i = GetFloat64("F_FLOAT64")
	Go(T).AssertEqual(*i, Fixtures["F_FLOAT64"])
}

func TestRequireFloat64(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireFloat64("F_FLOAT64")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	SetFixtures()
	i, e = RequireFloat64("F_FLOAT64")
	Go(T).AssertEqual(*i, Fixtures["F_FLOAT64"])
	Go(T).AssertNil(e)
}

func TestGetBool(T *testing.T) {
	defer UnSetFixtures()

	i := GetBool("F_BOOL")
	Go(T).AssertNil(i)

	SetFixtures()

	i = GetBool("F_BOOL")
	Go(T).AssertEqual(*i, Fixtures["F_BOOL"])
}

func TestGetOrSetBool(T *testing.T) {
	defer UnSetFixtures()

	def := bool(false)

	i := GetOrSetBool("F_BOOL", def)
	Go(T).AssertEqual(*i, bool(false))

	SetFixtures()

	i = GetBool("F_BOOL")
	Go(T).AssertEqual(*i, Fixtures["F_BOOL"])
}

func TestRequireBool(T *testing.T) {
	defer UnSetFixtures()

	i, e := RequireBool("F_BOOL")
	Go(T).RefuteNil(e)
	Go(T).AssertNil(i)

	SetFixtures()
	i, e = RequireBool("F_BOOL")
	Go(T).AssertEqual(*i, Fixtures["F_BOOL"])
	Go(T).AssertNil(e)
}

func Test_empty(T *testing.T) {
	s := ""
	Go(T).Assert(empty(&s))

	s = "not"
	Go(T).Refute(empty(&s))
}

func Test_get(T *testing.T) {
	defer UnSetFixtures()
	SetFixtures()

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

func Test_removeQuotes(T *testing.T) {
	s := `val`
	Go(T).AssertEqual(removeQuotes(s), "val")

	s = `"val"`
	Go(T).AssertEqual(removeQuotes(s), "val")

	s = `\"val\"`
	Go(T).AssertEqual(removeQuotes(s), `\"val\"`)

	s = `"\"val\""`
	Go(T).AssertEqual(removeQuotes(s), `\"val\"`)
}

func Test_stripComments(T *testing.T) {
	s := `foo="bar"# comment`
	Go(T).AssertEqual(stripComments(s), `foo="bar"`)

	s = `foo=bar#comment`
	Go(T).AssertEqual(stripComments(s), `foo=bar`)

	s = `foo="#bar"#comment`
	Go(T).AssertEqual(stripComments(s), `foo="#bar"`)

	s = `foo="#bar"####comment####`
	Go(T).AssertEqual(stripComments(s), `foo="#bar"`)
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

func ExampleLoad() {
	defer UnSetFixtures()

	Set("F_STRING", "old_string")

	// using `_fixtures/fixtures.env`
	Load(env)

	var f = GetFloat32("F_FLOAT32")
	var s = GetString("F_STRING")
	var b = GetBool("F_BOOL")
	var i = GetInt("F_INT")

	fmt.Printf("F_FLOAT32 ::: %v\n", *f)
	fmt.Printf("F_STRING  ::: %v\n", *s)
	fmt.Printf("F_BOOL    ::: %v\n", *b)
	fmt.Printf("F_INT     ::: %v\n", *i)

	// Output:
	// F_FLOAT32 ::: 9.1
	// F_STRING  ::: old_string
	// F_BOOL    ::: false
	// F_INT     ::: 9
}

func ExampleOverload() {
	defer UnSetFixtures()

	Set("F_STRING", "old_string")

	// using `_fixtures/fixtures.env`
	Overload(env)

	var f = GetFloat32("F_FLOAT32")
	var s = GetString("F_STRING")
	var b = GetBool("F_BOOL")
	var i = GetInt("F_INT")

	fmt.Printf("F_FLOAT32 ::: %v\n", *f)
	fmt.Printf("F_STRING  ::: %v\n", *s)
	fmt.Printf("F_BOOL    ::: %v\n", *b)
	fmt.Printf("F_INT     ::: %v\n", *i)

	// Output:
	// F_FLOAT32 ::: 9.1
	// F_STRING  ::: sample file
	// F_BOOL    ::: false
	// F_INT     ::: 9
}
