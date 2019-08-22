cmd
===
More convenient commands builder base on **Cobra**.The key feature of the package:

-  use filesystem path like format to route the subcommands

## Installation

````bash
$: go get github.com/x-mod/cmd
````

## Dependence

- [cobra](https://github.com/spf13/cobra)

## Quick Start

### Only Root Command

replace default Root Command settings. `cmd.Parent("")` means replace the default Root Command with the new command.

````go
import (
	"fmt"

	"github.com/x-mod/cmd"
)

func main() {
	cmd.Add(
		cmd.Name("root"),
		cmd.Main(Main),
	)
	cmd.Execute()
}

func Main(c *cmd.Command, args []string) error {
	fmt.Println("my root command running ...")
	return nil
}
````

run the code in bash:

````bash
$: go run main.go
my root command running ...
````

### Sub Commands

sub commands routing rules:

- `cmd.Parent("/")`  level 1
- `cmd.Parent("/foo/bar")` level 3

subcommand's `cmd.Parent("/command/path")` must be setting.

````go
import (
	"fmt"

	"github.com/x-mod/cmd"
)

func main() {
	cmd.Add(
		cmd.Parent("/foo/bar"),
		cmd.Name("v1"),
		cmd.Main(V1),
	).PersistentFlags().StringP("parameter", "p", "test", "flags usage")
	cmd.Version("version string")
	cmd.Execute()
}

func V1(c *cmd.Command, args []string) error {
	fmt.Println("V1 called")
	return nil
}
````

run the code in bash:

````bash
$: go run main.go foo bar v1
V1 called
````


