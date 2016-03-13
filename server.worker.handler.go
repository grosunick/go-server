package server

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"log"
)

type workerHandler struct {
	router     IRouter
	connection IConnection
}

// Обработчик клиентского соединения.
// -- Читает поступившую от клиента комманду
// -- Парсит метод комманды и ее параметры
// -- Выполняет комманду и возвращает ответ
func (this *workerHandler) Run() (result interface{}, err error) {
	// если обработчик запаниковал, залогируем это
	defer func() {
		if err := recover(); err != nil {
			this.connection.Close()
			log.Println(errors.Wrap(err, 2).ErrorStack())
		}
	}()

	err = nil

	for {
		var command string
		var request ServerRequest

		// читаем комманду, поступившую от клиента
		command, err = this.readCommand()
		if err != nil {
			return
		}

		// Получаем описание метода
		request, err = this.getMethodDesc(command)
		if err != nil {
			return
		}

		// Пробуем найти обработчик метода
		handler, ok := this.getMethodHandler(request.Method)
		if !ok {
			return
		}

		// запуск обработчика метода
		handler.Handle(string(request.Params), this.connection)
	}
}

// Читает json комманду, поступившую от клиента
func (this *workerHandler) readCommand() (command string, err error) {
	// пытаемся прочитать комманду клиента
	command, err = this.connection.Read()
	if err != nil {
		this.connection.Write("Read error: " + err.Error() + "\n")
		this.connection.Close()
	}

	return
}

// Получает описание метода
func (this *workerHandler) getMethodDesc(message string) (request ServerRequest, err error) {
	err = nil

	if err = json.Unmarshal([]byte(message), &request); err != nil {
		this.connection.Write("Invalid json \n")
		this.connection.Write(err.Error())
		this.connection.Close()
	}

	return
}

// Возвращает обработчик метода
func (this *workerHandler) getMethodHandler(method string) (handler IHandler, res bool) {
	handler, res = this.router.GetHandler(method)
	if !res {
		this.connection.Write("undefined method: " + method + " \n")
		this.connection.Close()
	}

	return
}
