// Package main
/*
Copyright © 2022 Jose Epifanio jose.epifanio90@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"github.com/JEpifanio90/bulldog-cli/internal/commands/list"
	"github.com/pterm/pterm"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "bulldog",
		Usage: "is a CLI that empowers developers by giving them full control over their cloud accounts and pipelines",
		Commands: []*cli.Command{
			{
				Name:        "tenant",
				Aliases:     []string{"tnt"},
				Usage:       "...",
				Subcommands: []*cli.Command{&list.Command},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		pterm.Error.Println(fmt.Errorf("bulldog main: %v", err))
	}
}
