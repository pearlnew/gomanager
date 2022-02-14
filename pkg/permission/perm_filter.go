package permission

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/pearlnew/gomanager/pkg/models"
	"github.com/pearlnew/gomanager/pkg/utils/option"
	"github.com/pearlnew/gomanager/pkg/utils/session"
)

func MiddleWare(req *restful.Request, resp *restful.Response) (*models.User, error){
	// 前端网页根据seesion即可获取
	sess := session.GlobalSession.SessionStart(resp, req.Request)
	userSession := sess.Get(session.SESSION_USER)
	if userSession == nil {
		return nil, fmt.Errorf("用户没有登陆")
	}
	user, ok := userSession.(models.User)
	if ok {
		return &user, nil
	} else {
		return &user, fmt.Errorf("用户没有登陆")
	}
}

func CSRFProtect(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	csrfHeader := req.Request.Header.Get(option.GlobalOption.CSRFName)
	if csrfHeader != option.GlobalOption.CSRFToken {
		resp.WriteErrorString(403, "csrf error!")
		return
	}
	chain.ProcessFilter(req, resp)
}

func RequireLogin(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	user, err := MiddleWare(req, resp)
	if err != nil {
		resp.WriteError(401, err)
		return
	}
	if user.UserID == 0 {
		resp.WriteErrorString(401, "用户未登陆")
		return
	}
	chain.ProcessFilter(req, resp)
}

func RequireAdmin(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	user, err := MiddleWare(req, resp)
	if err != nil {
		resp.WriteError(401, err)
		return
	}
	if user.UserID == 0 {
		resp.WriteErrorString(401, "用户未登陆")
		return
	}
	if user.IsAdmin == 0 {
		resp.WriteError(403, fmt.Errorf("您没有管理员权限！"))
		return
	}
	chain.ProcessFilter(req, resp)
}