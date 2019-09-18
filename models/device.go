package models

import (
	"plathome-backend/gen/models"
)

// for GORM (Database)
type Device struct {
	Description       string             `json:"description"`
	EzRequesterModels []EzRequesterModel `json:"ezRequesterModels" gorm:"foreignkey:UserRefer"`

	Hostname string `json:"hostname"`
	IP       string `json:"ip" gorm:"primary_key"`

	Type  string `json:"type"`
	State string `json:"state"`
}

type EzRequesterModel struct {
	*models.EzRequesterModel
	UserRefer string
}

func NewDevice(md *models.Device) Device {

	d := Device{
		Description: *md.Description,
		Hostname:    *md.Hostname,
		IP:          *md.IP,
		Type:        *md.Type,
		State:       *md.State,
	}
	ip := d.IP
	var em []EzRequesterModel
	for _, m := range md.EzRequesterModels {
		em = append(em, EzRequesterModel{
			UserRefer:        ip,
			EzRequesterModel: m,
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
		State:             &d.State,
	}
	return md
}

func ConvertDevices(ds []Device) []*models.Device {
	var mds []*models.Device
	if ds == nil {
		return []*models.Device{}
	}
	for i := range ds {
		mds = append(mds, convertDevice(&ds[i]))
	}
	return mds
}
