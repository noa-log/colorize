/*
 * @Author: nijineko
 * @Date: 2025-06-07 11:05:59
 * @LastEditTime: 2025-06-07 11:12:48
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
