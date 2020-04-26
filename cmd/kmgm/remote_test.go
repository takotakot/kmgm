package main_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func runKmgmServe(t *testing.T) string {
	basedir, teardown := prepareBasedir(t)
	t.Cleanup(teardown)

	testPort := 34000
	addrPort := fmt.Sprintf("127.0.0.1:%d", testPort)

	ctx, cancel := context.WithCancel(context.Background())

	joinC := make(chan struct{})
	go func() {
		logs, err := runKmgm(t, ctx, basedir, nil, []string{"serve", "--reuse-port", "--listen-addr", addrPort}, nowDefault)
		expectErr(t, err, context.Canceled)
		expectLogMessage(t, logs, "Started listening")
		close(joinC)
	}()

	for i := 0; i < 10; i++ {
		conn, err := net.Dial("tcp", addrPort)
		if err != nil {
			t.Logf("net.Dial(%s) error: %v", addrPort, err)

			time.Sleep(100 * time.Millisecond)
			continue
		}
		conn.Close()
		t.Logf("net.Dial(%s) success", addrPort)
		break
	}

	t.Cleanup(func() {
		cancel()
		<-joinC
	})
	return addrPort
}

func TestServe_Noop(t *testing.T) {
	_ = runKmgmServe(t)
}
