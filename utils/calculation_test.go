package utils

import "testing"

func TestStore_BalancePurchase(t *testing.T) {
	type fields struct {
		applicationInventory map[string]int
		inventryQueue        []string
	}
	type args struct {
		appid string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Pass laptop more",
			fields: fields{
				applicationInventory: map[string]int{"33_LAPTOP": 10, "33_DESKTOP": 5},
				inventryQueue:        []string{"33"},
			},
			args: args{appid: "33"},
		},
		{
			name: "Pass desktop is more",
			fields: fields{
				applicationInventory: map[string]int{"33_LAPTOP": 5, "33_DESKTOP": 10},
				inventryQueue:        []string{"33"},
			},
			args: args{appid: "33"},
		},
		{
			name: "Pass desktop and laptop count same",
			fields: fields{
				applicationInventory: map[string]int{"33_LAPTOP": 5, "33_DESKTOP": 5},
				inventryQueue:        []string{"33"},
			},
			args: args{appid: "33"},
		},
		{
			name: "Failed -appid not found",
			fields: fields{
				applicationInventory: map[string]int{"33_LAPTOP": 5, "33_DESKTOP": 5},
				inventryQueue:        []string{"", ""},
			},
			args: args{appid: "33"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				applicationInventory: tt.fields.applicationInventory,
				inventryQueue:        tt.fields.inventryQueue,
			}
			s.BalancePurchase(tt.args.appid)
		})
	}
}
