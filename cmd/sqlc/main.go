package main

import "github.com/sqlc-dev/sqlc/pkg/cli"

func main() {
	cli.Run([]string{"generate", "-f", "sqlc.yaml"})
}
