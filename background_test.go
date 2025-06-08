/*
 * @Author: nijineko
 * @Date: 2025-06-07 11:13:20
 * @LastEditTime: 2025-06-08 11:59:11
 * @LastEditors: nijineko
 * @Description: quick method for set background color
 * @FilePath: \colorize\background_test.go
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

func TestGrayBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Gray background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;100;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GrayBackground(tt.args.Text); got != tt.want {
				t.Errorf("GrayBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("GrayBackground() = ", got)
			}
		})
	}
}

func TestBrightRedBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bright Red background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;101;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrightRedBackground(tt.args.Text); got != tt.want {
				t.Errorf("BrightRedBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BrightRedBackground() = ", got)
			}
		})
	}
}

func TestBrightGreenBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bright Green background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;102;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrightGreenBackground(tt.args.Text); got != tt.want {
				t.Errorf("BrightGreenBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BrightGreenBackground() = ", got)
			}
		})
	}
}

func TestBrightYellowBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bright Yellow background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;103;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrightYellowBackground(tt.args.Text); got != tt.want {
				t.Errorf("BrightYellowBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BrightYellowBackground() = ", got)
			}
		})
	}
}

func TestBrightBlueBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bright Blue background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;104;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrightBlueBackground(tt.args.Text); got != tt.want {
				t.Errorf("BrightBlueBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BrightBlueBackground() = ", got)
			}
		})
	}
}

func TestBrightMagentaBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bright Magenta background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;105;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrightMagentaBackground(tt.args.Text); got != tt.want {
				t.Errorf("BrightMagentaBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BrightMagentaBackground() = ", got)
			}
		})
	}
}

func TestBrightCyanBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bright Cyan background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;106;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrightCyanBackground(tt.args.Text); got != tt.want {
				t.Errorf("BrightCyanBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BrightCyanBackground() = ", got)
			}
		})
	}
}

func TestBrightWhiteBackground(t *testing.T) {
	type args struct {
		Text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bright White background",
			args: args{
				Text: "Hello, Noa!",
			},
			want: "\x1b[0;107;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BrightWhiteBackground(tt.args.Text); got != tt.want {
				t.Errorf("BrightWhiteBackground() = %v, want %v", got, tt.want)
			} else {
				t.Log("BrightWhiteBackground() = ", got)
			}
		})
	}
}
