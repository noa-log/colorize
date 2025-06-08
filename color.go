/*
 * @Author: nijineko
 * @Date: 2025-06-07 10:44:04
 * @LastEditTime: 2025-06-08 11:42:11
 * @LastEditors: nijineko
 * @Description: define color constants
 * @FilePath: \colorize\color.go
 */
package colorize

// define default constants
const (
	DEFAULT_CONFIG     = 0         // Default configuration
	DEFAULT_BACKGROUND = 0         // Default background color
	DEFAULT_TEXT       = TextWhite // Default text color
)

// define color constants
const (
	ConfigDefault   = iota     // Default configuration
	ConfigBold                 // Bold text
	ConfigDim                  // Dim text
	ConfigItalic               // Italic text
	ConfigUnderline            // Underlined text
	ConfigBlinking             // Blinking text
	ConfigReverse   = iota + 1 // Reversed colors
	ConfigHidden               // Hidden text
)

// define background color constants
const (
	BackgroundBlack         = iota + 40 // Black background
	BackgroundRed                       // Red background
	BackgroundGreen                     // Green background
	BackgroundYellow                    // Yellow background
	BackgroundBlue                      // Blue background
	BackgroundMagenta                   // Magenta background
	BackgroundCyan                      // Cyan background
	BackgroundWhite                     // White background
	BackgroundGray          = iota + 92 // Gray background (bright black)
	BackgroundBrightRed                 // Bright red background
	BackgroundBrightGreen               // Bright green background
	BackgroundBrightYellow              // Bright yellow background
	BackgroundBrightBlue                // Bright blue background
	BackgroundBrightMagenta             // Bright magenta background
	BackgroundBrightCyan                // Bright cyan background
	BackgroundBrightWhite               // Bright white background
)

// define text color constants
const (
	TextBlack         = iota + 30 // Black text
	TextRed                       // Red text
	TextGreen                     // Green text
	TextYellow                    // Yellow text
	TextBlue                      // Blue text
	TextMagenta                   // Magenta text
	TextCyan                      // Cyan text
	TextWhite                     // White text
	TextGray          = iota + 82 // Gray text (bright black)
	TextBrightRed                 // Bright red text
	TextBrightGreen               // Bright green text
	TextBrightYellow              // Bright yellow text
	TextBrightBlue                // Bright blue text
	TextBrightMagenta             // Bright magenta text
	TextBrightCyan                // Bright cyan text
	TextBrightWhite               // Bright white text
)
