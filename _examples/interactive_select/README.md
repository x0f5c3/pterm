### interactive_select/demo

![Animation](https://raw.githubusercontent.com/x0f5c3/pterm/master/_examples/interactive_select/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"fmt"

	"github.com/x0f5c3/pterm"
)

func main() {
	var options []string

	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}

```

</details>

