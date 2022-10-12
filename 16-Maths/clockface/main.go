package main

import (
	"os"
	"time"

	"go-tdd-tutorial/16-Maths/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
