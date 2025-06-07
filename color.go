/*
 * @Author: nijineko
 * @Date: 2025-06-07 10:44:04
 * @LastEditTime: 2025-06-07 10:56:01
 * @LastEditors: nijineko
 * @Description: define color constants
 * @FilePath: \colorize\color.go
 */
package colorize

// define color constants
const (
	ConfigDefault   = iota // Default configuration
	ConfigBold             // Bold text
	ConfigItalic           // Italic text
	ConfigUnderline        // Underlined text
	ConfigBlinking         // Blinking text
	ConfigReverse          // Reversed colors
	ConfigHidden           // Hidden text
)

// define background color constants
const (
	BackgroundBlack   = iota + 40 // Black background
	BackgroundRed                 // Red background
	BackgroundGreen               // Green background
	BackgroundYellow              // Yellow background
	BackgroundBlue                // Blue background
	BackgroundMagenta             // Magenta background
	BackgroundCyan                // Cyan background
	BackgroundWhite               // White background
)

// define text color constants
const (
	TextBlack   = iota + 30 // Black text
	TextRed                 // Red text
	TextGreen               // Green text
	TextYellow              // Yellow text
	TextBlue                // Blue text
	TextMagenta             // Magenta text
	TextCyan                // Cyan text
	TextWhite               // White text
)
