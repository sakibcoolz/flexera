package utils

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStore_StoreApplicationInvetory(t *testing.T) {
	type fields struct {
		applicationInventory map[string]int
		inventryQueue        []string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Pass",
			fields: fields{
				applicationInventory: make(map[string]int),
				inventryQueue:        make([]string, 0),
			},
			args: args{
				str: "123|123123|23123|DESKTOP",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				applicationInventory: tt.fields.applicationInventory,
				inventryQueue:        tt.fields.inventryQueue,
			}
			s.StoreApplicationInvetory(tt.args.str)
		})
	}
}

func TestDataEngine(t *testing.T) {
	type args struct {
		dataTranferChannel chan []string
	}
	tests := []struct {
		name string
		args args
		want *Store
	}{
		{
			name: "Pass",
			args: args{
				dataTranferChannel: make(chan []string),
			},
			want: &Store{
				applicationInventory: map[string]int{"188_DESKTOP": 1, "606_DESKTOP": 1, "62_DESKTOP": 1, "APPLICATIONID_COMPUTERTYPE": 1},
				inventryQueue:        []string{"APPLICATIONID", "606", "62", "188"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func(filename string, dataTranferChannel chan []string) {
				err := FileReader(filename, dataTranferChannel)
				if err != nil {
					log.Fatal("cant process file due to error", err.Error())
				}
			}("../stage/sample-test.csv", tt.args.dataTranferChannel)
			got := DataEngine(tt.args.dataTranferChannel)
			if data := cmp.Diff(got.inventryQueue, tt.want.inventryQueue); data != "" {
				fmt.Println(data)
				t.Errorf("DataEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}
