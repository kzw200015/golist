package assets

import "embed"

//go:embed frontend/dist
var efs embed.FS

func Get() embed.FS {
	return efs
}
