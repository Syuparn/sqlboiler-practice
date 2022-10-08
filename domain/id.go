package domain

import "github.com/oklog/ulid/v2"

var generateID = generateULID

func generateULID() string {
	return ulid.Make().String()
}
