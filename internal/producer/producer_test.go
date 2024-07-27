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
		{
			name: "TestNewProducer",
			args: args{
				filePath: "output.txt",
			},
			want: &Producer{
				filePath: "output.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProducer(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProducer_Produce(t *testing.T) {
	type fields struct {
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name: "TestProducer_Produce",
			fields: fields{
				filePath: "input.txt",
			},
			want:    []string{"http://www.baidu.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			producer := &Producer{
				filePath: tt.fields.filePath,
			}
			got, err := producer.Produce()
			if (err != nil) != tt.wantErr {
				t.Errorf("Produce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Produce() got = %v, want %v", got, tt.want)
			}
		})
	}
}
