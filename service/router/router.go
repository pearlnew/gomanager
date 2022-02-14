package router

import (
	"github.com/emicklei/go-restful"
	"github.com/pearlnew/gomanager/pkg/permission"
	"github.com/pearlnew/gomanager/pkg/utils/logging"
	"github.com/pearlnew/gomanager/service/api"
	_ "github.com/pearlnew/gomanager/statik"
	"github.com/rakyll/statik/fs"
	"net/http"
	"strings"
	"sync"
)

var statikFS http.FileSystem

func init() {
	var err error
	statikFS, err = fs.New()
	if err != nil {
		logging.Log.Fatal("web file not found!")
	}
}

type Handler struct {
	*restful.Container
	extensionWebService *restful.WebService
	sync.RWMutex
}

func Register() *Handler {
	logging.Log.Infof("register web url ......")
	h := &Handler{
		Container: restful.NewContainer(),
	}
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders: []string{"appid"},
		AllowedHeaders: []string{"Content-Type","Accept", "appid", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: true,
		Container: h.Container,
	}
	h.Container.Filter(cors.Filter)
	h.Container.Filter(h.Container.OPTIONSFilter)

	registWeb(h.Container)
	registComAPI(h.Container)
	registUserAPI(h.Container)

	return h
}

func registWeb(container *restful.Container) {
	logging.Log.Infof("adding web page")
	t := http.FileServer(statikFS)
	container.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".html") {
			r.URL.Path = "/"
		}
		t.ServeHTTP(w, r)
	}))
}

func registComAPI(container *restful.Container) {
	proxy := new(restful.WebService)
	proxy.Filter(restful.NoBrowserCacheFilter)
	proxy.Consumes(restful.MIME_JSON, "application/json-patch+json", "text/plain").
		Produces(restful.MIME_JSON, "application/json-patch+json", "text/plain").
		Path("/api/login")
	proxy.Route(proxy.POST("/").To(api.Login))
	container.Add(proxy)
}

func registUserAPI(container *restful.Container) {
	proxy := new(restful.WebService)
	proxy.Filter(restful.NoBrowserCacheFilter)
	proxy.Consumes(restful.MIME_JSON, "application/json-patch+json", "text/plain").
		Produces(restful.MIME_JSON, "application/json-patch+json", "text/plain").
		Path("/api/user").Filter(permission.RequireLogin)
	proxy.Route(proxy.GET("/{userid:\\d}").To(api.GetUserDetail))
	proxy.Route(proxy.POST("/{userid:\\d}").To(api.EditUser))
	proxy.Route(proxy.GET("/list").To(api.GetUserList))
	proxy.Route(proxy.POST("/add").To(api.AddUser))
	container.Add(proxy)
}

func registOtherAPI(container *restful.Container) {

}