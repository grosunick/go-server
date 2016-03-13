package server

import (
	"github.com/go-errors/errors"
	"github.com/gorilla/websocket"
	"time"
)

// Объект подключения по протоколу вебсокет
type WebSocketConnection struct {
	conn *websocket.Conn
	config *WebSocketServerConfig
	onClose map[int]func()
}

// Создание объекта подключения по простоколу websocket
func NewWebSocketConnection(conn *websocket.Conn, config *WebSocketServerConfig) *WebSocketConnection {
	return &WebSocketConnection{conn, config, make(map[int]func())}
}

// Добавляет обрбаотчик, отрабатывающий перед закрытием подключения
func (this *WebSocketConnection) AddOnCloseFunction(fn func()) {
	size := len(this.onClose) - 1
	if size < 0 {
		size = 0
	}

	this.onClose[size] = fn
}

// Закрывает соединение
func (this *WebSocketConnection) Close() error {
	// Перед закрытием выполняем все обработчики
	for _, fn := range this.onClose {
		fn()
	}

	return this.conn.Close()
}

// Читает текстовое сообщение из сокета. Если пришли бинарные данные, вернет ошибку
func (this *WebSocketConnection) Read() (message string, err error) {
	if this.config.ReadTimeOut > 0 {
		this.conn.SetReadDeadline(
			time.Now().Add(
				time.Millisecond * time.Duration(this.config.ReadTimeOut),
			),
		)
	}

	messageType, data, err := this.conn.ReadMessage()
	message = string(data)

	if messageType == websocket.BinaryMessage {
		err = errors.New("not text message")
		return
	}

	return message, err
}

// Записывает данные в сокет
func (this *WebSocketConnection) Write(message string) error {
	if this.config.WriteTimeOut > 0 {
		this.conn.SetWriteDeadline(
			time.Now().Add(
				time.Millisecond * time.Duration(this.config.WriteTimeOut),
			),
		)
	}

	return this.conn.WriteMessage(websocket.TextMessage, []byte(message))
}