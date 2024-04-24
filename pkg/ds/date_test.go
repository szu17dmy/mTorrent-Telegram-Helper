package ds

import "testing"

func TestDateAfterNow(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				date: "2023-01-01 00:00:00",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args{
				date: "2099-01-01 00:00:00",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateAfterNow(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateAfterNow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DateAfterNow() got = %v, want %v", got, tt.want)
			}
		})
	}
}
