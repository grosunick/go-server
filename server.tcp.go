package server

import (
	"net"
	"time"
	"github.com/grosunick/go-common/worker"
	"github.com/grosunick/go-common/queue"
)

const (
	TIME_SLEEP_AFTER_ACCEPT_ERROR = 100
)

// TCP socket server struct
type SocketServerTcp struct {
	socketServer
}

// TCP socket server factory
func NewSocketServerTcp(config *SocketServerConfig, router IRouter) SocketServerTcp {
	return SocketServerTcp{
		socketServer{
			config,
			router,
			worker.NewTreadPoolWorker(
				config.MaxWorkerAmount,
				queue.NewChannelBasedLimitQueue(config.MaxUnhandledRequests),
			),
		},
	}
}

// Realize port listening.
func (this *SocketServerTcp) Listen() (err error) {
	var listener net.Listener

	// start listerning port
	listener, err = net.Listen("tcp", this.config.Addr)
	defer listener.Close()

	if err != nil {
		return
	}

	// run thread pull
	go this.Run()

	for {
		// new connection handling
		conn, err := listener.Accept()
		if err != nil {
			time.Sleep(time.Millisecond * TIME_SLEEP_AFTER_ACCEPT_ERROR)
			continue
		}

		// add task to the processing
		this.Add(
			&workerHandler{
				this.router,
				NewTcpConnection(conn, this.config),
			},
		)
	}
}
