package server

// Структура, определяющая настройки websocket сервера
type WebSocketServerConfig struct {
	Addr string // интерфейс, слушающий входящие запросы
	Url string // адрес websocket сщединения
	ReadTimeOut int16 // timeout чтения из websocket соединения (миллисекунды)
	WriteTimeOut int16 // timeout записи в websocket соединение (миллисекунды)
	ReadBufferSize int // буфер чтения из websocket
	WriteBufferSize int // буфер записи в websocket
	CheckOrigin bool // Проверять ли заголовок origin
	ShowLog bool // признак вывода отладочной информации
}
