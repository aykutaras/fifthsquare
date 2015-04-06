package main

import (
	"fmt"
	"github.com/aykutaras/fifthsquare"
)

func main() {
	serverUrl := "localhost:4001"
	fmt.Println(fmt.Sprintf("Server started at: %s", serverUrl))
	fifthsquare.InitHttpService(serverUrl)
}
