package main

import (
	"context"
)

func main() {
	l := &Loader{}

	d, err := l.GetDict(context.Background(), DictDir1, "v0.0.1")

	if err != nil {
		panic(err)
	}

	println(d.Meta.Description)
}
