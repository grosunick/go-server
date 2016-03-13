package server

import (
	"encoding/json"
	"log"
)

// Обработчик подключения websocket сервера
type WebSocketHandler struct {
	*workerHandler
	showLog bool
}

// Читает json комманду, поступившую от клиента
func (this *WebSocketHandler) readCommand() (command string, err error) {
	// пытаемся прочитать комманду клиента
	command, err = this.connection.Read()
	if err != nil {
		if this.showLog {
			log.Println("Read error: " + err.Error() + "\n")
		}

		this.connection.Close()
	}

	return
}

// Получает описание метода
func (this *WebSocketHandler) getMethodDesc(message string) (request ServerRequest, err error) {
	err = nil

	if err = json.Unmarshal([]byte(message), &request); err != nil {
		if this.showLog {
			log.Println("Invalid json \n")
			log.Println(err.Error())
		}

		this.connection.Close()
	}

	return
}

// Возвращает обработчик метода
func (this *WebSocketHandler) getMethodHandler(method string) (handler IHandler, res bool) {
	handler, res = this.router.GetHandler(method)
	if !res {
		if this.showLog {
			log.Println("undefined method: " + method + " \n")
		}

		this.connection.Close()
	}

	return
}
