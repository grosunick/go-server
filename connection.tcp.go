package server

import (
	"bufio"
	"net"
	"time"
)

// Объект подключения по протоколу tcp
type TcpConnection struct {
	conn   net.Conn
	config *SocketServerConfig
}

// Возвращает объект подключения по протоколу TCP
func NewTcpConnection(connection net.Conn, serverConfig *SocketServerConfig) *TcpConnection {
	return &TcpConnection{connection, serverConfig}
}

// Читает строку данных
func (this *TcpConnection) Read() (message string, err error) {
	if this.config.ReadTimeOut > 0 {
		this.conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(this.config.ReadTimeOut)))
	}

	// если не обнаружили перевод строки, то сообщение считается не валидным
	message, err = bufio.NewReader(this.conn).ReadString("\n"[0])
	return
}

// Пишет строку данных
func (this *TcpConnection) Write(message string) error {
	if this.config.WriteTimeOut > 0 {
		this.conn.SetWriteDeadline(time.Now().Add(time.Millisecond * time.Duration(this.config.WriteTimeOut)))
	}

	_, err := this.conn.Write([]byte(message))

	return err
}

// Закрывает соединение
func (this *TcpConnection) Close() error {
	return this.conn.Close()
}
