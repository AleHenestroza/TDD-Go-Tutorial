package main

import (
	clockface "go-tdd-tutorial/16-Maths"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SvgWriter(os.Stdout, t)
}
