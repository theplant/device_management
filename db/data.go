package db

import (
	"fmt"
	"github.com/qor/qor"
)

var DeviceCategories = [][]string{
	{"1", "自有设备"},
	{"2", "消耗品"},
	{"3", "客户设备"},
}

func WarehouseCollection(prop interface{}, c *qor.Context) (r [][]string) {
	var whs = []*Warehouse{}
	if err := DB.Find(&whs).Error; err != nil {
		panic(err)
	}
	for _, wh := range whs {
		r = append(r, []string{fmt.Sprintf("%d", wh.ID), wh.Name})
	}
	return
}
