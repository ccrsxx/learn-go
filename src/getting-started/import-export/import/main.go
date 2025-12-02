package main

import (
	"fmt"

	"github.com/ccrsxx/learn-go/src/getting-started/import-export/export"
)

func main() {
	emilia := export.Emilia

	fmt.Println("Best girl:", emilia)

	export.BestGirl()
}
