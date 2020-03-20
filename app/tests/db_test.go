package tests

import (
	"testing"

	"github.com/leanote-ng/leanote/app/db"
	//	. "github.com/leanote-ng/leanote/app/lea"
	//	"github.com/leanote-ng/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}
