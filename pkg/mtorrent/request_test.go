package mtorrent

import (
	"fmt"
	"testing"
)

func Test_postForm(t *testing.T) {
	type args struct {
		api  string
		data map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "System/hello",
			args: args{
				api:  "/api/system/hello",
				data: nil,
			},
			want: "200 OK",
		},
		{
			name: "System/state",
			args: args{
				api:  "/api/system/state",
				data: nil,
			},
			want: "200 OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := postForm(tt.args.api, &tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("postForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postForm() got = %v, want %v", got, tt.want)
			}
			fmt.Println(string(got1))
		})
	}
}

func Test_postJson(t *testing.T) {
	type args struct {
		api  string
		form any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []byte
		wantErr bool
	}{
		{
			name: "User/bases",
			args: args{
				api: "/api/member/bases",
				form: map[string][]int{
					"ids": {294703},
				},
			},
			want: "200 OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := postJson(tt.args.api, &tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("postJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("postJson() got = %v, want %v", got, tt.want)
			}
			fmt.Println(string(got1))
		})
	}
}
