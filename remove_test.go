/*
 * @Author: nijineko
 * @Date: 2025-06-07 10:59:07
 * @LastEditTime: 2025-06-07 11:42:53
 * @LastEditors: nijineko
 * @Description: remove color from text
 * @FilePath: \colorize\remove_test.go
 */
package colorize

import "testing"

func TestRemove(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Remove color codes",
			args: args{
				Text: Set("Hello, Noa!", ConfigBold, BackgroundBlack, TextWhite),
			},
			want: "Hello, Noa!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.Text); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
