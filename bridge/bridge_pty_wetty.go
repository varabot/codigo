package bridge

import (
	"io"
	"log"
	"sync"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/asciitransport"
)

/*
func Handler() http.Handler {
	var err error
	shell := os.Getenv("SHELL")

	if shell == "" {
		switch runtime.GOOS {
		case "windows":
			shell = "powershell.exe"
			shell, err = exec.LookPath(shell)
			if err != nil {
				shell, _ = exec.LookPath("cmd.exe")
			}
		default:
			shell = "bash"
			_, err = exec.LookPath(shell)
			if err != nil {
				shell = "sh"
			}
		}
	}
	cmd := []string{shell}

	return &auto{
		fac: factory.New(cmd),
	}
}
*/

func newWetty(cmd string) *auto {
	return &auto{
		fac: factory.New([]string{cmd}),
	}
}

type auto struct {
	fac agent.TtyFactory
}

func (a *auto) serveConn(conn io.ReadWriteCloser) {
	var (
		tryCommandOnce = &sync.Once{}
		cmdCh          = make(chan []string, 1)
		envCh          = make(chan map[string]string, 1)
		resizeCh       = make(chan struct{ rows, cols int }, 4)
	)

	server := asciitransport.Server(conn)
	// send
	// case output:

	// recv
	go func() {
		for {
			var (
				re   = <-server.ResizeEvent()
				rows = int(re.Height)
				cols = int(re.Width)
			)
			tryCommandOnce.Do(func() {
				cmdCh <- re.Command
				envCh <- re.Env
			})
			resizeCh <- struct{ rows, cols int }{rows, cols}
		}
		server.Close()
	}()

	cmd := <-cmdCh
	env := <-envCh

	_ = cmd
	_ = env

	term, err := a.fac.MakeTty()
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for {
			re := <-resizeCh
			err := term.Resize(re.rows, re.cols)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	opts := []asciitransport.Opt{
		asciitransport.WithReader(term),
		asciitransport.WithWriter(term),
		// asciitransport.WithLogger(os.Stderr),
	}
	server.ApplyOpts(opts...)

	<-server.Done()
	term.Close()
}
