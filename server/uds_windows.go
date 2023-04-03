//go:build windows

package server

import "github.com/xerion3800/evcc/core/site"

// HealthListener attaches listener to unix domain socket
func HealthListener(_ site.API, _ int) {
	// nop
}
