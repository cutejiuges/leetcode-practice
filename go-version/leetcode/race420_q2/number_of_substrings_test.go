package race420_q2

import "testing"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/20 上午11:36
 * @FilePath: number_of_substrings_test.go
 * @Description:
 */

func Test_numberOfSubstrings(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{s: "aaa", k: 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSubstrings(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("numberOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
