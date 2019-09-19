package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var httpServerInstant *http.Server

func init() {
	fmt.Println("http server init")
	http.HandleFunc("/web/login", LoginHandler)
}

func Start() {
	if nil == httpServerInstant {
		httpServerInstant = &http.Server{Addr: "127.0.0.1:8080"}
		if err := httpServerInstant.ListenAndServe(); nil != err && err != http.ErrServerClosed {
			log.Println(err.Error())
		}
	}
}

func Stop() {
	if nil != httpServerInstant {
		if err := httpServerInstant.Close(); nil != err {
			log.Println(err)
		}
		httpServerInstant = nil
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// 这些信息是输出到服务器端的打印信息
	fmt.Println("Request解析")
	// HTTP方法
	fmt.Println("method", r.Method)
	// RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
	fmt.Println("RequestURI", r.RequestURI)
	// URL类型,下方分别列出URL的各成员
	fmt.Println("URL_scheme", r.URL.Scheme)
	fmt.Println("URL_opaque", r.URL.Opaque)
	fmt.Println("URL_user", r.URL.User.String())
	fmt.Println("URL_host", r.URL.Host)
	fmt.Println("URL_path", r.URL.Path)
	fmt.Println("URL_RawQuery", r.URL.RawQuery)
	fmt.Println("URL_Fragment", r.URL.Fragment)
	// 协议版本
	fmt.Println("proto", r.Proto)
	fmt.Println("protomajor", r.ProtoMajor)
	fmt.Println("protominor", r.ProtoMinor)

	body, _ := ioutil.ReadAll(r.Body)
	values, err := url.ParseQuery(string(body))

	if ret, err := w.Write([]byte("SUCCESS")); nil != err {
		log.Println("login occur error" + err.Error() + string(ret))
	}

	log.Println("login response success")
}
