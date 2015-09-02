package envcfg

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// PanicOnRequire forces panics when Require- methods fail.
var PanicOnRequire = false

// Set sets via an interface
func Set(key string, val interface{}) {
	os.Setenv(key, toString(val))
}

// SetMap iterates over a map and sets keys to values
func SetMap(m map[string]interface{}) {
	for key, val := range m {
		Set(key, val)
	}
}

// IsSet return true if a key is not "" in os.Getenv
func IsSet(key string) bool {
	s := get(key)
	return !empty(&s)
}

// Get gets a key and returns a string
func Get(key string) *string {
	s := get(key)

	if s == "" {
		return nil
	}

	return &s
}

// GetString is an alias to Get
func GetString(key string) *string {
	return Get(key)
}

// Require gets a key and returns a string or an error if it's set to "" in
// os.Getenv
func Require(key string) (*string, error) {
	s := get(key)
	if empty(&s) {
		return nil, onError(fmt.Errorf("missing required string from %s", key))
	}
	return &s, nil
}

func RequireString(key string) (*string, error) {
	return Require(key)
}

// GetOrSet gets a key and returns a string or set's the default
func GetOrSet(key string, val interface{}) *string {
	if IsSet(key) {
		return Get(key)
	}
	Set(key, val)

	s := toString(val)
	return &s
}

// GetOrSetString is an alias to GetOrSet
func GetOrSetString(key, val string) *string {
	return GetOrSet(key, val)
}

func RequireDuration(key string) (*time.Duration, error) {
	s := get(key)

	d, e := time.ParseDuration(s)
	if e != nil {
		return nil, onError(e)
	}
	return &d, nil
}

func GetDuration(key string) *time.Duration {
	d, e := RequireDuration(key)
	if e != nil {
		return nil
	}
	return d
}

func GetOrSetDuration(key string, val time.Duration) *time.Duration {
	if IsSet(key) {
		return GetDuration(key)
	}
	Set(key, val)
	return &val
}

// GetInt gets a key and returns an int
func GetInt(key string) *int {
	if i, e := strconv.ParseInt(get(key), 10, 16); e != nil {
		return nil
	} else {
		n := int(i)
		return &n
	}
}

func GetOrSetInt(key string, val int) *int {
	if IsSet(key) {
		return GetInt(key)
	}
	Set(key, val)
	return &val
}

func RequireInt(key string) (*int, error) {
	b := GetInt(key)
	if b == nil {
		return nil, onError(fmt.Errorf("missing required int from %s", key))
	} else {
		return b, nil
	}
}

// GetUint gets a key and returns an uint
func GetUint(key string) *uint {
	if i, e := strconv.ParseInt(get(key), 10, 32); e != nil {
		return nil
	} else {
		n := uint(i)
		return &n
	}
}

func GetOrSetUint(key string, val uint) *uint {
	if IsSet(key) {
		return GetUint(key)
	}
	Set(key, val)
	return &val
}

func RequireUint(key string) (*uint, error) {
	b := GetUint(key)
	if b == nil {
		return nil, onError(fmt.Errorf("missing required uint from %s", key))
	} else {
		return b, nil
	}
}

// GetUint64 gets a key and returns an uint64
func GetUint64(key string) *uint64 {
	if i, e := strconv.ParseInt(get(key), 10, 64); e != nil {
		return nil
	} else {
		n := uint64(i)
		return &n
	}
}

func GetOrSetUint64(key string, val uint64) *uint64 {
	if IsSet(key) {
		return GetUint64(key)
	}
	Set(key, val)
	return &val
}

func RequireUint64(key string) (*uint64, error) {
	b := GetUint64(key)
	if b == nil {
		return nil, onError(fmt.Errorf("missing required uint64 from %s", key))
	} else {
		return b, nil
	}
}

// GetInt32 gets a key and returns an int32
func GetInt32(key string) *int32 {
	if i, e := strconv.ParseInt(get(key), 10, 32); e != nil {
		return nil
	} else {
		n := int32(i)
		return &n
	}
}

func GetOrSetInt32(key string, val int32) *int32 {
	if IsSet(key) {
		return GetInt32(key)
	}
	Set(key, val)
	return &val
}

func RequireInt32(key string) (*int32, error) {
	b := GetInt32(key)
	if b == nil {
		return nil, onError(fmt.Errorf("missing required int32 from %s", key))
	} else {
		return b, nil
	}
}

// GetInt64 gets a key and returns an int64
func GetInt64(key string) *int64 {
	if i, e := strconv.ParseInt(get(key), 10, 64); e != nil {
		return nil
	} else {
		return &i
	}
}

func GetOrSetInt64(key string, val int64) *int64 {
	if IsSet(key) {
		return GetInt64(key)
	}
	Set(key, val)
	return &val
}

func RequireInt64(key string) (*int64, error) {
	b := GetInt64(key)
	if b == nil {
		return nil, onError(fmt.Errorf("missing required int64 from %s", key))
	} else {
		return b, nil
	}
}

// GetFloat32 gets a key and returns an float32
func GetFloat32(key string) *float32 {
	if i, e := strconv.ParseFloat(get(key), 32); e != nil {
		return nil
	} else {
		n := float32(i)
		return &n
	}
}

func RequireFloat32(key string) (*float32, error) {
	b := GetFloat32(key)
	if b == nil {
		return nil, onError(fmt.Errorf("missing required float32 from %s", key))
	} else {
		return b, nil
	}
}

func GetOrSetFloat32(key string, val float32) *float32 {
	if IsSet(key) {
		return GetFloat32(key)
	}
	Set(key, val)
	return &val
}

// GetFloat64 gets a key and returns an float64
func GetFloat64(key string) *float64 {
	if i, e := strconv.ParseFloat(get(key), 64); e != nil {
		return nil
	} else {
		return &i
	}
}

func GetOrSetFloat64(key string, val float64) *float64 {
	if IsSet(key) {
		return GetFloat64(key)
	}
	Set(key, val)
	return &val
}

func RequireFloat64(key string) (*float64, error) {
	b := GetFloat64(key)
	if b == nil {
		return nil, onError(fmt.Errorf("missing required float64 from %s", key))
	} else {
		return b, nil
	}
}

// GetBool gets a key and sets to true, false or nil using the Truthy and Falsey
// variables
func GetBool(key string) *bool {
	if v, e := strconv.ParseBool(get(key)); e != nil {
		return nil
	} else {
		return &v
	}
}

func GetOrSetBool(key string, val bool) *bool {
	if IsSet(key) {
		return GetBool(key)
	}
	Set(key, val)
	return &val
}

func RequireBool(key string) (*bool, error) {
	b := GetBool(key)
	if b == nil {
		return nil, onError(fmt.Errorf("missing required bool from %s", key))
	} else {
		return b, nil
	}
}

func empty(s *string) bool {
	return (s == nil || *s == "")
}

func get(key string) string {
	return os.Getenv(key)
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func onError(e error) error {
	if PanicOnRequire {
		panic(e)
	}
	return e
}
