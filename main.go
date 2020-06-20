package main

import (
	"fmt"

	"github.com/parsnips/dupe/dupe"
)

func main() {
	var checker dupe.Dupe
	checker = &dupe.RealDupeChecker{}

	fmt.Println(checker.IsDupe("a.txt", "b.txt"))
}
