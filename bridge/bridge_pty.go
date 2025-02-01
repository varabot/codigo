package bridge

import (
	"tractor.dev/toolkit-go/duplex/rpc"
)

func (api *Bridge) Terminal(r rpc.Responder, c *rpc.Call) {
	c.Receive(nil)

	wetty := newWetty("/bin/bash")

	ch, err := r.Continue()
	if err != nil {
		panic(err)
	}

	wetty.serveConn(ch)

	ch.Close()
}
