package race420_q1

import (
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/20 上午10:44
 * @FilePath: string_sequence_test.go
 * @Description:
 */

func Test_stringSequence(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{target: "abc"},
			want: []string{"a", "aa", "ab", "aba", "abb", "abc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringSequence(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
