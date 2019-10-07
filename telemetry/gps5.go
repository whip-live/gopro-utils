package telemetry

import (
	"encoding/binary"
	"errors"
	"time"
)

// GPS sentence with lat/lon/alt/speed/3d speed
type GPS5 struct {
	Latitude  float64 `json:"latitude"`    // degrees lat
	Longitude float64 `json:"longitude"`    // degrees lon
	Altitude  float64 `json:"altitude"`    // meters above wgs84 ellipsoid ?
	Speed3D   float64 `json:"speed"` // m/s, standard error?
	TS        time.Time   `json:"timestamp"`
}

func (gps *GPS5) Parse(bytes []byte, scale *SCAL) error {
	if 20 != len(bytes) {
		return errors.New("Invalid length GPS5 packet")
	}

	gps.Latitude = float64(int32(binary.BigEndian.Uint32(bytes[0:4]))) / float64(scale.Values[0])
	gps.Longitude = float64(int32(binary.BigEndian.Uint32(bytes[4:8]))) / float64(scale.Values[1])

	// convert from mm
	gps.Altitude = float64(int32(binary.BigEndian.Uint32(bytes[8:12]))) / float64(scale.Values[2])

	// convert from mm/s
	gps.Speed3D = float64(int32(binary.BigEndian.Uint32(bytes[16:20]))) / float64(scale.Values[4])

	return nil
}
