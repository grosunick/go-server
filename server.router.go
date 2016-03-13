package server

type Router map[string]IHandler

// Создает объект роутера
func NewServerRouter() Router {
	return Router(make(map[string]IHandler))
}

// Добавляет обработчик для метода method
func (this Router) AddRoute(method string, handler IHandler) {
	this[method] = handler
}

// Возвращает обработчик метода method
func (this Router) GetHandler(method string) (handler IHandler, err bool) {
	handler, err = this[method]
	return
}
