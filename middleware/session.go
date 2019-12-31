package middleware

import (
	"context"
	"github.com/SongLiangChen/common/session"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"grpc-gateway/config"
	"net/http"
)

const (
	KeyUserId       = "userId"
	KeyNickName     = "nickName"
	KeyLoginAccount = "loginAccount"
)

func getUserId(store session.Store) string {
	userId := store.Get(KeyUserId)
	if userId == nil {
		return ""
	}
	return userId.(string)
}

func InitSession() gin.HandlerFunc {
	cnf := config.GetConfig()

	manager, err := session.NewManager(cnf.SessionCnf.ProviderType, &session.ManagerConfig{
		CookieName:      cnf.SessionCnf.CookieName,
		EnableSetCookie: cnf.SessionCnf.EnableSetCookie,
		Maxlifetime:     cnf.SessionCnf.MaxLifeTime,
		Secure:          cnf.SessionCnf.Secure,
		ProviderConfig:  cnf.SessionCnf.ProviderConfig,
		CookieLifeTime:  cnf.SessionCnf.CookieLifeTime,
	})
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		store, err := manager.SessionStart(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}

		sid := store.SessionID()
		userId := getUserId(store)

		// TODO 校验单点登录

	}
}

func SessionResponseOption(context.Context, http.ResponseWriter, proto.Message) error {
	return nil
}
