package model

import (
	"gorm.io/gorm"
)

type List struct {
	ID       uint
	App      string
	Account  string
	Password string
}

func (m *List) TableName() string {
	return "list"
}

type ListModel struct {
	db *gorm.DB
}

func NewListModel(db *gorm.DB) *ListModel {
	return &ListModel{
		db: db,
	}
}

func (m *ListModel) FindAll() ([]*List, error) {
	var lists []*List
	err := m.db.Model(&List{}).Find(&lists).Error
	return lists, err
}
