package service

//go:generate go run service-enroll-generate.go -a=annotation:(?P<name>.*)-service

import (
	"context"
	"backend/cmn"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"net/http"

	"time"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"


	"regexp"

	"sort"

)

var (
	z *zap.Logger

	pgxConn *pgxpool.Pool
	sqlxDB  *sqlx.DB
	rConn   redis.Conn
)



var store = sessions.NewCookieStore([]byte("aLongStory"))







var rIsAPI = regexp.MustCompile(`(?i)^/api/(.*)?$`)

var (
	rWxIOS     = regexp.MustCompile(`(iPhone)(.*)(MicroMessenger)`)
	rWxAndroid = regexp.MustCompile(`(Android)(.*)(MicroMessenger)`)
	rMacWx     = regexp.MustCompile(`\(Macintosh; .*(?P<osVer> \d*_\d*_\d*\)).* MicroMessenger/(?P<wxVer>\d*\.\d*\.\d*)\((?P<wxVerHex>.*)\) MacWechat`)
	rWinWx     = regexp.MustCompile(`\(Windows \S* (?P<osVer>\d*\.\d*)(; )?(?P<subSys>\S*)\).* MicroMessenger/(?P<wxVer>\d*\.\d*\.\d*)`)
	rIsWx      = regexp.MustCompile(`MicroMessenger`)
	rMobile    = regexp.MustCompile(`(Android|iPhone)`)
)

func reqProc(reqPath string, w http.ResponseWriter, r *http.Request) {
	//以单例运行

	// ---------------------------
	q := &cmn.ServiceCtx{

		R: r,
		W: w,

		Redis: rConn,

		Ep: cmn.Services[reqPath],

		ReqAdminFnc: r.URL.Query().Get("admin") == "true",

		BeginTime: time.Now(),
	}

	ctx := context.WithValue(context.Background(), cmn.QNearKey, q)


	cmn.Services[reqPath].Fn(ctx)
}

func WebServe() {
	Enroll()

	router := mux.NewRouter()

	var rootExists bool
	var pathList []string
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

		fmt.Println(k)
		//添加路由
		router.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			reqProc(k, w, r)
		})
	}

	//host := "qnear.cn"
	//if viper.IsSet("webServe.serverName") {
	//	host = viper.GetString("webServe.serverName")
	//}
	//
	//appLaunchPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	z.Fatal(err.Error())
	//	return
	//}
	//
	//certPath := appLaunchPath + "/certs"
	//var hostWhiteList string
	//if viper.IsSet("webServe.hostWhiteList") {
	//	hostWhiteList = viper.GetString("webServe.hostWhiteList")
	//	names := strings.Split(hostWhiteList, ",")
	//	host := "qnear.cn"
	//	if viper.IsSet("webServe.serverName") {
	//		host = viper.GetString("webServe.serverName")
	//	}
	//	var exists bool
	//	for _, v := range names {
	//		if v == host {
	//			exists = true
	//			break
	//		}
	//	}
	//	if !exists {
	//		log.Fatal(fmt.Sprintf("webServe.serverName:%s must exists in webServe.hostWhiteList: %s",
	//			host, hostWhiteList))
	//	}
	//}
	//
	//if hostWhiteList == "" {
	//	hostWhiteList = host
	//}
	//
	//certManager := autocert.Manager{
	//	Prompt: autocert.AcceptTOS,
	//
	//	HostPolicy: autocert.HostWhitelist(
	//		strings.Split(hostWhiteList, ",")...), //Your domain here
	//
	//	Cache: autocert.DirCache(certPath), //Folder for storing certificates
	//}
	//
	////getWxAccessToken(2)
	//
	//httpListenPort := 8080
	//if viper.IsSet("webServe.httpListenPort") {
	//	httpListenPort = viper.GetInt("webServe.httpListenPort")
	//}
	//
	//httpsListenPort := 8443
	//if viper.IsSet("webServe.httpsListenPort") {
	//	httpsListenPort = viper.GetInt("webServe.httpsListenPort")
	//}
	//
	//var autoCert bool
	//if viper.IsSet("webServe.autoCert") {
	//	autoCert = viper.GetBool("webServe.autoCert")
	//}
	//
	//var ep string
	//if autoCert {
	//	ep = fmt.Sprintf(":%v", httpsListenPort)
	//} else {
	//	ep = fmt.Sprintf(":%v", httpListenPort)
	//}
	//
	//serv := &http.Server{
	//	Addr:    ep,
	//	Handler: GzipHandler(router),
	//}
	//
	//if autoCert {
	//	serv.TLSConfig = &tls.Config{GetCertificate: certManager.GetCertificate}
	//	go func() { _ = http.ListenAndServe(":http", certManager.HTTPHandler(nil)) }()
	//	_ = serv.ListenAndServeTLS("", "")
	//	return
	//}
	//
	//cmn.AppStartTime = time.Now()
	//
	//z.Info(cmn.AppStartTime.Format(cmn.AppStartTimeLayout))
	//_ = serv.ListenAndServe()
}
