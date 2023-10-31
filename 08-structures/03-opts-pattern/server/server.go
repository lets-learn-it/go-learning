package server

type OptFunc func(*Opts)

type Opts struct {
	MaxConn int
	Id      string
	Tls     bool
}

func defaultOpts() Opts {
	return Opts{
		MaxConn: 10,
		Id:      "default",
		Tls:     false,
	}
}

type Server struct {
	Opts
}

func NewServer(ofn ...OptFunc) *Server {
	opts := defaultOpts()
	for _, fn := range ofn {
		fn(&opts)
	}
	return &Server{
		Opts: opts,
	}
}
