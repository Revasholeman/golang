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
	type fields struct {
		prod *mock.Mockproducer
		pres *mock.Mockpresenter
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
		{
			name: "TestService_SpamMasker",
			fields: fields{
				prod: mock.NewMockproducer(ctrl),
				pres: mock.NewMockpresenter(ctrl),
			},
			args: args{
				message: "http://123",
			},
			want: "http://***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				prod: tt.fields.prod,
				pres: tt.fields.pres,
			}
			if got := s.SpamMasker(tt.args.message); got != tt.want {
				t.Errorf("SpamMasker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProducer := mock.NewMockproducer(ctrl)
	mockPresenter := mock.NewMockpresenter(ctrl)

	expectedFile := []string{"input.txt", "output.txt"}
	mockProducer.EXPECT().
		Produce().
		Return(expectedFile, nil)

	mockPresenter.EXPECT().
		Present(gomock.Any()).
		Return(nil)

	svc := &Service{
		prod: mockProducer,
		pres: mockPresenter,
	}

	err := svc.Run()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
