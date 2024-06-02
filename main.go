package main

import (
	"github.com/newton-miku/nali/internal/constant"

	"github.com/newton-miku/nali/cmd"
	"github.com/newton-miku/nali/internal/config"

	_ "github.com/newton-miku/nali/internal/migration"
)

func main() {
	config.ReadConfig(constant.ConfigDirPath)
	cmd.Execute()
}
