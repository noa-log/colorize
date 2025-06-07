# Colorize
A Golang package to add colors to your terminal output.

# Installation
```bash
go get -u github.com/noa-log/colorize
```

# Quick Start
```go
package main

import (
    "fmt"
    "github.com/noa-log/colorize"
)

func main() {
    // Add color to text
    fmt.Println(colorize.BlueText("This is blue text"))

    // Add background color to text
    fmt.Println(colorize.BlueBackground("This is text with a blue background"))

    // Set custom text color, background color, and style
    fmt.Println(colorize.Set("This is custom styled text", colorize.ConfigBold, colorize.BackgroundBlue, colorize.TextWhite))
}
```

# Features
- All available colors and styles are defined as constants.
- Provides methods to quickly set text color and background color.
- Supports custom combinations of colors and styles.
- Uses ANSI escape sequences, compatible with most terminals.
- 100% unit test coverage.

# License
This project is open-sourced under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0). Please comply with the terms when using it.