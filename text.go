/*
 * @Author: nijineko
 * @Date: 2025-06-07 11:05:59
 * @LastEditTime: 2025-06-08 11:45:47
 * @LastEditors: nijineko
 * @Description: quick method for set text color
 * @FilePath: \colorize\text.go
 */
package colorize

/**
 * @description: Set text color to black
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BlackText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBlack)
}

/**
 * @description: Set text color to red
 * @param {string} Text
 * @return {string} formatted text with red color
 */
func RedText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextRed)
}

/**
 * @description: Set text color to green
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func GreenText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextGreen)
}

/**
 * @description: Set text color to yellow
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func YellowText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextYellow)
}

/**
 * @description: Set text color to blue
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BlueText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBlue)
}

/**
 * @description: Set text color to magenta
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func MagentaText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextMagenta)
}

/**
 * @description: Set text color to cyan
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func CyanText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextCyan)
}

/**
 * @description: Set text color to white
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func WhiteText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextWhite)
}

/**
 * @description: Set text color to gray
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func GrayText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextGray)
}

/**
 * @description: Set text color to bright red
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BrightRedText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBrightRed)
}

/**
 * @description: Set text color to bright green
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BrightGreenText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBrightGreen)
}

/**
 * @description: Set text color to bright yellow
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BrightYellowText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBrightYellow)
}

/**
 * @description: Set text color to bright blue
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BrightBlueText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBrightBlue)
}

/**
 * @description: Set text color to bright magenta
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BrightMagentaText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBrightMagenta)
}

/**
 * @description: Set text color to bright cyan
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BrightCyanText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBrightCyan)
}

/**
 * @description: Set text color to bright white
 * @param {string} Text
 * @return {string} formatted text with black color
 */
func BrightWhiteText(Text string) string {
	return Set(Text, ConfigDefault, DEFAULT_BACKGROUND, TextBrightWhite)
}
