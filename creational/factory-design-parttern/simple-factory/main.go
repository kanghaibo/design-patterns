package main

import (
	"design-patterns/simple-factory/factory"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	t := factory.NewTranslator("english")
	fmt.Println(t.Translate("english"))

	t = factory.NewTranslator("haha")
	fmt.Println(t.Translate("haha"))
}
