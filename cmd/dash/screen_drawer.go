package dash

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

type ScreenDrawer struct {
	l *LineDrawer
}

func NewScreenDrawer(s tcell.Screen) *ScreenDrawer {
	return &ScreenDrawer{l: NewLineDrawer(0, s)}
}

func (d *ScreenDrawer) Print(s string, style tcell.Style) {
	d.l.Draw(s, style)
}

func (d *ScreenDrawer) Println(s string, style tcell.Style) {
	d.Print(s, style)
	d.NL()
}

// FillLine prints the given rune until the end of the current line
// and adds a newline.
func (d *ScreenDrawer) FillLine(r rune, style tcell.Style) {
	w, _ := d.Screen().Size()
	if w-d.l.col < 0 {
		d.NL()
		return
	}
	s := strings.Repeat(string(r), w-d.l.col)
	d.Print(s, style)
	d.NL()
}

func (d *ScreenDrawer) FillUntil(r rune, style tcell.Style, limit int) {
	if d.l.col > limit {
		return // already passed the limit
	}
	s := strings.Repeat(string(r), limit-d.l.col)
	d.Print(s, style)
}

func (d *ScreenDrawer) NL() {
	d.l.row++
	d.l.col = 0
}

func (d *ScreenDrawer) Screen() tcell.Screen {
	return d.l.s
}

func (d *ScreenDrawer) Goto(x, y int) {
	d.l.col = x
	d.l.row = y
}

func (d *ScreenDrawer) GoToBottom() {
	_, h := d.Screen().Size()
	d.l.row = h - 1
	d.l.col = 0
}

type LineDrawer struct {
	s   tcell.Screen
	row int
	col int
}

func NewLineDrawer(row int, s tcell.Screen) *LineDrawer {
	return &LineDrawer{row: row, col: 0, s: s}
}

func (d *LineDrawer) Draw(s string, style tcell.Style) {
	for _, r := range s {
		d.s.SetContent(d.col, d.row, r, nil, style)
		d.col += runewidth.RuneWidth(r)
	}
}
