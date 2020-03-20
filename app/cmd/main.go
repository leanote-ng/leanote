package main

import (
	"fmt"

	"github.com/revel/cmd/harness"
	"github.com/revel/cmd/model"
	"github.com/revel/revel"
)

func main() {
	// go run main.go
	// 生成routes.go, main.go
	Cmd := model.CommandConfig{}
	Cmd.InitGoPaths()
	Cmd.InitPackageResolver()

	var callback model.RevelCallback = model.NewWrappedRevelCallback(nil, Cmd.PackageResolver)
	RevelContainers, err := model.NewRevelPaths("dev", "github.com/leanote/leanote", "", callback)
	if err != nil {
		return
	}

	revel.Init("", "github.com/leanote/leanote", "")
	_, err = harness.Build(&Cmd, RevelContainers) // ok, err
	if err != nil {
		panic(err)
	}
	fmt.Println("Ok")
	//	panicOnError(reverr, "Failed to build")
}
