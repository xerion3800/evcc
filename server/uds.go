//go:build !windows

package server

import (
	"net"
	"net/http"
	"os"

	"github.com/xerion3800/evcc/cmd/shutdown"
	"github.com/xerion3800/evcc/core/site"
)

// SocketPath is the unix domain socket path
const SocketPath = "/tmp/evcc"

// removeIfExists deletes file if it exists or fails
func removeIfExists(file string) {
	if _, err := os.Stat(file); err == nil {
		if err := os.RemoveAll(file); err != nil {
			log.FATAL.Fatal(err)
		}
	}
}

// HealthListener attaches listener to unix domain socket and runs listener
func HealthListener(site site.API) {
	removeIfExists(SocketPath)

	l, err := net.Listen("unix", SocketPath)
	if err != nil {
		log.FATAL.Fatal(err)
	}

	mux := http.NewServeMux()
	httpd := http.Server{Handler: mux}
	mux.HandleFunc("/health", healthHandler(site))

	go func() { _ = httpd.Serve(l) }()

	shutdown.Register(func() {
		_ = l.Close()
		removeIfExists(SocketPath) // cleanup
	})
}
