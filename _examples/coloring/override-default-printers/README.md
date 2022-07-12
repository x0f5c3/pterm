# coloring/override-default-printers

![Animation](animation.svg)

```go
package main

import "github.com/x0f5c3/pterm"

func main() {
	// Print default error.
	pterm.Error.Println("This is the default Error")

	// Customize default error.
	pterm.Error.Prefix = pterm.Prefix{
		Text:  "OVERRIDE",
		Style: pterm.NewStyle(pterm.BgCyan, pterm.FgRed),
	}

	// Print new default error.
	pterm.Error.Println("This is the default Error after the prefix was overridden")
}

```
