package main

import (
	"bytes"
	"fmt"
	_ "image/jpeg"
	"net/http"

	"github.com/edot92/gotherimg/engine/views"
	"github.com/edot92/gotherimg/router"
	"github.com/go-ini/ini"

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
	gin.SetMode(gin.ReleaseMode)
	router.Router.Static("/static", "static")
	router.Router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/index")
	})
	router.Router.GET("/index", func(c *gin.Context) {
		dataRender := make(map[string]string)
		dataRender["PageTitle"] = "Template example: map"
		dataRender["Message"] = "Hello"
		dataRender["User"] = "World"
		buffer := new(bytes.Buffer)
		template.HalamanRender(dataRender, buffer)
		c.Writer.Write(buffer.Bytes())

	})
	cfg, err := ini.InsensitiveLoad("conf//app.conf")
	if err != nil {
		fmt.Println("file conf//app.conf notfound , Err:", err.Error())
		router.Router.Run(":8080")
	} else {
		valPort := cfg.Section("").Key("PORT_WEB").Value()
		if len(valPort) == 0 {
			fmt.Println("Key PORT_WEB not found")
			router.Router.Run(":8080")
		} else {
			fmt.Println("Key PORT_WEB  found value:" + valPort)
			router.Router.Run(":" + valPort)
		}
	}
}
