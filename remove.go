/*
 * @Author: nijineko
 * @Date: 2025-06-07 10:59:07
 * @LastEditTime: 2025-06-07 11:00:07
 * @LastEditors: nijineko
 * @Description: remove color from text
 * @FilePath: \colorize\remove.go
 */
package colorize

import "regexp"

/**
 * @description: remove color from text
 * @param {string} Text colorized text that needs to be cleaned
 * @return {string} cleaned text without color codes
 */
func Remove(Text string) string {
	return regexp.MustCompile(`\x1b(\[.*?[@-~]|\].*?(\x07|\x1b\\))`).ReplaceAllString(Text, "")
}
