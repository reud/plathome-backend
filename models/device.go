package models

import (
	"PlatHome-Backend/gen/models"
	"github.com/jinzhu/gorm"
)

// for GORM (Database)
type Device struct {
	Description       string             `json:"description"`
	EzRequesterModels []EzRequesterModel `json:"ezRequesterModels"`

	Hostname string `json:"hostname"`
	IP       string `json:"ip"`

	Type string `json:"type"`
	gorm.Model
}

type EzRequesterModel struct {
	DeviceID uint
	*models.EzRequesterModel
	gorm.Model
}

func NewDevice(md *models.Device) Device {

	d := Device{
		Description: *md.Description,
		Hostname:    *md.Hostname,
		IP:          *md.IP,
		Type:        *md.Type,
		Model:       gorm.Model{},
	}
	id := d.ID
	var em []EzRequesterModel
	for _, m := range md.EzRequesterModels {
		em = append(em, EzRequesterModel{
			DeviceID:         id,
			EzRequesterModel: m,
			Model:            gorm.Model{},
		})
	}
	d.EzRequesterModels = em
	return d
}

func convertDevice(d *Device) *models.Device {
	var erms []*models.EzRequesterModel
	for _, el := range d.EzRequesterModels {
		erms = append(erms, &models.EzRequesterModel{
			Parameter: el.Parameter,
			Protocol:  el.Protocol,
		})
	}

	md := &models.Device{
		Description:       &d.Description,
		EzRequesterModels: erms,
		Hostname:          &d.Hostname,
		IP:                &d.IP,
		Type:              &d.Type,
	}
	return md
}

func ConvertDevices(ds []Device) []*models.Device {
	var mds []*models.Device
	if ds == nil {
		return []*models.Device{}
	}
	for i, _ := range ds {
		mds = append(mds, convertDevice(&ds[i]))
	}
	return mds
}
