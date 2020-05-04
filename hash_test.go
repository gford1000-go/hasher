package hasher

import (
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		i interface{}
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
		want    DataHash
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
			want: DataHash{107, 155, 250, 45, 100, 39, 24, 65, 194, 101, 16, 159, 195, 239, 176, 2, 85, 62, 34, 119, 216, 52, 216, 241, 211, 28, 20, 240, 99, 189, 128, 4},
		},
		{
			name: "String",
			args: args{
				i: "Hello world",
			},
			want: DataHash{55, 253, 231, 128, 85, 174, 168, 85, 98, 99, 96, 141, 236, 78, 221, 174, 172, 151, 211, 237, 174, 98, 120, 65, 69, 59, 214, 24, 76, 251, 178, 13},
		},
		{
			name: "Struct1",
			args: args{
				i: Struct1{
					A: 42,
					B: "Hello World",
				},
			},
			want: DataHash{41, 72, 166, 110, 158, 224, 183, 249, 18, 102, 68, 91, 124, 37, 40, 157, 186, 234, 35, 184, 21, 185, 160, 83, 59, 152, 59, 241, 47, 117, 220, 121},
		},
		{
			name: "Struct1 reference",
			args: args{
				i: &Struct1{
					A: 42,
					B: "Hello World",
				},
			},
			want: DataHash{41, 72, 166, 110, 158, 224, 183, 249, 18, 102, 68, 91, 124, 37, 40, 157, 186, 234, 35, 184, 21, 185, 160, 83, 59, 152, 59, 241, 47, 117, 220, 121},
		},
		{
			name: "Struct2",
			args: args{
				i: Struct2{
					a: 42,
					B: "Hello World",
				},
			},
			want: DataHash{85, 250, 140, 46, 12, 8, 228, 34, 254, 128, 48, 100, 205, 242, 169, 196, 145, 232, 190, 210, 240, 30, 17, 82, 100, 172, 194, 229, 168, 159, 128, 189},
		},
		{
			name: "Struct2 reference",
			args: args{
				i: &Struct2{
					a: 42,
					B: "Hello World",
				},
			},
			want: DataHash{85, 250, 140, 46, 12, 8, 228, 34, 254, 128, 48, 100, 205, 242, 169, 196, 145, 232, 190, 210, 240, 30, 17, 82, 100, 172, 194, 229, 168, 159, 128, 189},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hash(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
