// ReadyStack Caddy build entrypoint (xcaddy pattern).
//
// This tiny main module depends on github.com/caddyserver/caddy/v2 at a pinned
// version; building it produces the caddy binary with the standard module set.
// Because caddy reads its version from the caddy/v2 dependency in build info,
// the resulting `caddy version` self-reports the pinned upstream tag — no
// -ldflags version stamping required, and -trimpath does not strip it.
package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// Pull in the standard set of Caddy modules (reverse_proxy, file_server,
	// tls, http handlers, etc.). This mirrors the official caddy binary.
	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

func main() {
	caddycmd.Main()
}
