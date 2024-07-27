package producer

import (
	"reflect"
	"testing"
)

func TestNewProducer(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want *Producer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProducer(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProducer_produce(t *testing.T) {
	type fields struct {
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			producer := &Producer{
				filePath: tt.fields.filePath,
			}
			got, err := producer.produce()
			if (err != nil) != tt.wantErr {
				t.Errorf("Producer.produce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Producer.produce() = %v, want %v", got, tt.want)
			}
		})
	}
}
