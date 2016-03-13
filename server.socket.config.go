package server

// Структура, определяющая настройки socket сервера
type SocketServerConfig struct {
	Addr string 				// интерфейс, слушающий входящие запросы
	MaxWorkerAmount int32 		// максимальное количество обработчиков
	MaxUnhandledRequests uint32 // максимальный размер очереди необработанных запросов
	ReadTimeOut int16 			// socket read timeout (миллисекунды)
	WriteTimeOut int16 			// socket write timeout (миллисекунды)
}
