package main

import "testing"

func Test_ftoc(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "freezing",
			args: args{f: 32},
			want: 0,
		},
		{
			name: "boiling",
			args: args{f: 212},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ftoc(tt.args.f); got != tt.want {
				t.Errorf("ftoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
