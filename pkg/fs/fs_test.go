package fs

import "testing"

func TestByteSize_String(t *testing.T) {
	tests := []struct {
		name string
		b    ByteSize
		want string
	}{
		{
			name: "1B",
			b:    1,
			want: "1.00B",
		},
		{
			name: "1KB",
			b:    1024,
			want: "1.00KB",
		},
		{
			name: "1.27TB",
			b:    1393812756406,
			want: "1.27TB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
