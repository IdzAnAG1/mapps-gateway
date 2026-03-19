package conf

// Bootstrap - корневая конфигурация приложения.
// Структуры загружаются из configs/config.yaml через Kratos config.
// Сгенерировать conf.pb.go из conf.proto можно командой: make config
type Bootstrap struct {
	Server *Server `json:"server"`
	Data   *Data   `json:"data"`
}

type Server struct {
	Http *Server_HTTP `json:"http"`
	Grpc *Server_GRPC `json:"grpc"`
}

type Server_HTTP struct {
	Network string `json:"network"`
	Addr    string `json:"addr"`
	Timeout string `json:"timeout"`
}

type Server_GRPC struct {
	Network string `json:"network"`
	Addr    string `json:"addr"`
	Timeout string `json:"timeout"`
}

type Data struct {
	Redis        *Data_Redis           `json:"redis"`
	Auth         *Data_ServiceEndpoint `json:"auth"`
	Product      *Data_ServiceEndpoint `json:"product"`
	AssetManager *Data_ServiceEndpoint `json:"asset_manager"`
}

type Data_Redis struct {
	Network      string `json:"network"`
	Addr         string `json:"addr"`
	ReadTimeout  string `json:"read_timeout"`
	WriteTimeout string `json:"write_timeout"`
}

// Data_ServiceEndpoint - конфиг подключения к downstream gRPC сервису.
type Data_ServiceEndpoint struct {
	Network string `json:"network"`
	Addr    string `json:"addr"`
	Timeout string `json:"timeout"`
}

// GetAddr возвращает адрес сервиса.
func (s *Data_ServiceEndpoint) GetAddr() string {
	if s == nil {
		return ""
	}
	return s.Addr
}

// GetAddr для HTTP-сервера
func (s *Server_HTTP) GetAddr() string {
	if s == nil {
		return ""
	}
	return s.Addr
}

// GetNetwork для HTTP-сервера
func (s *Server_HTTP) GetNetwork() string {
	if s == nil {
		return ""
	}
	return s.Network
}

// GetAddr для gRPC-сервера
func (s *Server_GRPC) GetAddr() string {
	if s == nil {
		return ""
	}
	return s.Addr
}

// GetNetwork для gRPC-сервера
func (s *Server_GRPC) GetNetwork() string {
	if s == nil {
		return ""
	}
	return s.Network
}
