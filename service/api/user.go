package api

import (
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/pearlnew/gomanager/pkg/models"
	"github.com/pearlnew/gomanager/pkg/utils"
	"github.com/pearlnew/gomanager/pkg/utils/session"
	"io/ioutil"
	"strconv"
)

func Login(request *restful.Request, response *restful.Response) {
	user := models.User{UserID: 1}
	u, err := models.UserSearch(&user)
	ret := map[string]interface{}{"userInfo": u}
	sess := session.GlobalSession.SessionStart(response, request.Request)
	sess.Set(session.SESSION_USER, user)
	utils.RespondStruct(response, 200, ret, err)
}

func AddUser(request *restful.Request, response *restful.Response) {
	user := models.User{}
	body, _ := ioutil.ReadAll(request.Request.Body)
	err := json.Unmarshal(body, &user)
	if user.UnionID == "" || user.OpenID == ""{
		utils.RespondStruct(response, 500, user, fmt.Errorf("用户OpenID和UnionID不存在！"))
		return
	}
	err = models.UserAdd(&user)
	utils.RespondStruct(response, 200, map[string]interface{}{"user": user}, err)

}

func EditUser(request *restful.Request, response *restful.Response) {
	user := models.User{}
	body, _ := ioutil.ReadAll(request.Request.Body)
	err := json.Unmarshal(body, &user)
	if user.UnionID == "" || user.OpenID == ""{
		utils.RespondStruct(response, 500, user, fmt.Errorf("用户OpenID和UnionID不存在！"))
		return
	}
	err = models.UserAdd(&user)
	utils.RespondStruct(response, 200, map[string]interface{}{"user": user}, err)

}

func GetUserDetail(request *restful.Request, response *restful.Response) {
	idString := request.PathParameter("userid")
	userID, _ := strconv.ParseInt(idString, 10, 64)
	user := models.User{UserID: userID}
	u, err := models.UserSearch(&user)
	ret := map[string]interface{}{"userInfo": u}
	utils.RespondStruct(response, 200, ret, err)
}

func GetUserList(request *restful.Request, response *restful.Response) {
	users, err := models.UserList()
	ret := map[string]interface{}{"list": users}
	utils.RespondStruct(response, 200, ret, err)
}