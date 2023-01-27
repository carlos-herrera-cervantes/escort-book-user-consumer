package config

import "os"

type dictum struct {
	FromUser string
}

var singleDictum *dictum

func InitializeDictum() *dictum {
	if singleDictum != nil {
		return singleDictum
	}

	lock.Lock()
	defer lock.Unlock()

	singleDictum = &dictum{
		FromUser: os.Getenv("FROM_USER"),
	}

	return singleDictum
}
