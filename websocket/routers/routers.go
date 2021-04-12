package routers

import (
	"html/template"
	"net/http"

	"websocket/api/bind2group"
	"websocket/api/closeclient"
	"websocket/api/getonlinelist"
	"websocket/api/register"
	"websocket/api/send2client"
	"websocket/api/send2clients"
	"websocket/api/send2group"
	"websocket/api/sendtoclient"
	"websocket/servers"
)

func Init() {
	registerHandler := &register.Controller{}
	sendToClientHandler := &send2client.Controller{}
	sendToClientsHandler := &send2clients.Controller{}
	sendToGroupHandler := &send2group.Controller{}
	bindToGroupHandler := &bind2group.Controller{}
	getGroupListHandler := &getonlinelist.Controller{}
	closeClientHandler := &closeclient.Controller{}
	websendtoclientHandler := &sendtoclient.Controller{}

	http.HandleFunc("/abc", home)
	http.HandleFunc("/api/register", registerHandler.Run)
	http.HandleFunc("/api/send_to_client", AccessTokenMiddleware(sendToClientHandler.Run))
	http.HandleFunc("/api/send_to_clients", AccessTokenMiddleware(sendToClientsHandler.Run))
	http.HandleFunc("/api/send_to_group", AccessTokenMiddleware(sendToGroupHandler.Run))
	http.HandleFunc("/api/bind_to_group", AccessTokenMiddleware(bindToGroupHandler.Run))
	http.HandleFunc("/api/get_online_list", AccessTokenMiddleware(getGroupListHandler.Run))
	http.HandleFunc("/api/close_client", AccessTokenMiddleware(closeClientHandler.Run))
	http.HandleFunc("/send_message", WebTokenMiddleware(websendtoclientHandler.Run))

	servers.StartWebSocket()

	//启动线程，分类信息
	servers.MessageChannel = servers.NewChannel_Pool(5)
	//开启监视任务
	servers.MessageChannel.NewChannel_PoolGo()
	//处理任务
	servers.MessageGetResult()

	go servers.WriteMessage()
}

func home(w http.ResponseWriter, r *http.Request) {

	homeTemplate.Execute(w, "123")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!doctype html>
<html>
  	<head>
    	<title>Hello World</title>
		<meta http-equiv="content-type" content="text/html; charset=utf-8">
	</head> 	
<body>
Hello world !!!
</body>
</html>
`))
