package main

import (
	"design-patterns/factory-method/factory"
	"fmt"
)

func main() {
	f := factory.NewTranslatorFactory("english")
	t := f.Create()
	fmt.Println(t.Translate("haha"))
}
