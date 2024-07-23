package main

import (
	service2 "golang/internal/service"
	"os"
)

func main() {
	inputFile := "input.txt"
	outputFile := "output.txt"

	if len(os.Args) != 1 {
		args := os.Args[1:]
		inputFile = args[0]
		outputFile = args[1]
	}

	prod := service2.NewProducer(inputFile)
	pres := service2.NewPresenter(outputFile)

	serv := service2.NewService(prod, pres)

	err := serv.Run()
	if err != nil {
		return
	}
}
