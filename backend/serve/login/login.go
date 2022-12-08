package login

//author: {"name":"auth","email":"XUnion@GMail.com"}
//annotation:login-service

import (
	"backend/cmn"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4"
	"io/ioutil"

	"log"
	"os"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

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
		Fn: login,

		Path: "/api/login",
		Name: "login",

		Developer: developer,
	})

	cmn.AddService(&cmn.ServeEndPoint{
		Fn: signup,

		Path: "/api/signup",
		Name: "signup",

		Developer: developer,
	})
}

func login(ctx context.Context) {
	fmt.Println("登录")
	q := cmn.GetCtxValue(ctx)
	b, err := ioutil.ReadAll(q.R.Body)
	if err != nil {
		_, _ = q.W.Write([]byte(fmt.Sprintf(`{"status":0,"err":"%s"}`,err)))
		return
	}

	body := make(map[string]interface{})
	err = json.Unmarshal(b, &body)
	if err != nil {
		_, _ = q.W.Write([]byte(fmt.Sprintf(`{"status":0,"err":"%s"}`,err)))
		return
	}

	username := body["username"]
	password := body["password"]
	if username == "" || password == "" {
		_, _ = q.W.Write([]byte(`{"status":0,"err":"用户名或账号为空"}`))
		return
	}

	var uid interface{}
	s := `select uid from user_login where username = $1 and password = $2`


	err = cmn.Connection.QueryRow(context.Background(), s, username, password).Scan(uid)
	if err == pgx.ErrNoRows {
		_, _ = q.W.Write([]byte(`{"status":0,"err":"密码错误"}`))
	} else {
		_, _ = q.W.Write([]byte(`{"status":200}`))
	}

	_, _ = q.W.Write([]byte(`{"status":200}`))
}

func signup(ctx context.Context) {
	fmt.Println("注册")
	//q := cmn.GetCtxValue(ctx)

	// if exit

	// if not exit
}
