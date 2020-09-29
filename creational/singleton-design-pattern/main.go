package main

import (
	"design-patterns/singleton/singleton"
)

func main() {
	instance := singleton.GetInstance()
	_ = instance
}
