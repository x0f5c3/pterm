### interactive_confirm/demo

![Animation](https://raw.githubusercontent.com/x0f5c3/pterm/master/_examples/interactive_confirm/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/x0f5c3/pterm"
)

func main() {
	result, _ := pterm.DefaultInteractiveConfirm.Show()
	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", boolToText(result))
}

func boolToText(b bool) string {
	if b {
		return pterm.Green("Yes")
	}
	return pterm.Red("No")
}

```

</details>

