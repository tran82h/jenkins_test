package tests

import (
	"testing"
)

func Test_keyOp_Degenerate(t *testing.T) {

	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantX   int
		wantY   int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				s: "40_99",
			},
			wantX: 40,
			wantY: 99,
		},
		{
			name: "failure",
			args: args{
				s: "4099",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kp := GetKeyOperator()
			gotX, gotY, gotErr := kp.Degenerate(tt.args.s)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("keyOp.Degenerate() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if gotX != tt.wantX {
				t.Errorf("keyOp.Degenerate() gotX = %v, wantX %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("keyOp.Degenerate() gotY = %v, wantX %v", gotY, tt.wantY)
			}
		})
	}
}
