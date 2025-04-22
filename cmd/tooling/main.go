package main

import (
	"fmt"
	"os"

	"github.com/qba73/stamp"
)

func main() {
	if err := stamp.GenerateCertAndKey(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
}
