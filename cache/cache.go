package cache

import (
	"github.com/avialeta/api/log"

	"github.com/xuyu/goredis"
)

var client *goredis.Redis

func init() {
	client = dial()
}

func dial() *goredis.Redis {
	dialConfig := goredis.DialConfig{
		Address: "avialeta-cache:6379",
	}
	client, err := goredis.Dial(&dialConfig)
	if err != nil {
		log.Error.Fatal(err)
	}
	return client
}

func Set(k, v string) {
	err := client.Set(k, v, 0, 0, false, false)
	if err != nil {
		log.Error.Print(err)
	}
}

func Get(k string) string {
	v, err := client.Get(k)
	if err != nil {
		log.Error.Print(err)
		return ""
	}
	return string(v)
}
