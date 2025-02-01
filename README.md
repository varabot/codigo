# codigo

VSCode as a Go library. Embed an editor in your Go programs.

Set up `codigo.Workbench` with a terminal factory and filesystem (both of which can be virtual), then you can serve it as an HTTP handler to access your custom VSCode editor in the browser. Use with a webview window library to give the editor its own native window.

```go
func main() {
	cwd, _ := os.Getwd()
	fsys := workingpathfs.New(osfs.New(), cwd)

	wb := &codigo.Workbench{
		ProductConfiguration: product.Configuration{
			NameLong: "My Custom Editor",
		},
		FS: fsys,
	}

	log.Println("editor serving on :8080 ...")
	if err := http.ListenAndServe(":8080", wb); err != nil {
		log.Fatal(err)
	}

}

```

Let me know what else you'd like to customize from Go!

## License

MIT
