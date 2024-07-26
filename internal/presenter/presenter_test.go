package presenter

import (
	"reflect"
	"testing"
)

func TestNewPresenter(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want *Presenter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPresenter(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPresenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPresenter_present(t *testing.T) {
	type fields struct {
		filePath string
	}
	type args struct {
		text []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Presenter{
				filePath: tt.fields.filePath,
			}
			if err := p.present(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("Presenter.present() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
