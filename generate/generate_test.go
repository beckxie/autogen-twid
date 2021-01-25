package generate

import (
	"reflect"
	"strconv"
	"testing"
)

const randomCount = 100

func Test_generatePrefix(t *testing.T) {
	for i := 1; i <= randomCount; i++ {
		t.Run("random-test#"+strconv.Itoa(i), func(t *testing.T) {
			if got := generatePrefix(); !isValidPrefixRange(got) {
				t.Errorf("generatePrefix() = %v", got)
			}
		})
	}
}

func isValidPrefixRange(number int) bool {
	if number >= 10 && number <= 35 {
		return true
	}
	return false
}

func Test_genGenderNumber(t *testing.T) {
	for i := 1; i <= randomCount; i++ {
		t.Run("random-test#"+strconv.Itoa(i), func(t *testing.T) {
			if got := genGenderNumber(); !isValidSecond(got) {
				t.Errorf("genGenderNumber() = %v", got)
			}
		})
	}
}

func isValidSecond(i int) bool {
	switch i {
	case 1, 2, 8, 9:
		return true
	}
	return false
}

func Test_intToSliceInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name         string
		args         args
		wantSliceInt []int
	}{
		{
			name: "case1",
			args: args{
				i: 5566,
			},
			wantSliceInt: []int{5, 5, 6, 6},
		},
		{
			name: "case2",
			args: args{
				i: 987,
			},
			wantSliceInt: []int{9, 8, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSliceInt := intToSliceInt(tt.args.i); !reflect.DeepEqual(gotSliceInt, tt.wantSliceInt) {
				t.Errorf("intToSliceInt() = %v, want %v", gotSliceInt, tt.wantSliceInt)
			}
		})
	}
}

func TestGenerateID(t *testing.T) {
	for i := 1; i <= randomCount; i++ {
		t.Run("random test#"+strconv.Itoa(i), func(t *testing.T) {
			resultID := GenerateID()
			if len(resultID) <= 0 {
				t.Errorf("GenerateID() = %v", resultID)
			}
		})
	}
}

func Test_generateSuffix(t *testing.T) {
	type args struct {
		pre int
		mid int
	}
	tests := []struct {
		name    string
		args    args
		wantSuf int
		wantErr bool
	}{
		{
			name: "happy case#1",
			args: args{
				pre: 10,
				mid: 12345678,
			},
			wantSuf: 9,
			wantErr: false,
		},
		{
			name: "bad case#1",
			args: args{
				pre: 10,
				mid: 1234567,
			},
			wantSuf: 0,
			wantErr: true,
		},
		{
			name: "bad case#2",
			args: args{
				pre: 0,
				mid: 0,
			},
			wantSuf: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuf, err := generateSuffix(tt.args.pre, tt.args.mid)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateSuffix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuf != tt.wantSuf {
				t.Errorf("generateSuffix() = %v, want %v", gotSuf, tt.wantSuf)
			}
		})
	}
}

func Test_generateMiddleNumber(t *testing.T) {
	for i := 1; i <= randomCount; i++ {
		t.Run("random test#"+strconv.Itoa(i), func(t *testing.T) {
			if got := generateMiddleNumber(); !isValidMiddleRange(got) {
				t.Errorf("generateMiddleNumber() = %v", got)
			}
		})
	}
}

func isValidMiddleRange(number int) bool {
	if number >= 10000000 && number <= 99999999 {
		return true
	}
	return false
}
