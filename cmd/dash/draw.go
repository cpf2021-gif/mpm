package dash

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/cpf2021-gif/mpm/model"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

var (
	baseStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
)

type drawer interface {
	Draw(state *State)
}

type dashDrawer struct {
	s tcell.Screen
}

func (dd *dashDrawer) Draw(state *State) {
	s := dd.s
	s.Clear()
	d := NewScreenDrawer(s)

	switch state.view {
	case viewTypeLists:
		d.Println("=== Lists ===", baseStyle.Bold(true))
		d.NL()
		drawListTable(d, baseStyle, state)

	case viewTypeListDetail:
		d.Println("=== List Detail ===", baseStyle.Bold(true))
		d.NL()
		d.Println(fmt.Sprintf("App: %s  Account: %s", state.selectList.App, state.selectList.Account), baseStyle)
		drawModal(d, state)
	}

	d.GoToBottom()
	drawFooter(d, state)
}

func drawFooter(d *ScreenDrawer, state *State) {
	style := baseStyle.Background(tcell.ColorDarkSlateGray).Foreground(tcell.ColorWhite)
	switch state.view {
	case viewTypeLists:
		d.Print("<Ctrl+C>: Exit", style)
	case viewTypeListDetail:
		d.Print("<q>: GoBack    <Ctrl+C>: Exit", style)
	}

}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	tmpl := fmt.Sprintf("%%-%ds ", padding)
	return fmt.Sprintf(tmpl, s)
}

// lpad adds padding to the left of a string.
func lpad(s string, padding int) string {
	tmpl := fmt.Sprintf("%%%ds ", padding)
	return fmt.Sprintf(tmpl, s)
}

var listColumnConfigs = []*columnConfig[*model.List]{
	{"ID", alignRight, func(l *model.List) string { return strconv.Itoa(int(l.ID)) }},
	{"App", alignLeft, func(l *model.List) string { return l.App }},
	{"Account", alignLeft, func(l *model.List) string { return l.Account }},
	{"Password", alignLeft, func(l *model.List) string { return l.Password }},
}

func drawListTable(d *ScreenDrawer, style tcell.Style, state *State) {
	drawTable(d, style, listColumnConfigs, state.lists, state.listTableRowIdx, state.pageNum)
}

func drawModal(d *ScreenDrawer, state *State) {
	if state.showModal {
		fns := []func(d *modalRowDrawer){
			func(d *modalRowDrawer) { d.Print("=== Info ===", baseStyle.Bold(true)) },
			func(d *modalRowDrawer) { d.Print("", baseStyle) },
			func(d *modalRowDrawer) {
				d.Print("this is a test", baseStyle)
			},
		}
		withModal(d, fns)
	}
}

type modalRowDrawer struct {
	d        *ScreenDrawer
	width    int // current width occupied by content
	maxWidth int
}

// Note: s should not include newline
func (d *modalRowDrawer) Print(s string, style tcell.Style) {
	if d.width >= d.maxWidth {
		return // no longer write to this row
	}
	if d.width+runewidth.StringWidth(s) > d.maxWidth {
		s = truncate(s, d.maxWidth-d.width)
	}
	d.d.Print(s, style)
}

// withModal draws a modal with the given functions row by row.
func withModal(d *ScreenDrawer, rowPrintFns []func(d *modalRowDrawer)) {
	w, h := d.Screen().Size()
	var (
		modalWidth  = int(math.Floor(float64(w) * 0.6))
		modalHeight = int(math.Floor(float64(h) * 0.6))
		rowOffset   = int(math.Floor(float64(h) * 0.2)) // 20% from the top
		colOffset   = int(math.Floor(float64(w) * 0.2)) // 20% from the left
	)
	if modalHeight < 3 {
		return // no content can be shown
	}
	d.Goto(colOffset, rowOffset)
	d.Print(string(tcell.RuneULCorner), baseStyle)
	d.Print(strings.Repeat(string(tcell.RuneHLine), modalWidth-2), baseStyle)
	d.Print(string(tcell.RuneURCorner), baseStyle)
	d.NL()
	rowDrawer := modalRowDrawer{
		d:        d,
		width:    0,
		maxWidth: modalWidth - 4, /* borders + paddings */
	}
	for i := 1; i < modalHeight-1; i++ {
		d.Goto(colOffset, rowOffset+i)
		d.Print(fmt.Sprintf("%c ", tcell.RuneVLine), baseStyle)
		if i <= len(rowPrintFns) {
			rowPrintFns[i-1](&rowDrawer)
		}
		d.FillUntil(' ', baseStyle, colOffset+modalWidth-2)
		d.Print(fmt.Sprintf(" %c", tcell.RuneVLine), baseStyle)
		d.NL()
	}
	d.Goto(colOffset, rowOffset+modalHeight-1)
	d.Print(string(tcell.RuneLLCorner), baseStyle)
	d.Print(strings.Repeat(string(tcell.RuneHLine), modalWidth-2), baseStyle)
	d.Print(string(tcell.RuneLRCorner), baseStyle)
	d.NL()
}

// truncates s if s exceeds max length.
func truncate(s string, max int) string {
	if runewidth.StringWidth(s) <= max {
		return s
	}
	return string([]rune(s)[:max-1]) + "…"
}
