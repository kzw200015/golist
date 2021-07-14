package handlers

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/kzw200015/go-list/args"
	"github.com/kzw200015/go-list/types"
)

func ListPath() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := filepath.Clean(c.DefaultQuery("path", "/"))

		err := checkPath(path)
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		pathInfo, err := getInfo(path)
		if err != nil {
			c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}

		c.JSON(200, pathInfo)
	}

}

func Down(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := filepath.Clean(c.Query("path"))
		c.Request.URL.Path = filepath.ToSlash(filepath.Join("/file", path))
		router.HandleContext(c)
	}
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

func checkPath(path string) error {
	p1, err := filepath.Abs(filepath.Join(args.GetSrcPath(), path))
	if err != nil {
		return err
	}

	p2, err := filepath.Abs(args.GetSrcPath())
	if err != nil {
		return err
	}

	if !strings.HasPrefix(p1, p2) {
		return errors.New("path no permission")
	}

	return nil
}
