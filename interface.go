package server

// Интерфейс обработчика запроса
type IHandler interface {
	Handle(params string, connection IConnection)
}

// Интерфейс маршрутизатора запросов
type IRouter interface {
	// Добавляет обработчик для метода method
	AddRoute(method string, handler IHandler)
	// Возвращает обработчик метода method
	GetHandler(method string) (IHandler, bool)
}

// Интерфейс сервера
type IServer interface {
	Listen() error      // Прослушивает порт
}

// интерфейс подключения
type IConnection interface {
	Read() (string, error)  	// Читает данные текстовые
	Write(message string) error // Записывает текстовые данные
	Close() error               // Закрывает соединение
}
