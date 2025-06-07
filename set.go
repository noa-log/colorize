/*
 * @Author: nijineko
 * @Date: 2025-06-07 10:53:28
 * @LastEditTime: 2025-06-07 11:00:21
 * @LastEditors: nijineko
 * @Description: set color for text
 * @FilePath: \colorize\set.go
 */
package colorize

import "fmt"

/**
 * @description: set color for text
 * @param {string} Text Text that needs color added
 * @param {int} Config
 * @param {int} BackgroundColor
 * @param {int} TextColor
 * @return {string} formatted text with color
 */
func Set(Text string, Config int, BackgroundColor int, TextColor int) string {
	return fmt.Sprintf("\x1b[%d;%d;%dm%s\x1b[0m",
		Config,
		BackgroundColor,
		TextColor,
		Text,
	)
}
