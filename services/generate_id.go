package services

import gid "github.com/matoous/go-nanoid/v2"

func Nanoid() string {
	id, err := gid.New(36)
	if err != nil {
		panic(err)
	}

	return id
}
