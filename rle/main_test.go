package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	type args struct {
		original []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "basic compress",
			args: args{
				original: []byte("baababab"),
			},
			want: []byte("b2ababab"),
		},
		{
			name: "complex compress",
			args: args{
				original: []byte("aaabbbbbbabbab"),
			},
			want: []byte("3a6ba2bab"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compress := CompressByString(tt.args.original)
			compressRLE := CompressOptimized(tt.args.original)

			assert.Equal(t, tt.want, compress)
			assert.Equal(t, tt.want, compressRLE)
		})
	}
}

func BenchmarkCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressByString([]byte("aaabbbbbbabbabaaabbbbbbabbabaaaaaabbbbbbabbabaaa"))
	}
}

func BenchmarkCompressOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressOptimized([]byte("aaabbbbbbabbabaaabbbbbbabbabaaaaaabbbbbbabbabaaa"))
	}
}
