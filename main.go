package main

import (
	"github.com/Olusamimaths/tries/app"
	"github.com/Olusamimaths/tries/impl"
)

func main() {
	impl.TestTrie()
	app.Start(":8080")
}
