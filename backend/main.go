package main

import (
	"backend/cmn"
	"backend/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sort"
	"time"
)

func reqProc(reqPath string, w http.ResponseWriter, r *http.Request) {
	//以单例运行

	// ---------------------------
	q := &cmn.ServiceCtx{
		R: r,
		W: w,

		Ep: cmn.Services[reqPath],

		BeginTime: time.Now(),
	}

	ctx := context.WithValue(context.Background(), cmn.QNearKey, q)

	fmt.Println("reqPath:",reqPath)
	cmn.Services[reqPath].Fn(ctx)
}

func WebServe(){
	router := mux.NewRouter()
	service.Enroll()
	var rootExists bool
	var pathList []string
	fmt.Println("services:",cmn.Services)
	for k := range cmn.Services {
		if k == "/" {
			rootExists = true
			continue
		}
		pathList = append(pathList, k)
	}
	sort.Strings(pathList)
	if rootExists {
		pathList = append(pathList, "/")
	}

	for _, k := range pathList {
		k:=k
		router.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			reqProc(k, w, r)
		})
	}

	http.ListenAndServe(":1244",router)
}

func main() {
	fmt.Println("start")
	log.SetFlags(log.Ldate |log.Lshortfile )
	WebServe()
}