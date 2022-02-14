//go:generate statik -src=./../web/dist -dest=./../ -f
package main

import (
	"fmt"
	"github.com/pearlnew/gomanager/pkg/utils/logging"
	"github.com/pearlnew/gomanager/pkg/utils/option"
	"github.com/pearlnew/gomanager/pkg/utils/session"
	"github.com/pearlnew/gomanager/service/router"
	"net/http"
)

func main() {
	option.GlobalOption.InitOptions()
	logging.InitLogger(option.GlobalOption.LogFile)
	session.InitSession(option.GlobalOption.SessionName)
	routerHandler := router.Register()
	server := &http.Server{Addr: fmt.Sprintf(":%d", option.GlobalOption.Port), Handler: routerHandler}
	////if err := server.ListenAndServeTLS("./config/ca.crt","./config/ca.key"); err != nil && err != http.ErrServerClosed {
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logging.Log.Fatal(err)
	}
}

