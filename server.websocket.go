package server

import (
	"github.com/go-errors/errors"
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

// Структура вебсокет сервера
type WebSocketServer struct {
	config *WebSocketServerConfig // настройки сервера сервера
	server *http.Server // http сервер
	router IRouter // роутер json комманд
	upgrader websocket.Upgrader
}

// Создает объект websocket сервера
func NewWebSocketServer(server *http.Server, config *WebSocketServerConfig, router IRouter) *WebSocketServer {
	server.Addr = config.Addr

	upgrader := websocket.Upgrader{
		ReadBufferSize:  config.ReadBufferSize,
		WriteBufferSize: config.WriteBufferSize,
	}

	// отменяем проверку заголовка origin
	if config.CheckOrigin == false {
		upgrader.CheckOrigin = func(r *http.Request) bool {return true}
	}

	return &WebSocketServer{
		config: config,
		server: server,
		router: router,
		upgrader: upgrader,
	}
}

// Запускает сервер
func (this *WebSocketServer) Listen() {
	// устанавливаем обработчик подключения по протоколу websocket
	http.HandleFunc(this.config.Url, this.handle)

	err := this.server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Метод обработчик клиентских соединений
func (this *WebSocketServer) handle(w http.ResponseWriter, r *http.Request) {
	// если обработчик запаниковал, залогируем это
	defer func() {
		if err := recover(); err != nil {
			log.Println(errors.Wrap(err, 3).ErrorStack())
		}
	}()

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	// осушествляем переход на протокол websocket
	ws, err := this.upgrader.Upgrade(w, r, nil)
	if err != nil {
		if this.config.ShowLog {
			log.Println(err)
		}
		return
	}

	// создаем обработчик соединения и запускаем его
	handler := &WebSocketHandler{
		&workerHandler{
			this.router,
			NewWebSocketConnection(ws, this.config),
		},
		this.config.ShowLog,
	}

	handler.Run()
}