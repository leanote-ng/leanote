package tests

import (
	"testing"

	"github.com/leanote-ng/leanote/app/db"
	"github.com/revel/revel"

	//	. "github.com/leanote-ng/leanote/app/lea"
	"github.com/leanote-ng/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func init() {
	revel.Init("dev", "github.com/leanote/leanote", "/Users/life/Documents/Go/package_base/src")
	db.Init("mongodb://localhost:27017/leanote", "leanote")
	service.InitService()
	service.ConfigS.InitGlobalConfigs()
}

// 测试登录
func TestSendMail(t *testing.T) {
	ok, err := service.EmailS.SendEmail("life@leanote.com", "你好", "你好吗")
	t.Log(ok)
	t.Log(err)
}
