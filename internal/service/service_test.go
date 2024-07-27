package service

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		prod producer
		pres presenter
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.prod, tt.args.pres); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Run(t *testing.T) {
	type fields struct {
		prod producer
		pres presenter
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				prod: tt.fields.prod,
				pres: tt.fields.pres,
			}
			if err := s.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Service.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_SpamMasker(t *testing.T) {
	type fields struct {
		prod producer
		pres presenter
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				prod: tt.fields.prod,
				pres: tt.fields.pres,
			}
			if got := s.SpamMasker(tt.args.message); got != tt.want {
				t.Errorf("Service.SpamMasker() = %v, want %v", got, tt.want)
			}
		})
	}
}
