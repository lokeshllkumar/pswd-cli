package utils

import (
	"fmt"
	"io/ioutil"
)

func readKey() ([]byte, error){
	data, err := ioutil.ReadFile("./utils/key.txt")
	if err != nil {
		fmt.Printf("Error reading key from stored file: %v\n", err)
		return []byte(""), err
	}

	return data, nil

}

var Key, _ = readKey()