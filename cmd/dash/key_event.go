package dash

import (
	"math"
	"os"

	"github.com/gdamore/tcell/v2"
)

type keyEventHandler struct {
	s     tcell.Screen
	state *State
	done  chan struct{}

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
	} else if ev.Key() == tcell.KeyUp || ev.Rune() == 'k' {
		h.handleUpkey()
	} else if ev.Key() == tcell.KeyDown || ev.Rune() == 'j' {
		h.handleDownkey()
	} else if ev.Key() == tcell.KeyLeft || ev.Rune() == 'h' {
		h.prev()
	} else if ev.Key() == tcell.KeyRight || ev.Rune() == 'l' {
		h.next()
	}
}

func (h *keyEventHandler) handleUpkey() {
	curPageSize := pageSize

	state := h.state
	last := len(state.lists) - 1
	curPageEnd := state.pageNum*curPageSize + 9

	if curPageEnd <= last {
		state.listTableRowIdx = (state.listTableRowIdx + curPageSize - 1) % curPageSize
	} else {
		curPageSize = (last + 1) % pageSize
		state.listTableRowIdx = (state.listTableRowIdx + curPageSize - 1) % curPageSize
	}
	h.drawer.Draw(state)
}

func (h *keyEventHandler) handleDownkey() {
	curPageSize := pageSize

	state := h.state
	last := len(state.lists) - 1
	curPageEnd := state.pageNum*curPageSize + 9

	if curPageEnd <= last {
		state.listTableRowIdx = (state.listTableRowIdx + curPageSize + 1) % curPageSize
	} else {
		curPageSize = (last + 1) % pageSize
		state.listTableRowIdx = (state.listTableRowIdx + curPageSize + 1) % curPageSize
	}
	h.drawer.Draw(state)
}

func (h *keyEventHandler) prev() {
	state := h.state

	length := len(state.lists)

	maxPage := int(math.Ceil(float64(length) / float64(pageSize)))

	state.pageNum = (state.pageNum + maxPage - 1) % maxPage
	state.listTableRowIdx = -1

	h.drawer.Draw(state)
}

func (h *keyEventHandler) next() {
	state := h.state

	length := len(state.lists)

	maxPage := int(math.Ceil(float64(length) / float64(pageSize)))

	state.pageNum = (state.pageNum + maxPage + 1) % maxPage
	state.listTableRowIdx = -1

	h.drawer.Draw(state)
}
