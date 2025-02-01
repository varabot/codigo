package vscode

import (
	"archive/zip"
	"bytes"
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	_ "embed"

	"github.com/btwiuse/codigo/bridge"
	"github.com/btwiuse/codigo/product"
	"github.com/btwiuse/codigo/zipfs"
	"github.com/btwiuse/wsconn"
	"tractor.dev/toolkit-go/duplex/codec"
	"tractor.dev/toolkit-go/duplex/fn"
	"tractor.dev/toolkit-go/duplex/mux"
	"tractor.dev/toolkit-go/duplex/talk"
	"tractor.dev/toolkit-go/engine/fs"
	"tractor.dev/toolkit-go/engine/fs/workingpathfs"
)

//go:embed extension assets
var embedded embed.FS

var vscodeReader *zip.Reader

func init() {
	b, err := embedded.ReadFile("assets/vscode-web.zip")
	if err != nil {
		panic(err)
	}
	vscodeReader, err = zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		panic(err)
	}
}

type URIComponents struct {
	Scheme    string `json:"scheme"`
	Authority string `json:"authority,omitempty"`
	Path      string `json:"path,omitempty"`
	Query     string `json:"query,omitempty"`
	Fragment  string `json:"fragment,omitempty"`
}

type Workbench struct {
	ProductConfiguration        product.Configuration `json:"productConfiguration"`
	AdditionalBuiltinExtensions []*URIComponents      `json:"additionalBuiltinExtensions,omitempty"`
	FolderURI                   *URIComponents        `json:"folderUri,omitempty"`
	InitialColorTheme           ColorScheme           `json:"initialColorTheme,omitempty"`

	FS fs.FS `json:"-"`
}

func (wb *Workbench) GetFS() fs.FS {
	return wb.FS
}

type ColorScheme struct {
	ThemeType string            `json:"themeType"` // "dark" | "light" | "hcLight" | "hcDark"
	Colors    map[string]string `json:"colors,omitempty"`
}

func (wb *Workbench) setColorScheme(t string) {
	wb.InitialColorTheme.ThemeType = t
}

func (wb *Workbench) ensureExtension(r *http.Request) {
	origin := r.Header.Get("Origin")

	o, err := url.Parse(origin)
	if err != nil {
		log.Println(err)
		return
	}

	abe := &URIComponents{
		Scheme:    o.Scheme,
		Authority: o.Host,
		Path:      "/extension",
	}

	foundExtension := false
	for i, e := range wb.AdditionalBuiltinExtensions {
		if e.Path == "/extension" {
			wb.AdditionalBuiltinExtensions[i] = abe
			foundExtension = true
			break
		}
	}

	if !foundExtension {
		wb.AdditionalBuiltinExtensions = append(wb.AdditionalBuiltinExtensions, abe)
	}
}

func (wb *Workbench) ensureFolder() {
	if wb.FolderURI == nil {
		wb.FolderURI = &URIComponents{
			Scheme: "hostfs",
			Path:   "/",
		}
	}
}

func (wb *Workbench) handleBridge(w http.ResponseWriter, r *http.Request) {
	conn, err := wsconn.Wrconn(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	sess := mux.New(conn)
	defer sess.Close()

	peer := talk.NewPeer(sess, codec.CBORCodec{})
	peer.Handle("vscode/", fn.HandlerFrom(&bridge.Bridge{wb}))
	peer.Respond()
}

func (wb *Workbench) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// wb.setColorScheme("dark")
	wb.ensureFolder()

	mux := http.NewServeMux()

	mux.Handle("/workbench.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wb.ensureExtension(r)
		w.Header().Add("content-type", "application/json")
		enc := json.NewEncoder(w)
		if err := enc.Encode(wb); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

	mux.Handle("/extension/", http.FileServerFS(embedded))

	fsys := workingpathfs.New(zipfs.New(vscodeReader), "dist")
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFileFS(w, r, embedded, "assets/index.html")
			return
		}

		if r.URL.Path == "/bridge.js" {
			http.ServeFileFS(w, r, embedded, "assets/bridge.js")
			return
		}

		http.FileServerFS(fsys).ServeHTTP(w, r)
	}))

	mux.HandleFunc("/bridge", wb.handleBridge)

	mux.ServeHTTP(w, r)
}
