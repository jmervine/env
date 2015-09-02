# envcfg
Simple configuration utility around os.{Get,Set}env

# usage

```
PACKAGE DOCUMENTATION

package envcfg
    import "github.com/jmervine/envcfg"


VARIABLES

var PanicOnRequire = false
    PanicOnRequire forces panics when Require- methods fail

FUNCTIONS

func Get(key string) *string
    Get gets a key and returns a string

func GetBool(key string) *bool
    GetBool gets a key and sets to true, false or nil using the Truthy and
    Falsey variables

func GetDuration(key string) *time.Duration

func GetFloat32(key string) *float32
    GetFloat32 gets a key and returns an float32

func GetFloat64(key string) *float64
    GetFloat64 gets a key and returns an float64

func GetInt(key string) *int
    GetInt gets a key and returns an int

func GetInt32(key string) *int32
    GetInt32 gets a key and returns an int32

func GetInt64(key string) *int64
    GetInt64 gets a key and returns an int64

func GetOrSet(key string, val interface{}) *string
    GetOrSet gets a key and returns a string or set's the default

func GetOrSetBool(key string, val bool) *bool

func GetOrSetDuration(key string, val time.Duration) *time.Duration

func GetOrSetFloat32(key string, val float32) *float32

func GetOrSetFloat64(key string, val float64) *float64

func GetOrSetInt(key string, val int) *int

func GetOrSetInt32(key string, val int32) *int32

func GetOrSetInt64(key string, val int64) *int64

func GetOrSetString(key, val string) *string
    GetOrSetString is an alias to GetOrSet

func GetOrSetUint(key string, val uint) *uint

func GetOrSetUint64(key string, val uint64) *uint64

func GetString(key string) *string
    GetString is an alias to Get

func GetUint(key string) *uint
    GetUint gets a key and returns an uint

func GetUint64(key string) *uint64
    GetUint64 gets a key and returns an uint64

func IsSet(key string) bool
    IsSet return true if a key is not "" in os.Getenv

func Require(key string) (*string, error)
    Require gets a key and returns a string or an error if it's set to "" in
    os.Getenv

func RequireBool(key string) (*bool, error)

func RequireDuration(key string) (*time.Duration, error)

func RequireFloat32(key string) (*float32, error)

func RequireFloat64(key string) (*float64, error)

func RequireInt(key string) (*int, error)

func RequireInt32(key string) (*int32, error)

func RequireInt64(key string) (*int64, error)

func RequireString(key string) (*string, error)

func RequireUint(key string) (*uint, error)

func RequireUint64(key string) (*uint64, error)

func Set(key string, val interface{})
    Set sets via an interface

func SetMap(m map[string]interface{})
    SetMap iterates over a map and sets keys to values

```

# LICENSE

```
Copyright (c) 2015 Joshua Mervine

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
