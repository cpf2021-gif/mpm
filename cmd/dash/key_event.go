package dash

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

type keyEventHandler struct {
	s tcell.Screen
	state *State
	done chan struct{}

	drawer drawer
}

func (h *keyEventHandler) quit() {
	h.s.Fini()
	close(h.done)
	os.Exit(0)
}

func (h *keyEventHandler) HandleKeyEvent(ev *tcell.EventKey) {
	if ev.Key() == tcell.KeyCtrlC {
		h.quit()
	}
}