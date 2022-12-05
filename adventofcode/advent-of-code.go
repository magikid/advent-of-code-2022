package adventofcode

import (
	"log"
)

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		log.Fatalf("Runtime error: %v\n", err)
		return 1
	}

	return 0
}
