### center/demo

![Animation](https://raw.githubusercontent.com/x0f5c3/pterm/master/_examples/center/demo/animation.svg)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/x0f5c3/pterm"

func main() {
	pterm.DefaultCenter.Println("This text is centered!\nIt centeres the whole block by default.\nIn that way you can do stuff like this:")

	// Generate BigLetters
	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("PTerm")).Srender()
	pterm.DefaultCenter.Println(s) // Print BigLetters with the default CenterPrinter

	pterm.DefaultCenter.WithCenterEachLineSeparately().Println("This text is centered!\nBut each line is\ncentered\nseparately")
}

```

</details>

