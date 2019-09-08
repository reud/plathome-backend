package controller

import (
	"PlatHome-Backend/models"
	"github.com/jinzhu/gorm"
)

type Database struct {
	db *gorm.DB
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		panic(err)
	}

}

func NewDatabase(dialect string, settings string) *Database {
	db, err := gorm.Open(dialect, settings)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.EzRequesterModel{}, &models.Device{})
	return &Database{db: db}
}

func (d *Database) Create(device *models.Device) {
	d.db.Create(device)
}

func (d *Database) FindAll() *[]models.Device {
	var devices []models.Device
	d.db.Find(&devices)
	for i := range devices {
		d.db.Model(devices[i]).Related(&devices[i].EzRequesterModels, "EzRequesterModels")
	}
	return &devices
}

func (d *Database) Delete(id uint) {
	dice := models.Device{}
	dice.ID = id
	d.db.Delete(dice)
}

func (d *Database) Update(device *models.Device) {
	fromRecord := models.Device{}
	fromRecord.ID = device.ID
	d.db.First(&fromRecord)
	d.db.Model(&fromRecord).Update(&device)
	d.db.Save(&device)

}
