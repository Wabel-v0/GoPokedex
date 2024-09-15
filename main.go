package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)

	}
	fmt.Println("status code:", res.Status)
	fmt.Println("==========")
	fmt.Println("Header:", res.Header)
	fmt.Println("==========")
	fmt.Println("Body:", string(data))
}
