package config

type SocketConfig struct {
	ReadBufferSize  int `envconfig:"SOCKET_READ_BUFFER_SIZE" default:"1024"`
	WriteBufferSize int `envconfig:"SOCKET_WRITE_BUFFER_SIZE" default:"1024"`
}
