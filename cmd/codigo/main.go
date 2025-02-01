package main

import (
	"log"
	"net/http"
	"os"

	"github.com/btwiuse/codigo"
	"github.com/btwiuse/codigo/product"
	"github.com/webteleport/utils"
	"github.com/webteleport/wtf"
	"tractor.dev/toolkit-go/engine/fs/osfs"
	"tractor.dev/toolkit-go/engine/fs/workingpathfs"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	cwd, _ := os.Getwd()
	fsys := workingpathfs.New(osfs.New(), cwd)

	wb := &vscode.Workbench{
		ProductConfiguration: product.Configuration{
			NameLong: "My Custom Editor",
		},
		FS: fsys,
	}

	var handler http.Handler = wb
	handler = utils.GinLoggerMiddleware(wb)

	/*
		port := getPortOrDefault(":8080")
		log.Println("serving on", port)
		if err := http.ListenAndServe(port, handler); err != nil {
			log.Fatal(err)
		}
	*/

	err := wtf.Serve(RELAY, handler)
	if err != nil {
		log.Fatalln(err)
	}

}

var RELAY = EnvRELAY("https://ufo.k0s.io")

func EnvRELAY(s string) string {
	if relay := os.Getenv("RELAY"); relay != "" {
		return relay
	}
	return s
}

func getPortOrDefault(d string) string {
	port := os.Getenv("PORT")
	if port == "" {
		return d
	}
	return ":" + port
}
