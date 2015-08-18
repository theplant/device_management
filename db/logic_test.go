package db

import (
	"testing"
	"time"
)

func employeeAndWarehouse() (felix Employee, wensanlu Warehouse) {
	DB.Where(&Employee{Name: "孙凤民"}).Assign(&Employee{}).FirstOrCreate(&felix)
	DB.Where(&Warehouse{Name: "文三路仓库"}).Assign(&Warehouse{}).FirstOrCreate(&wensanlu)
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
