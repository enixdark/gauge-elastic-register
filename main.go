package main

import (
	"net"
	"os"

	"github.com/enixdark/gauge-elastic-register/logger"

	gm "github.com/exnidark/gauge-elastic-register/gauge_messages"
	"google.golang.org/grpc"
)

const oneGB = 1024 * 1024 * 1024

func main() {
	findPluginAndProjectRoot()
	if os.Getenv(pluginActionEnv) == executionAction {
		os.Chdir(projectRoot)
		address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
		if err != nil {
			logger.Fatal("failed to start server.")
		}
		l, err := net.ListenTCP("tcp", address)
		if err != nil {
			logger.Fatal("failed to start server.")
		}
		server := grpc.NewServer(grpc.MaxRecvMsgSize(oneGB))
		h := &handler{server: server}
		gm.RegisterReporterServer(server, h)
		logger.Info("Listening on port:%d", l.Addr().(*net.TCPAddr).Port)
		server.Serve(l)
	}
}
