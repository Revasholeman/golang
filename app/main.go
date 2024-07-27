package main

import (
	presenter "golang/internal/presenter"
	producer "golang/internal/producer"
	service "golang/internal/service"
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

	prod := producer.NewProducer(inputFile)
	pres := presenter.NewPresenter(outputFile)

	serv := service.NewService(prod, pres)

	err := serv.Run()
	if err != nil {
		return
	}
}
