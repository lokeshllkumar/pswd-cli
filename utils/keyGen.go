package utils

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
)

func KeyGen() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

func KeyWrite() {
	key, err := KeyGen()
	if err != nil {
		fmt.Println("Error generating key")
		return
	}
	err = ioutil.WriteFile("./utils/key.txt", key, 0666)
	if err != nil {
		fmt.Println("Error writing key to stored file")
	}
}
