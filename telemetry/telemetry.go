package telemetry

import (
	"time"
)

// Represents one second of telemetry data
type TELEM struct {
	Accl        []ACCL
	Gps         []GPS5
	Gyro        []GYRO
	GpsFix      GPSF
	GpsAccuracy GPSP
	Time        GPSU
	Temp        TMPC
}

// the thing we want, json-wise
// GPS data might have a generated timestamp and derived track
type TELEM_OUT struct {
	*GPS5

    Accuracy    float32 `json:"accuracy"`
    AltAccuracy float32 `json:"altitude_accuracy"`
    Heading     float32 `json:"heading"`
}

var last_good_track float64 = 0

// zeroes out the telem struct
func (t *TELEM) Clear() {
	t.Accl = t.Accl[:0]
	t.Gps = t.Gps[:0]
	t.Gyro = t.Gyro[:0]
	t.Time.Time = time.Time{}
}

// determines if the telem has data
func (t *TELEM) IsZero() bool {
	// hack.
	return t.Time.Time.IsZero()
}

// try to populate a timestamp for every GPS row. probably bogus.
func (t *TELEM) FillTimes(until time.Time) error {
	len := len(t.Gps)
	diff := until.Sub(t.Time.Time)

	offset := diff.Seconds() / float64(len)

	for i, _ := range t.Gps {
		dur := time.Duration(float64(i)*offset*1000) * time.Millisecond
		ts := t.Time.Time.Add(dur)
		t.Gps[i].TS = ts
	}

	return nil
}

func (t *TELEM) ShitJson() []TELEM_OUT {
	var out []TELEM_OUT

	for i, _ := range t.Gps {
		jobj := TELEM_OUT{&t.Gps[i], 0, 0, 0}
		if 0 == i {
			jobj.Accuracy = 1.5
			jobj.AltAccuracy = 1.5
			jobj.Heading = 0
		}

		out = append(out, jobj)
	}

	return out
}
