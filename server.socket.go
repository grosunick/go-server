package server

import (
	"github.com/grosunick/go-common/worker"
)

// Базовая структура сокет сервера
type socketServer struct {
	// Обхект настроек сервера
	config *SocketServerConfig
	// Роутер комманд. Определяет обработчики, соответсвующие определенной комманде
	router IRouter
	// Пул потоков. Каждый поток будет занималься обработкой конкретного соединения
	worker.IMultyThreadWorker
}
