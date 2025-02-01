package bridge

import (
	"strings"

	"tractor.dev/toolkit-go/engine/fs"
)

type entry struct {
	IsDir bool
	Ctime int
	Mtime int
	Size  int
	Name  string
}

func fixPath(path string) string {
	p := strings.TrimLeft(path, "/")
	if p == "" {
		p = "."
	}
	return p
}

func (api *Bridge) Stat(path string) (*entry, error) {
	fi, err := fs.Stat(api.GetFS(), fixPath(path))
	if err != nil {
		return nil, err
	}
	return &entry{
		Name:  fi.Name(),
		Mtime: int(fi.ModTime().Unix()),
		IsDir: fi.IsDir(),
		Ctime: 0,
		Size:  int(fi.Size()),
	}, nil
}

func (api *Bridge) ReadFile(path string) ([]byte, error) {
	return fs.ReadFile(api.GetFS(), fixPath(path))
}

func (api *Bridge) ReadDir(path string) ([]entry, error) {
	dir, err := fs.ReadDir(api.GetFS(), fixPath(path))
	if err != nil {
		return nil, err
	}
	var entries []entry
	for _, e := range dir {
		fi, _ := e.Info()
		entries = append(entries, entry{
			Name:  fi.Name(),
			Mtime: int(fi.ModTime().Unix()),
			IsDir: fi.IsDir(),
			Ctime: 0,
			Size:  int(fi.Size()),
		})
	}
	return entries, nil
}

func (api *Bridge) WriteFile(path string, data []byte) error {
	return fs.WriteFile(api.GetFS(), fixPath(path), data, 0644)
}

func (api *Bridge) MakeDir(path string) error {
	return fs.MkdirAll(api.GetFS(), fixPath(path), 0744)
}
