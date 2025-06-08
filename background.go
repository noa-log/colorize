/*
 * @Author: nijineko
 * @Date: 2025-06-07 11:13:20
 * @LastEditTime: 2025-06-07 11:13:26
 * @LastEditors: nijineko
 * @Description: quick method for set background color
 * @FilePath: \colorize\background.go
 */
package colorize

/**
 * @description: Set background color to black
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BlackBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBlack, DEFAULT_TEXT)
}

/**
 * @description: Set background color to red
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func RedBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundRed, DEFAULT_TEXT)
}

/**
 * @description: Set background color to green
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func GreenBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundGreen, DEFAULT_TEXT)
}

/**
 * @description: Set background color to yellow
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func YellowBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundYellow, DEFAULT_TEXT)
}

/**
 * @description: Set background color to blue
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BlueBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBlue, DEFAULT_TEXT)
}

/**
 * @description: Set background color to magenta
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func MagentaBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundMagenta, DEFAULT_TEXT)
}

/**
 * @description: Set background color to cyan
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func CyanBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundCyan, DEFAULT_TEXT)
}

/**
 * @description: Set background color to white
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func WhiteBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundWhite, DEFAULT_TEXT)
}

/**
 * @description: Set background color to gray
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func GrayBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundGray, DEFAULT_TEXT)
}

/**
 * @description: Set background color to bright red
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BrightRedBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBrightRed, DEFAULT_TEXT)
}

/**
 * @description: Set background color to bright green
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BrightGreenBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBrightGreen, DEFAULT_TEXT)
}

/**
 * @description: Set background color to bright yellow
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BrightYellowBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBrightYellow, DEFAULT_TEXT)
}

/**
 * @description: Set background color to bright blue
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BrightBlueBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBrightBlue, DEFAULT_TEXT)
}

/**
 * @description: Set background color to bright magenta
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BrightMagentaBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBrightMagenta, DEFAULT_TEXT)
}

/**
 * @description: Set background color to bright cyan
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BrightCyanBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBrightCyan, DEFAULT_TEXT)
}

/**
 * @description: Set background color to bright white
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BrightWhiteBackground(Text string) string {
	return Set(Text, DEFAULT_CONFIG, BackgroundBrightWhite, DEFAULT_TEXT)
}
