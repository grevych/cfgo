package main

import (
	"flag"
	"fmt"

	"github.com/grevych/cfgo"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "", "Path to the configuration directory.")
	cfgo.Load(&cfgo.Cfg{
		Path:  configPath,
		Scope: "main",
	})
}

func main() {
	svcCfg := cfgo.Get("services.main")
	fmt.Printf("%+v", svcCfg)
}
