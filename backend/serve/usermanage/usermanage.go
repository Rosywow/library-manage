package usermanage
//author: {"name":"auth","email":"XUnion@GMail.com"}
//annotation:user-service

import (
	"backend/cmn"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func Enroll(author string) {
	var developer *cmn.ModuleAuthor
	if author != "" {
		var d cmn.ModuleAuthor
		err := json.Unmarshal([]byte(author), &d)
		if err != nil {
			log.Println(err.Error())
			return
		}
		developer = &d
	}

	cmn.AddService(&cmn.ServeEndPoint{
		Fn: getuser,

		Path: "/api/getuser",
		Name: "getuser",

		Developer: developer,
	})
}


func getuser(ctx context.Context) {
	fmt.Println("getFileList")
	//q := cmn.GetCtxValue(ctx)
}