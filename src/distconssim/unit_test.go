package distconssim

import (
	"fmt"
	"testing"

	u "github.com/shamuelmanrrique/petrynet/src/utils"
)

// TestConnections create connections
func TestConnect(t *testing.T) {
	var LocalIPs = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003",
		"127.0.1.1:5004", "127.0.1.1:5005", "127.0.1.1:5006"}
	conn := u.NewConnec(LocalIPs)
	fmt.Println(*conn[1])
}