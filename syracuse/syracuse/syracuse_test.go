package syracuse_test

import (
	"testing"

	"github.com/Chroq/syracuse/syracuse"
	"github.com/stretchr/testify/assert"
)

func TestSyracuse(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Syracuse(1)",
			args: args{1},
			want: []int{},
		},
		{
			name: "Syracuse(512)",
			args: args{512},
			want: []int{256, 128, 64, 32, 16, 8, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, syracuse.Syracuse(tt.args.n))
		})
	}
}

func BenchmarkSyracuse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syracuse.Syracuse(3)
	}
}
