# Colorize
一个Golang包，为你的终端输出添加颜色。

# 安装
```bash
go get -u github.com/noa-log/colorize
```

# 快速开始
```go
package main

import (
    "fmt"
    "github.com/noa-log/colorize"
)

func main() {
	// 为文本添加颜色
	fmt.Println(colorize.BlueText("这是蓝色文本"))

	// 为文本添加背景色
	fmt.Println(colorize.BlueBackground("这是蓝色背景文本"))

	// 为文本添加自定义颜色，背景色和样式
	fmt.Println(colorize.Set("这是自定义颜色文本", colorize.ConfigBold, colorize.BackgroundBlue, colorize.TextWhite))
}
```

# 功能
- 使用常量定义了所有可用的颜色和样式。
- 提供方法来快速的设置文本颜色、背景颜色。
- 支持自定义颜色和样式组合。
- 使用 ANSI 转义序列，兼容大多数终端。
- 100% 单元测试覆盖。

# 许可
本项目基于[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)协议开源。使用时请遵守协议的条款。