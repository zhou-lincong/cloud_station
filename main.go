package main

import (
	"fmt"

	"github.com/zhou-lincong/cloud_station/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
