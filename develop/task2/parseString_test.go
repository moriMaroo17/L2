package parseString

import (
	"testing"
)

func Test_parseString(t *testing.T) {
	type args struct {
		stringForUnpack string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "first",
			args:    args{stringForUnpack: "a4bc2d5e"},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "second",
			args:    args{stringForUnpack: "abcd"},
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "third",
			args:    args{stringForUnpack: "45"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "fourth",
			args:    args{stringForUnpack: ""},
			want:    "",
			wantErr: false,
		},
		{
			name:    "fifth",
			args:    args{stringForUnpack: `qwe\4\5`},
			want:    "qwe45",
			wantErr: false,
		},
		{
			name:    "sixth",
			args:    args{stringForUnpack: `qwe\45`},
			want:    "qwe44444",
			wantErr: false,
		},
		{
			name:    "seventh",
			args:    args{stringForUnpack: `qwe\\5`},
			want:    `qwe\\\\\`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseString(tt.args.stringForUnpack)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
