package main

import (
	"os"

	"github.com/JosipVodafone/countbeat/cmd"

	_ "github.com/JosipVodafone/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
