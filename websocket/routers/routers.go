package routers

import (
	"html/template"
	"net/http"

	"github.com/woodylan/go-websocket/api/bind2group"
	"github.com/woodylan/go-websocket/api/closeclient"
	"github.com/woodylan/go-websocket/api/getonlinelist"
	"github.com/woodylan/go-websocket/api/register"
	"github.com/woodylan/go-websocket/api/send2client"
	"github.com/woodylan/go-websocket/api/send2clients"
	"github.com/woodylan/go-websocket/api/send2group"
	"github.com/woodylan/go-websocket/servers"
)

func Init() {
	registerHandler := &register.Controller{}
	sendToClientHandler := &send2client.Controller{}
	sendToClientsHandler := &send2clients.Controller{}
	sendToGroupHandler := &send2group.Controller{}
	bindToGroupHandler := &bind2group.Controller{}
	getGroupListHandler := &getonlinelist.Controller{}
	closeClientHandler := &closeclient.Controller{}

	http.HandleFunc("/abc", home)
	http.HandleFunc("/api/register", registerHandler.Run)
	http.HandleFunc("/api/send_to_client", AccessTokenMiddleware(sendToClientHandler.Run))
	http.HandleFunc("/api/send_to_clients", AccessTokenMiddleware(sendToClientsHandler.Run))
	http.HandleFunc("/api/send_to_group", AccessTokenMiddleware(sendToGroupHandler.Run))
	http.HandleFunc("/api/bind_to_group", AccessTokenMiddleware(bindToGroupHandler.Run))
	http.HandleFunc("/api/get_online_list", AccessTokenMiddleware(getGroupListHandler.Run))
	http.HandleFunc("/api/close_client", AccessTokenMiddleware(closeClientHandler.Run))

	servers.StartWebSocket()

	go servers.WriteMessage()
}

func home(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	get_act := query["systemId"][0]

	homeTemplate.Execute(w, get_act)
}

var homeTemplate = template.Must(template.New("").Parse(`
<!doctype html>
<html>
  	<head>
    	<title>第三方支付中间件</title>
		<meta http-equiv="content-type" content="text/html; charset=utf-8">
	</head> 	
<body>
<table>

<h2>WebSocket Test</h2> 
<div id="output"></div> 



<form action="" id="loginForm" class="fh5co-form animate-box" onsubmit="return false" method="post"
                  data-animate-effect="fadeInLeft">
                <h2>Register</h2>
                <div class="form-group">
                    <label for="systemId" class="sr-only">systemId</label>
                    <input type="text" class="form-control" id="systemId" name="systemId" placeholder="systemId"
                           autocomplete="off">
                </div>
                <div class="form-group">
                    <input type="button" onclick="register()" value="register" class="btn btn-primary">
                </div>
</form>
</table>

<input type="text" class="form-control" id="Connect" name="Connect" placeholder="Connect"
                           autocomplete="off">
<button id="open">Connect to websocket</button>



<table>
<form action="" id="loginForm" class="fh5co-form animate-box" onsubmit="return false" method="post"
                  data-animate-effect="fadeInLeft">
                <h2>Send to Client</h2>
                <div class="form-group">
                    <label class="sr-only">send_to_client</label>
                    <input type="text" class="form-control" id="message" name="message" placeholder="message"
                           autocomplete="off">
                    <input type="text" class="form-control" id="clientId" name="clientId" placeholder="clientId"
                           autocomplete="off">
                </div>
                <div class="form-group">
                    <input type="button" onclick="send_to_client()" value="send_to_client" class="btn btn-primary">
                </div>
</form>
</table>



<table>
<form action="" id="loginForm" class="fh5co-form animate-box" onsubmit="return false" method="post"
                  data-animate-effect="fadeInLeft">
                <h2>Send to Server</h2>
                <div class="form-group">
                    <label class="sr-only">Send to Server</label>
                    <input type="text" class="form-control" id="message_server" name="message_server" placeholder="message_server"
                           autocomplete="off">
                </div>
                <div class="form-group">
                    <input type="button" onclick="client_server()" value="send_to_server" class="btn btn-primary">
                </div>
</form>
</table>


<table>
<form action="" id="loginForm" class="fh5co-form animate-box" onsubmit="return false" method="post"
                  data-animate-effect="fadeInLeft">
                <h2>bind_to_group</h2>
                <div class="form-group">
                    <label class="sr-only">bind_to_group</label>
                    <input type="text" class="form-control" id="groupName" name="groupName" placeholder="groupName"
                         autocomplete="off">
                  <input type="text" class="form-control" id="userId" name="userId" placeholder="userId"
                         autocomplete="off">
                           
                </div>
                <div class="form-group">
                    <input type="button" onclick="bind_to_group()" value="bind_to_group" class="btn btn-primary">
                </div>
</form>
</table>


<table>
<form action="" id="loginForm" class="fh5co-form animate-box" onsubmit="return false" method="post"
                  data-animate-effect="fadeInLeft">
                <h2>get_online_list</h2>
                <div class="form-group">
                    <input type="button" onclick="get_online_list()" value="get_online_list" class="btn btn-primary">
                </div>
</form>
</table>


<table>
<form action="" id="loginForm" class="fh5co-form animate-box" onsubmit="return false" method="post"
                  data-animate-effect="fadeInLeft">
                <h2>send_to_group</h2>
                <div class="form-group">
                    <input type="button" onclick="send_to_group()" value="send_to_group" class="btn btn-primary">
                </div>
</form>
</table>





</body>



<script type="text/javascript">
function register() {

var httpRequest = new XMLHttpRequest();
httpRequest.open('POST', '/api/register', true); 
httpRequest.setRequestHeader("Content-type","application/json");
httpRequest.send(JSON.stringify({'systemId':{{.}}}));
httpRequest.onreadystatechange = function (){
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {//验证请求是否发送成功
        var json = httpRequest.responseText;//获取到服务端返回的数据
		console.log(json);
    }
};
 }

 function send_to_client() {
var systemId = document.getElementById("systemId").value;
var clientId = document.getElementById("clientId").value;
var message = document.getElementById("message").value;
var systemId = {{.}}
var httpRequest = new XMLHttpRequest();
httpRequest.open('POST', '/api/send_to_client', true); 
httpRequest.setRequestHeader("Content-type","application/json");
httpRequest.setRequestHeader("SystemId",systemId);
httpRequest.send(JSON.stringify({'systemId':systemId,'clientId':clientId,'sendUserId':'userid','code':200,'msg':message,'data':'data'}));
httpRequest.onreadystatechange = function (){
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {//验证请求是否发送成功
        var json = httpRequest.responseText;//获取到服务端返回的数据
		console.log(json);
		alert(json);
    }
};
 }


function bind_to_group() {
var clientId = document.getElementById("clientId").value;
var userid = document.getElementById("userId").value;
var groupName = document.getElementById("groupName").value;
var systemId = {{.}}
var httpRequest = new XMLHttpRequest();
httpRequest.open('POST', '/api/bind_to_group', true); 
httpRequest.setRequestHeader("Content-type","application/json");
httpRequest.setRequestHeader("SystemId",systemId);
httpRequest.send(JSON.stringify({'systemId':systemId,'clientId':clientId,'userId':userid,'groupName':groupName,'extend':'extend'}));
httpRequest.onreadystatechange = function (){
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {//验证请求是否发送成功
        var json = httpRequest.responseText;//获取到服务端返回的数据
		console.log(json);
		alert(json);
    }
};
 }

 function get_online_list() {
var groupName = document.getElementById("groupName").value;
var systemId = {{.}}
var httpRequest = new XMLHttpRequest();
httpRequest.open('POST', '/api/get_online_list', true); 
httpRequest.setRequestHeader("Content-type","application/json");
httpRequest.setRequestHeader("SystemId",systemId);
httpRequest.send(JSON.stringify({'groupName':groupName}));
httpRequest.onreadystatechange = function (){
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {//验证请求是否发送成功
        var json = httpRequest.responseText;//获取到服务端返回的数据
		console.log(json);
		alert(json);
    }
};
 }



 function send_to_group() {
var groupName = document.getElementById("groupName").value;
var systemId = {{.}}
var message = document.getElementById("message").value;
var userid = document.getElementById("userId").value;
var httpRequest = new XMLHttpRequest();
httpRequest.open('POST', '/api/send_to_group', true); 
httpRequest.setRequestHeader("Content-type","application/json");
httpRequest.setRequestHeader("SystemId",systemId);
httpRequest.send(JSON.stringify({'msg':message,'sendUserId':userid,'groupName':groupName,'code':200,'data':'data'}));
httpRequest.onreadystatechange = function (){
    if (httpRequest.readyState == 4 && httpRequest.status == 200) {//验证请求是否发送成功
        var json = httpRequest.responseText;//获取到服务端返回的数据
		console.log(json);
		alert(json);
    }
};
 }

  window.addEventListener("load", function(evt) {
 
        register()
        var systemId = {{.}}
        websocket = new WebSocket("ws://127.0.0.1:8081/ws?systemId="+systemId);

        websocket.onopen = function(evt) {
            onOpen(evt)
        };
        websocket.onclose = function(evt) {
            onClose(evt)
        };
        websocket.onmessage = function(evt) {
            onMessage(evt)
        };
        websocket.onerror = function(evt) {
            onError(evt)
        };

});

function onOpen(evt) {
        writeToScreen("CONNECTED");
        doSend("WebSocket rocks");
    } 
  
    function onClose(evt) {
        writeToScreen("DISCONNECTED");
    } 
  
    function onMessage(evt) {
        writeToScreen('<span style="color: LightSalmon;">RESPONSE: '+ evt.data+'</span>');
        //websocket.close();
    } 
  
    function onError(evt) {
        writeToScreen('<span style="color: red;">ERROR:</span> '+ evt.data);
    } 
  
    function doSend(message) {
        writeToScreen("SENT: " + message); 
        websocket.send(message);
    } 

    function client_server(){
        var message = document.getElementById("message_server").value;
        writeToScreen("SENT: " + message); 
        websocket.send(message);
    }
  
    function writeToScreen(message) {
        var pre = document.createElement("p");
        pre.style.wordWrap = "break-word";
        pre.innerHTML = message;
        output.appendChild(pre);
    } 
    </script>
</html>
`))
