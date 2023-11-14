package dash

import (
	"fmt"
	"os"

	"github.com/cpf2021-gif/mpm/model"
	"github.com/gdamore/tcell/v2"
)

type State struct {
	lists []*model.List

	listTableRowIdx int // highlighted row in list table

	pageNum int
}

func Run() {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("failed to create a screen: %v\n", err)
	}
	if err := s.Init(); err != nil {
		fmt.Printf("failed to initialize screen: %v\n", err)
		os.Exit(1)
	}
	s.SetStyle(baseStyle)

	// get data
	db := model.MustNewSqlite()
	listmodel := model.NewListModel(db.DB)
	res, err := listmodel.FindAll()
	if err != nil {
		fmt.Printf("get list data error: %v\n", err)
		os.Exit(1)
	}

	var (
		state = State{lists: res}
	)

	d := dashDrawer{
		s,
	}

	d.Draw(&state)

	for {
		s.Show()
	}
}
