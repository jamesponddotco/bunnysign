package main

import (
	"os"

	"git.sr.ht/~jamesponddotco/bunnysign/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
