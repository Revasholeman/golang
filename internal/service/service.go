package service

import (
	"sync"
)

//go:generate mockgen -source=service.go -destination=../internal/mock/service.go -package=mock
type producer interface {
	Produce() ([]string, error)
}

//go:generate mockgen -source=service.go -destination=../internal/mock/service.go -package=mock
type presenter interface {
	Present([]string) error
}

type Service struct {
	prod producer
	pres presenter
	wg   sync.WaitGroup
}

func NewService(prod producer, pres presenter) *Service {
	return &Service{pres: pres, prod: prod, wg: sync.WaitGroup{}}
}

func (s *Service) Run() error {
	file, err := s.prod.Produce()
	wg := &s.wg
	if err != nil {
		return err
	}

	fanIn := make(chan string, len(file))
	fanOut := make(chan string, len(file))

	var spamMasked []string

	wg.Add(len(file))
	go func() {
		for _, text := range file {
			fanIn <- text
		}
	}()

	go func() {
		for text := range fanIn {
			fanOut <- s.SpamMasker(text)
		}
		close(fanIn)
	}()

	go func() {
		for val := range fanOut {
			spamMasked = append(spamMasked, val)
			wg.Done()
		}
	}()

	err = s.pres.Present(spamMasked)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SpamMasker(message string) string {
	http := [7]rune{'h', 't', 't', 'p', ':', '/', '/'}
	count := 0
	result := make([]rune, 0, len(message))
	for _, w := range message {
		switch {
		case w == ' ':
			result = append(result, w)
			count = 0
		case count == 7 && w != ' ':
			result = append(result, '*')
		case w == http[count]:
			result = append(result, w)
			count++
		default:
			result = append(result, w)
		}
	}
	return string(result)
}
