package routers

import (
	"net/http"
	"web/urls"
	"web/utils"

	"github.com/gorilla/pat"
)

func GetRouter() *pat.Router {
	common := pat.New()
	fs := utils.FileSystemStatic{FileSystem: http.Dir("static")}
	common.PathPrefix(urls.StaticPath).Handler(
		http.StripPrefix(urls.StaticPath, http.FileServer(http.Dir("static"))))
	common.PathPrefix(urls.HomePath).Handler(http.FileServer(fs))
	return common
}
