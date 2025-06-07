/*
 * @Author: nijineko
 * @Date: 2025-06-07 11:05:59
 * @LastEditTime: 2025-06-07 11:43:59
 * @LastEditors: nijineko
 * @Description: quick method for set text color
 * @FilePath: \colorize\text_test.go
 */
package colorize

import (
	"testing"
)

func TestBlackText(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Black text",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;0;30mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlackText(tt.args.Text); got != tt.want {
				t.Errorf("BlackText() = %v, want %v", got, tt.want)
			} else {
				t.Log("BlackText() = ", got)
			}
		})
	}
}

func TestRedText(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Red text",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;0;31mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RedText(tt.args.Text); got != tt.want {
				t.Errorf("RedText() = %v, want %v", got, tt.want)
			} else {
				t.Log("RedText() = ", got)
			}
		})
	}
}

func TestGreenText(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Green text",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;0;32mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreenText(tt.args.Text); got != tt.want {
				t.Errorf("GreenText() = %v, want %v", got, tt.want)
			} else {
				t.Log("GreenText() = ", got)
			}
		})
	}
}

func TestYellowText(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yellow text",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;0;33mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YellowText(tt.args.Text); got != tt.want {
				t.Errorf("YellowText() = %v, want %v", got, tt.want)
			} else {
				t.Log("YellowText() = ", got)
			}
		})
	}
}

func TestBlueText(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Blue text",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;0;34mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlueText(tt.args.Text); got != tt.want {
				t.Errorf("BlueText() = %v, want %v", got, tt.want)
			} else {
				t.Log("BlueText() = ", got)
			}
		})
	}
}

func TestMagentaText(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Magenta text",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;0;35mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MagentaText(tt.args.Text); got != tt.want {
				t.Errorf("MagentaText() = %v, want %v", got, tt.want)
			} else {
				t.Log("MagentaText() = ", got)
			}
		})
	}
}

func TestCyanText(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Cyan text",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;0;36mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CyanText(tt.args.Text); got != tt.want {
				t.Errorf("CyanText() = %v, want %v", got, tt.want)
			} else {
				t.Log("CyanText() = ", got)
			}
		})
	}
}

func TestWhiteText(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "White text",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;0;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhiteText(tt.args.Text); got != tt.want {
				t.Errorf("WhiteText() = %v, want %v", got, tt.want)
			} else {
				t.Log("WhiteText() = ", got)
			}
		})
	}
}
