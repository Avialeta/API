package job

import (
	"fmt"
	"time"
)

func Foo() {
	for {
		fmt.Println("foo")
		time.Sleep(time.Second)
	}
}

func Bar() {
	for {
		fmt.Println("bar")
		time.Sleep(time.Second)
	}
}
