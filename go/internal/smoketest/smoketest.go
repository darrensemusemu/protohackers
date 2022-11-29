package smoketest

import (
	"fmt"
	"io"
	"net"

	"github.com/darrensemusemu/protohackers/go/pkg/server"
)

func New() *handler {
	return &handler{}
}

// handler implements server.Service connection
var _ server.Service = (*handler)(nil)

type handler struct{}

func (h *handler) HandleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Received: %s\n", conn.RemoteAddr())

	if _, err := io.Copy(conn, conn); err != nil {
		fmt.Printf("server handle connection err: %s\n", err)
	}
}
