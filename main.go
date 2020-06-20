package main

import (
	"fmt"

	"github.com/parsnips/dupe/util"
)

func main() {
	var checker util.Dupe
	checker = &util.RealDupeChecker{}

	fmt.Println(checker.IsDupe("a.txt", "b.txt"))
}
