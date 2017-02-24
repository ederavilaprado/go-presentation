package main

import (
	"errors"
	"fmt"
)

func funcTeste(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nome não informado")
	}

	return fmt.Sprintf("Hello %s, welcome to golang\n", name), nil
}

func main() {
	r, err := funcTeste("")
	if err != nil {
		panic(err)
	}
	fmt.Printf(r)
}
