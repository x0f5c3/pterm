# TreePrinter

<!--
Replace all of the following strings with the current printer.
     tree Tree TreePrinter DefaultTree
-->

![TreePrinter Example](https://raw.githubusercontent.com/x0f5c3/pterm/master/_examples/tree/animation.svg)

<p align="center"><a href="https://github.com/x0f5c3/pterm/blob/master/_examples/tree/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
pterm.DefaultTree.WithRoot(pterm.NewTreeFromLeveledList(pterm.LeveledList{pterm.LeveledListItem{Level: 0, Text: "Hello, World!"}})).Render()
```

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultTree.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultTree.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                                           | Type                                                           |
| ---------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------- |
| [HorizontalString](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithHorizontalString)                   | string                                                         |
| [Indent](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithIndent)                                       | int                                                            |
| [Root](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithRoot)                                           | [TreeNode](https://pkg.go.dev/github.com/x0f5c3/pterm#TreeNode) |
| [TextStyle](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithTextStyle)                                 | [\*Style](https://pkg.go.dev/github.com/x0f5c3/pterm#Style)     |
| [TopRightCornerString](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithTopRightCornerString)           | string                                                         |
| [TopRightDownStringOngoing](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithTopRightDownStringOngoing) | string                                                         |
| [TreeStyle](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithTreeStyle)                                 | [\*Style](https://pkg.go.dev/github.com/x0f5c3/pterm#Style)     |
| [VerticalString](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithVerticalString)                       | string                                                         |
| [Writer](https://pkg.go.dev/github.com/x0f5c3/pterm#TreePrinter.WithWriter)                                       | io.Writer                                                      |

### Output functions

> This printer implements the interface [`RenderablePrinter`](https://github.com/x0f5c3/pterm/blob/master/interface_renderable_printer.go)

| Function  | Description        |
| --------- | ------------------ |
| Render()  | Prints to Terminal |
| Srender() | Returns a string   |

## Related

- [Override default printers](docs/customizing/override-default-printer.md)
