package helpers

import "log"

func Safe[T any](x T, err error) T {
	if err != nil {
		log.Fatal(err)
	}

	return x
}
