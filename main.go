package main

import (
	"bytes"
	_ "image/jpeg"
	"net/http"

	"github.com/edot92/gotherimg/engine/views"
	"github.com/edot92/gotherimg/router"
	// assetfs "github.com/elazarl/go-bindata-assetfs"/* comment jika ingin statis binary*/
	"github.com/gin-gonic/gin"
)

type ResponseJSON struct {
	status  string
	data    interface{}
	message string
}

type binaryFileSystem struct {
	fs http.FileSystem
}

/* comment jika ingin statis binary
func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	// viewsBinary lokasi url go
	fs := &assetfs.AssetFS{viewsBinary.Asset, viewsBinary.AssetDir, viewsBinary.AssetInfo, root}
	return &binaryFileSystem{
		fs,
	}
}
*/
func main() {
	// b := BinaryFileSystem("static")
	// router.Router.Use(static.Serve("/static", b))/* comment jika ingin statis binary*/
	router.Router.Static("/static", "static")
	router.Router.GET("/index", func(c *gin.Context) {
		dataRender := make(map[string]string)
		dataRender["PageTitle"] = "Template example: map"
		dataRender["Message"] = "Hello"
		dataRender["User"] = "World"
		buffer := new(bytes.Buffer)
		template.HalamanRender(dataRender, buffer)
		c.Writer.Write(buffer.Bytes())

	})

	router.Router.Run(":8081")
}
