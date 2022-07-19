package main

import (
	"fmt"
	"kazdream/app"
	"time"
)

func main() {
	t := time.Now()
	app.App()
	fmt.Printf("%s\n", time.Since(t))
}
