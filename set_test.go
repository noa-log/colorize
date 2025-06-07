/*
 * @Author: nijineko
 * @Date: 2025-06-07 11:18:45
 * @LastEditTime: 2025-06-07 11:26:30
 * @LastEditors: nijineko
 * @Description: set color for text
 * @FilePath: \colorize\set_test.go
 */
package colorize

import "testing"

func TestSet(t *testing.T) {
	type args struct {
		Text            string
		Config          int
		BackgroundColor int
		TextColor       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Text black",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       TextBlack,
			},
			want: "\x1b[0;0;30mHello, Noa!\x1b[0m",
		},
		{
			name: "Text red",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       TextRed,
			},
			want: "\x1b[0;0;31mHello, Noa!\x1b[0m",
		},
		{
			name: "Text green",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       TextGreen,
			},
			want: "\x1b[0;0;32mHello, Noa!\x1b[0m",
		},
		{
			name: "Text yellow",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       TextYellow,
			},
			want: "\x1b[0;0;33mHello, Noa!\x1b[0m",
		},
		{
			name: "Text blue",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       TextBlue,
			},
			want: "\x1b[0;0;34mHello, Noa!\x1b[0m",
		},
		{
			name: "Text magenta",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       TextMagenta,
			},
			want: "\x1b[0;0;35mHello, Noa!\x1b[0m",
		},
		{
			name: "Text cyan",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       TextCyan,
			},
			want: "\x1b[0;0;36mHello, Noa!\x1b[0m",
		},
		{
			name: "Text white",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       TextWhite,
			},
			want: "\x1b[0;0;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Background black",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: BackgroundBlack,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[0;40;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Background red",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: BackgroundRed,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[0;41;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Background green",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: BackgroundGreen,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[0;42;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Background yellow",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: BackgroundYellow,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[0;43;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Background blue",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: BackgroundBlue,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[0;44;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Background magenta",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: BackgroundMagenta,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[0;45;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Background cyan",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: BackgroundCyan,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[0;46;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Background white",
			args: args{
				Text:            "Hello, Noa!",
				Config:          DEFAULT_CONFIG,
				BackgroundColor: BackgroundWhite,
				TextColor:       TextBlack,
			},
			want: "\x1b[0;47;30mHello, Noa!\x1b[0m",
		},
		{
			name: "Config bold",
			args: args{
				Text:            "Hello, Noa!",
				Config:          ConfigBold,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[1;0;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Config dim",
			args: args{
				Text:            "Hello, Noa!",
				Config:          ConfigDim,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[2;0;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Config italic",
			args: args{
				Text:            "Hello, Noa!",
				Config:          ConfigItalic,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[3;0;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Config underline",
			args: args{
				Text:            "Hello, Noa!",
				Config:          ConfigUnderline,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[4;0;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Config blinking",
			args: args{
				Text:            "Hello, Noa!",
				Config:          ConfigBlinking,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[5;0;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Config reverse",
			args: args{
				Text:            "Hello, Noa!",
				Config:          ConfigReverse,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[7;0;37mHello, Noa!\x1b[0m",
		},
		{
			name: "Config hidden",
			args: args{
				Text:            "Hello, Noa!",
				Config:          ConfigHidden,
				BackgroundColor: DEFAULT_BACKGROUND,
				TextColor:       DEFAULT_TEXT,
			},
			want: "\x1b[8;0;37mHello, Noa!\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Set(tt.args.Text, tt.args.Config, tt.args.BackgroundColor, tt.args.TextColor); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			} else {
				t.Log("Set() = ", got)
			}
		})
	}
}
