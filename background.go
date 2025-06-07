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
	return Set(Text, ConfigDefault, BackgroundBlack, DEFAULT_TEXT)
}

/**
 * @description: Set background color to red
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func RedBackground(Text string) string {
	return Set(Text, ConfigDefault, BackgroundRed, DEFAULT_TEXT)
}

/**
 * @description: Set background color to green
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func GreenBackground(Text string) string {
	return Set(Text, ConfigDefault, BackgroundGreen, DEFAULT_TEXT)
}

/**
 * @description: Set background color to yellow
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func YellowBackground(Text string) string {
	return Set(Text, ConfigDefault, BackgroundYellow, DEFAULT_TEXT)
}

/**
 * @description: Set background color to blue
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func BlueBackground(Text string) string {
	return Set(Text, ConfigDefault, BackgroundBlue, DEFAULT_TEXT)
}

/**
 * @description: Set background color to magenta
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func MagentaBackground(Text string) string {
	return Set(Text, ConfigDefault, BackgroundMagenta, DEFAULT_TEXT)
}

/**
 * @description: Set background color to cyan
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func CyanBackground(Text string) string {
	return Set(Text, ConfigDefault, BackgroundCyan, DEFAULT_TEXT)
}

/**
 * @description: Set background color to white
 * @param {string} Text
 * @return {string} formatted text with black background
 */
func WhiteBackground(Text string) string {
	return Set(Text, ConfigDefault, BackgroundWhite, DEFAULT_TEXT)
}
