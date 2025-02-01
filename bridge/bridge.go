package bridge

import (
	_ "embed"
	"io/fs"
)

type Bridge struct {
	Workbench
}

type Workbench interface {
	GetFS() fs.FS
}
