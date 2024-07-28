package service

import (
	"github.com/golang/mock/gomock"
	mock "golang/internal/mock/service"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		prod *mock.Mockproducer
		pres *mock.Mockpresenter
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "TestNewService",
			args: args{
				prod: mock.NewMockproducer(ctrl),
				pres: mock.NewMockpresenter(ctrl),
			},
			want: &Service{
				prod: mock.NewMockproducer(ctrl),
				pres: mock.NewMockpresenter(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.prod, tt.args.pres); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_SpamMasker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name    string
		fields  *Service
		message string
		want    string
	}{
		{
			name: "TestService_SpamMasker",
			fields: &Service{
				prod: mock.NewMockproducer(ctrl),
				pres: mock.NewMockpresenter(ctrl),
			},
			message: "http://123",
			want:    "http://***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.fields.SpamMasker(tt.message); got != tt.want {
				t.Errorf("SpamMasker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		prod *mock.Mockproducer
		pres *mock.Mockpresenter
	}
	tests := []struct {
		name             string
		fields           fields
		expectedFiles    []string
		expectedRunError error
	}{
		{
			name: "Successful Run",
			fields: fields{
				prod: mock.NewMockproducer(ctrl),
				pres: mock.NewMockpresenter(ctrl),
			},
			expectedFiles:    []string{"input.txt", "output.txt"},
			expectedRunError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.prod.EXPECT().
				Produce().
				Return(tt.expectedFiles, nil)

			tt.fields.pres.EXPECT().
				Present(gomock.Any()).
				Return(nil)

			s := &Service{
				prod: tt.fields.prod,
				pres: tt.fields.pres,
			}

			err := s.Run()
			if err != tt.expectedRunError {
				t.Fatalf("Expected no error, got %v", err)
			}
		})
	}
}
