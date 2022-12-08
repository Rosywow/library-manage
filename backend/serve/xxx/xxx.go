package xxx

//author: {"name":"xxx","email":"XUnion@GMail.com"}
//annotation:xxx-service

import (
	"backend/cmn"
	"context"
	"encoding/json"
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
		Fn: xxx,

		Path: "/api/xxx",
		Name: "xxx",

		Developer: developer,
	})
}


func xxx(ctx context.Context) {
	q := cmn.GetCtxValue(ctx)
	q.W.Write([]byte(`{"data":"I'm from xxx"}`))
}

