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
				date: "2006-01-02 15:04:05",
			},
			want:    false,
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
