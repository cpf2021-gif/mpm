package dash

import (
	"fmt"
	"strconv"

	"github.com/cpf2021-gif/mpm/model"
	"github.com/gdamore/tcell/v2"
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

	d.Println("=== Lists ===", baseStyle.Bold(true))
	d.NL()
	drawListTable(d, baseStyle, state)
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
