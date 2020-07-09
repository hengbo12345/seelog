package seelog

import (
	"errors"
	"fmt"
	"github.com/xmge/seelog/page"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	PageIndex = "index.html"
	Page403   = "403.html"
)

// 开启 httpServer
func server(port int, password string) {

	defer func() {
		if err := recover(); err != nil {
			printError(errors.New("server panic"))
		}
	}()

	// socket链接
	http.Handle("/ws", websocket.Handler(genConn))

	// 访问页面
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if !(strings.Replace(request.RequestURI, "/", "", -1) == password) {
			showPage(writer, Page403, nil)
			return
		}
		showPage(writer, PageIndex, slogs)
	})
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// 输出page
func showPage(writer http.ResponseWriter, pageName string, data interface{}) {
	t, err := template.New(pageName).Parse(page.GetPage(pageName))
	if err != nil {
		printError(err)
	}
	t.Execute(writer, data)
}

// 创建client对象
func genConn(ws *websocket.Conn) {
	client := &client{time.Now().String(), ws, make(chan msg, 1024), slogs[0].Name}
	manager.register <- client
	go client.read()
	client.write()
}
