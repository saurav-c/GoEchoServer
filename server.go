package echo

type multiEchoServer struct {
}

// New creates and returns (but does not start) a new MultiEchoServer.
func New() MultiEchoServer {
	return nil
}

func (mes *multiEchoServer) Start(port int) error {
	return nil
}

func (mes *multiEchoServer) Close() {
}

func (mes *multiEchoServer) Count() int {
	return -1
}