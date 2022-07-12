### style/demo

![Animation](https://raw.githubusercontent.com/x0f5c3/pterm/master/_examples/style/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/x0f5c3/pterm"

func main() {
	// Create styles as new variables
	primary := pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold)
	secondary := pterm.NewStyle(pterm.FgLightGreen, pterm.BgWhite, pterm.Italic)

	// Use created styles
	primary.Println("Hello, World!")
	secondary.Println("Hello, World!")
}

```

</details>

