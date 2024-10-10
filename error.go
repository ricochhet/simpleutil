package simpleutil

import "log"

func WrapError(a func() (string, error)) string {
	s, err := a()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return s
}
