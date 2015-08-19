package db

import (
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

func employeeAndWarehouse() (felix Employee, wensanlu Warehouse) {
	DB.Where(&Employee{Name: "孙凤民"}).Assign(&Employee{}).FirstOrCreate(&felix)
	DB.Where(&Warehouse{Name: "文三路仓库"}).Assign(&Warehouse{}).FirstOrCreate(&wensanlu)
	return
}

func deviceiPhone(warehouseID uint) (iPhone Device) {
	DB.Where(&Device{Name: "苹果iPhone", Code: "IPHONE6", TotalQuantity: 20, WarehouseID: warehouseID, CategoryID: 1}).Assign(&Device{}).FirstOrCreate(&iPhone)
	return
}

func TestClientDeviceIn(t *testing.T) {
	DB.Unscoped().Delete(&ClientDeviceIn{})
	DB.Unscoped().Delete(&ClientDeviceOut{})
	DB.Unscoped().Delete(&ReportItem{})

	felix, wensanlu := employeeAndWarehouse()
	// t.Error(felix, wensanlu)
	cdIn := &ClientDeviceIn{
		DeviceName: "游标卡尺",
		ClientName: "杭州浩天化工有限公司",
		Quantity:   5,
		Date:       time.Now(),
		Warehouse:  wensanlu,
		ByWhom:     felix,
	}

	DB.Create(cdIn)

	ris := []*ReportItem{}
	DB.Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Find(&ris)
	if len(ris) == 0 {
		t.Error("report item not created")
	}

	cdOut := &ClientDeviceOut{
		ClientDeviceInID: cdIn.ID,
		Date:             time.Now(),
		ByWhom:           felix,
	}

	DB.Create(cdOut)
	DB.Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Find(&ris)
	if len(ris) > 0 {
		t.Error("report item not removed when return")
	}

	if cdOut.WarehouseName != wensanlu.Name {
		t.Error("additional property not assigned.")
	}

	err := DB.Delete(&cdIn).Error
	if err == nil {
		t.Error("shouldn't be able to delete")
	}

	DB.Delete(&cdOut)
	DB.Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Find(&ris)
	if len(ris) == 0 {
		t.Error("report item not recovered")
	}

	DB.Delete(&cdIn)
	DB.Where(&ReportItem{ClientDeviceInID: cdIn.ID}).Find(&ris)
	if len(ris) > 0 {
		t.Error("report item not removed when delete in")
	}

}

func TestDeviceOutAndIn(t *testing.T) {
	gorm.Delete(DB.Unscoped().NewScope(&DeviceIn{}))
	gorm.Delete(DB.Unscoped().NewScope(&DeviceOut{}))
	gorm.Delete(DB.Unscoped().NewScope(&Device{}))
	gorm.Delete(DB.Unscoped().NewScope(&ReportItem{})) // call without callbacks

	felix, wensanlu := employeeAndWarehouse()
	iphone := deviceiPhone(wensanlu.ID)
	from, _ := getOrCreateReportItem(wensanlu, &iphone, 0)

	dOut := DeviceOut{
		FromReportItemID: from.ID,
		ToWhomID:         felix.ID,
		Quantity:         5,
		Date:             time.Now(),
		ByWhomID:         felix.ID,
	}

	DB.Create(&dOut)

	var ri *ReportItem
	ri, _ = getOrCreateReportItem(felix, &iphone, 0)

	if ri.ID == 0 {
		t.Error("report item not created")
	}

	dOut2 := dOut
	dOut2.ID = 0
	DB.Create(&dOut2)
	ri, _ = getOrCreateReportItem(felix, &iphone, 0)

	if ri.Count != 10 {
		t.Error("report item count updated wrong, should be 10")
	}

	DB.Delete(&dOut2)

	ri, _ = getOrCreateReportItem(felix, &iphone, 0)
	if ri.Count != 5 {
		t.Error("report item count updated wrong, should be 5")
	}

	inFrom, _ := getOrCreateReportItem(felix, &iphone, 0)
	dIn := &DeviceIn{
		FromReportItemID: inFrom.ID,
		Quantity:         3,
		ToWarehouseID:    wensanlu.ID,
		Date:             time.Now(),
		ByWhomID:         felix.ID,
	}
	DB.Create(dIn)
	ri, _ = getOrCreateReportItem(felix, &iphone, 0)

	if ri.Count != 2 {
		t.Error("report item count updated wrong, should be 2")
	}

	DB.Delete(&dIn)
	ri, _ = getOrCreateReportItem(felix, &iphone, 0)
	if ri.Count != 5 {
		t.Error("report item count updated wrong, should be 5 again")
	}
}
