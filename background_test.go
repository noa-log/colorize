/*
 * @Author: nijineko
 * @Date: 2025-06-07 11:13:20
 * @LastEditTime: 2025-06-07 11:13:26
 * @LastEditors: nijineko
 * @Description: quick method for set background color
 * @FilePath: \colorize\background.go
 */
package colorize

import (
	"testing"
)

func TestBlackBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Black background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;40;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlackBackground(tt.args.Text); got != tt.want {
				t.Errorf("BlackBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BlackBackground() = ", got)
			}
		})
	}
}

func TestRedBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Red background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;41;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RedBackground(tt.args.Text); got != tt.want {
				t.Errorf("RedBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("RedBackground() = ", got)
			}
		})
	}
}

func TestGreenBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Green background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;42;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreenBackground(tt.args.Text); got != tt.want {
				t.Errorf("GreenBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("GreenBackground() = ", got)
			}
		})
	}
}

func TestYellowBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Yellow background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;43;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YellowBackground(tt.args.Text); got != tt.want {
				t.Errorf("YellowBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("YellowBackground() = ", got)
			}
		})
	}
}

func TestBlueBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Blue background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;44;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlueBackground(tt.args.Text); got != tt.want {
				t.Errorf("BlueBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BlueBackground() = ", got)
			}
		})
	}
}

func TestMagentaBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Magenta background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;45;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MagentaBackground(tt.args.Text); got != tt.want {
				t.Errorf("MagentaBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("MagentaBackground() = ", got)
			}
		})
	}
}

func TestCyanBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Cyan background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;46;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CyanBackground(tt.args.Text); got != tt.want {
				t.Errorf("CyanBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("CyanBackground() = ", got)
			}
		})
	}
}

func TestWhiteBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "White background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;47;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WhiteBackground(tt.args.Text); got != tt.want {
				t.Errorf("WhiteBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("WhiteBackground() = ", got)
			}
		})
	}
}
