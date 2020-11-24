package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start Project")

	application := app.Generate()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
