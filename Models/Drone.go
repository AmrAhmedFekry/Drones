package Models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type droneState string

const (
	IDLE       droneState = "IDLE"
	LOADING    droneState = "LOADING"
	LOADED     droneState = "LOADED"
	DELIVERING droneState = "DELIVERING"
	DELIVERED  droneState = "DELIVERED"
	RETURNING  droneState = "RETURNING"
)

type droneModel string

const (
	Lightweight   droneModel = "Lightweight"
	Middleweight  droneModel = "Middleweight"
	Cruiserweight droneModel = "Cruiserweight"
	Heavyweight   droneModel = "Heavyweight"
)

type Drone struct {
	gorm.Model
	SerialNumber    string     `json:"serial_number" gorm:"type:varchar(100) not null;unique"`
	DroneModel      droneModel `json:"model" sql:"drone_model"`
	DroneState      droneState `json:"state" sql:"droneState"`
	WeightLimit     int        `json:"weight_limit"`
	BatteryCapacity int        `json:"battery_capacity" gorm:"type:varchar(3)"`
	DroneLoad       []DroneLoad
}

func (dm *droneModel) Scan(value interface{}) error {
	*dm = droneModel(value.([]byte))
	return nil
}

func (dm droneModel) Value() (driver.Value, error) {
	return string(dm), nil
}

func (ds *droneState) Scan(value interface{}) error {
	*ds = droneState(value.([]byte))
	return nil
}

func (ds droneState) Value() (driver.Value, error) {
	return string(ds), nil
}
