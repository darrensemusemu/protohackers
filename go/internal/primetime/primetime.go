package primetime

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"net"

	"github.com/darrensemusemu/protohackers/go/pkg/server"
)

func New() *handler {
	return &handler{}
}

// handler implements server.Service interface
var _ server.Service = (*handler)(nil)

type handler struct{}

func (h *handler) HandleConnection(conn net.Conn) {
	defer conn.Close()

	malformedRes := []byte("{}\n")

	buf := bufio.NewScanner(conn)
	for buf.Scan() {
		var req Request
		b := buf.Bytes()
		fmt.Printf("primetime handling request: %s\n", string(b))

		if err := json.Unmarshal(b, &req); err != nil {
			conn.Write(malformedRes)
			return
		}

		if req.Method != "isPrime" || req.Number == nil {
			conn.Write(malformedRes)
			return
		}

		fmt.Printf("isPrime:  %v\n", isPrime(*req.Number))
		handleResponse(conn, isPrime(*req.Number))
	}

	if err := buf.Err(); err != nil {
		fmt.Println("err", err)
		return
	}
}

func handleResponse(conn net.Conn, isPrime bool) {
	res := Response{
		Method: "isPrime",
		Prime:  isPrime,
	}

	b, _ := json.Marshal(res)
	b = append(b, "\n"...)
	if _, err := conn.Write(b); err != nil {
		fmt.Printf("primetime handle response err: %s\n", err)
	}
}

func isPrime(n float64) bool {
	if n <= 1 || float64(int(n)) != n {
		return false
	}
	if n <= 3 {
		return true
	}

	for i := 2.0; i < math.Sqrt(n)+1; i++ {
		if math.Mod(n, i) == 0 {
			return false
		}
	}

	return true
}

type Request struct {
	Method string   `json:"method"`
	Number *float64 `json:"number"`
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}
