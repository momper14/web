package router

import (
	"net/http"
	"web/app/url"
	"web/app/util"

	"github.com/gorilla/pat"
)

// GetRouter returns all routers
func GetRouter() *pat.Router {
	common := pat.New()
	fs := util.FileSystemStatic{FileSystem: http.Dir("static")}
	common.PathPrefix(url.StaticPath).Handler(
		http.StripPrefix(url.StaticPath, http.FileServer(http.Dir("static"))))
	common.PathPrefix(url.HomePath).Handler(http.FileServer(fs))
	return common
}
