package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/x-mod/cmd"
)

func main() {
	cmd.ExitCode(true)
	cmd.Add(
		cmd.Name("root"),
		cmd.Main(RootMain),
	).PersistentFlags().StringP("parameter", "p", "test", "flags usage")
	cmd.Execute()
}

func RootMain(c *cmd.Command, args []string) error {
	fmt.Println("$HOME=", viper.GetString("HOME"))
	fmt.Println("my root command running ...parameter:", viper.GetString("parameter"))
	return fmt.Errorf("error return")
}
