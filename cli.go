package main

import "github.com/urfave/cli"

type CLI struct {
	driver *cli.App
}

func (c *CLI) Run(args []string) error {
	return c.driver.Run(args)
}
