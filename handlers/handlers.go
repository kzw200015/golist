package handlers

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/kzw200015/go-list/args"
	"github.com/kzw200015/go-list/types"
)

func ListPath(c *gin.Context) {
	destPath := c.Query("path")
	if destPath == "" {
		destPath = "/"
	}

	pathInfo, err := getInfo(destPath)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, pathInfo)
}

func getInfo(path string) (types.PathInfo, error) {
	dirEntries, err := os.ReadDir(filepath.Join(args.GetSrcPath(), path))
	if err != nil {
		return types.PathInfo{}, err
	}

	var pathInfo types.PathInfo

	pathInfo.Path = formatPath(path)

	for _, dirEntry := range dirEntries {
		pathItemInfo := types.PathItemInfo{
			Name:  dirEntry.Name(),
			Path:  formatPath(path, dirEntry.Name()),
			IsDir: dirEntry.IsDir(),
		}

		if !dirEntry.IsDir() {
			s, err := mimetype.DetectFile(filepath.Join(args.GetSrcPath(), path, dirEntry.Name()))
			if err != nil {
				return types.PathInfo{}, err
			}
			pathItemInfo.MimeType = s.String()
		}

		pathInfo.Items = append(pathInfo.Items, pathItemInfo)
	}

	sort.SliceStable(pathInfo.Items, func(i, j int) bool { return pathInfo.Items[i].IsDir && !pathInfo.Items[j].IsDir })

	pathInfo.Items = append([]types.PathItemInfo{{
		Name:  "..",
		Path:  formatPath(filepath.Dir(path)),
		IsDir: true,
	}}, pathInfo.Items...)

	return pathInfo, nil
}

func formatPath(path ...string) string {
	path = append([]string{"/"}, path...)
	return filepath.ToSlash(filepath.Join(path...))
}
