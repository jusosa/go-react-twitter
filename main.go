package main

import (
	"github.com/jusosa/go-react-twitter/bd"
	"github.com/jusosa/go-react-twitter/handlers"
	"log"
)

func main() {
	if bd.CheckConnection()== 0 {
		log.Fatal("DB Connection Fail")
		return
	}

	handlers.Handle()
}