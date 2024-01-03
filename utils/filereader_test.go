package utils

import (
	"fmt"
	"testing"
)

func TestFileReader(t *testing.T) {
	type args struct {
		filename           string
		dataTranferChannel chan []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Pass",
			args: args{
				filename:           "../stage/sample-test.csv",
				dataTranferChannel: make(chan []string, 1000),
			},
			wantErr: false,
		},
		{
			name: "Fail",
			args: args{
				filename:           "stage/sample-small.csv",
				dataTranferChannel: make(chan []string, 1000),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func(dataTranferChannel chan []string) {
				for data := range dataTranferChannel {
					fmt.Println(data)
				}
			}(tt.args.dataTranferChannel)
			if err := FileReader(tt.args.filename, tt.args.dataTranferChannel); (err != nil) != tt.wantErr {
				t.Errorf("FileReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
