package page

import (
	"fmt"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
	"path"
	"runtime"
)

var _pageMap = map[string]string{}

func init()  {
	_, currentFile, _, _ := runtime.Caller(0) // ignore error
	pageDir := path.Join(path.Dir(currentFile))
	fmt.Println(pageDir)
	box := packr.NewBox(pageDir)
	box.Walk(func(s string, file packd.File) error {
		_pageMap[s]=file.String()
		return nil
	})
}

func GetPage(name string) string {
	return _pageMap[name]
}
