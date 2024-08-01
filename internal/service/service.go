package service

import (
	"sync"
	"time"
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

	limit := 10

	fanIn := make(chan string)
	fanOut := make(chan string)

	semaphore := make(chan struct{}, limit)

	var spamMasked []string

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, text := range file {
			fanIn <- text
			<-semaphore
		}
	}()

	go func() {
		for text := range fanIn {
			semaphore <- struct{}{}
			fanOut <- s.SpamMasker(text)
			<-semaphore
		}
		close(fanIn)
	}()

	go func() {
		wg.Wait()
		time.Sleep(100 * time.Microsecond)
		close(fanOut)
	}()

	for val := range fanOut {
		spamMasked = append(spamMasked, val)
	}

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
