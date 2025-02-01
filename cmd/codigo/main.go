package main

import (
	"log"
	"net/http"
	"os"

	"github.com/btwiuse/codigo"
	"github.com/btwiuse/codigo/product"
	"github.com/btwiuse/codigo/upgrade"
	"github.com/btwiuse/codigo/version"
	"github.com/btwiuse/multicall"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
	"tractor.dev/toolkit-go/engine/fs/osfs"
	"tractor.dev/toolkit-go/engine/fs/workingpathfs"
)

var cmdRun multicall.RunnerFuncMap = map[string]multicall.RunnerFunc{
	// version info
	"version": version.Run,
	// binary upgrade
	"upgrade": upgrade.Run,
	// start cmd
	"start": Run,
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	err := cmdRun.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}

func Run(args []string) error {
	codigo.DownloadAndUnzipVSCode()

	cwd, _ := os.Getwd()
	fsys := workingpathfs.New(osfs.New(), cwd)

	wb := &codigo.Workbench{
		ProductConfiguration: product.Configuration{
			NameLong: "My Custom Editor",
		},
		FS: fsys,
	}

	var handler http.Handler = wb
	handler = utils.GinLoggerMiddleware(wb)

	return wtf.Serve(RELAY, handler)
}

var RELAY = EnvRELAY("https://ufo.k0s.io")

func EnvRELAY(s string) string {
	if relay := os.Getenv("RELAY"); relay != "" {
		return relay
	}
	return s
}
