package cache_test

import (
	"testing"

	. "github.com/avialeta/api/cache"
)

func TestStore(t *testing.T) {
	Set("foo", "FOO")
	v := Get("foo")
	t.Log("v:", v)
}
