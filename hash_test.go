package hasher

import (
	"bytes"
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		i any
	}

	type Struct1 struct {
		A int
		B string
	}

	type Struct2 struct {
		a int
		B string
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Nil interface",
			args: args{
				i: nil,
			},
			wantErr: true,
		},
		{
			name: "Integer",
			args: args{
				i: 42,
			},
			want: []byte{107, 155, 250, 45, 100, 39, 24, 65, 194, 101, 16, 159, 195, 239, 176, 2, 85, 62, 34, 119, 216, 52, 216, 241, 211, 28, 20, 240, 99, 189, 128, 4},
		},
		{
			name: "String",
			args: args{
				i: "Hello world",
			},
			want: []byte{55, 253, 231, 128, 85, 174, 168, 85, 98, 99, 96, 141, 236, 78, 221, 174, 172, 151, 211, 237, 174, 98, 120, 65, 69, 59, 214, 24, 76, 251, 178, 13},
		},
		{
			name: "Struct1",
			args: args{
				i: Struct1{
					A: 42,
					B: "Hello World",
				},
			},
			want: []byte{52, 149, 63, 48, 99, 112, 7, 227, 18, 218, 193, 93, 41, 166, 45, 14, 42, 231, 74, 88, 142, 131, 46, 220, 136, 250, 40, 119, 132, 8, 34, 19},
		},
		{
			name: "Struct1 reference",
			args: args{
				i: &Struct1{
					A: 42,
					B: "Hello World",
				},
			},
			want: []byte{52, 149, 63, 48, 99, 112, 7, 227, 18, 218, 193, 93, 41, 166, 45, 14, 42, 231, 74, 88, 142, 131, 46, 220, 136, 250, 40, 119, 132, 8, 34, 19},
		},
		{
			name: "Struct2",
			args: args{
				i: Struct2{
					a: 42,
					B: "Hello World",
				},
			},
			want: []byte{74, 75, 174, 205, 4, 102, 158, 45, 148, 24, 43, 215, 166, 24, 121, 67, 47, 182, 192, 57, 143, 132, 165, 235, 94, 132, 2, 233, 142, 76, 242, 156},
		},
		{
			name: "Struct2 reference",
			args: args{
				i: &Struct2{
					a: 42,
					B: "Hello World",
				},
			},
			want: []byte{74, 75, 174, 205, 4, 102, 158, 45, 148, 24, 43, 215, 166, 24, 121, 67, 47, 182, 192, 57, 143, 132, 165, 235, 94, 132, 2, 233, 142, 76, 242, 156},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hash(tt.args.i)
			if err != nil && !tt.wantErr {
				t.Errorf("Test Name = '%s', Hash() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Test Name = '%s', Hash() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
