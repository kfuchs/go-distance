package main

import (
	"testing"
)

// Tests that a call to NewPoint should return a pointer to a Point with the specified values assigned correctly.
func TestNewCoords(t *testing.T) {
	p := NewCoords(40.622712, -74.016590)

	if p == nil {
		t.Error("Expected to get a pointer to a new point, but got nil instead.")
	}

	if p.lat != 40.622712 {
		t.Errorf("Expected to be able to specify 40.5 as the lat value of a new point, but got %f instead", p.lat)
	}

	if p.lon != -74.016590 {
		t.Errorf("Expected to be able to specify 120.5 as the lng value of a new point, but got %f instead", p.lon)
	}
}

// Tests the Haversine method for a known point and ensures it's accuracy is within 0.3%
func TestHaversine(t *testing.T) {
	whiteHouse := &Coords{lat: 38.897628, lon: -77.036617}
	marthsVinyard := &Coords{lat: 41.380630, lon: -70.645639}

	dist := whiteHouse.Haversine(marthsVinyard)

	// Distance should be 609.19 km
	distanceShouldBe := 609.19

	withinRange := dist < distanceShouldBe*1.003 && dist < distanceShouldBe*1.003

	if !withinRange {
		t.Error("Result not within acceptable range")
	}
}

//
//func TestPointAtDistanceAndBearing(t *testing.T) {
//	sea := &Point{lat: 47.44745785, lng: -122.308065668024}
//	p := sea.PointAtDistanceAndBearing(1090.7, 180)
//
//	// Expected results of transposing point
//	// ~1091km at bearing of 180 degrees
//	resultLat := 37.638557
//	resultLng := -122.308066
//
//	withinLatBounds := p.lat < resultLat+0.001 && p.lat > resultLat-0.001
//	withinLngBounds := p.lng < resultLng+0.001 && p.lng > resultLng-0.001
//	if !(withinLatBounds && withinLngBounds) {
//		t.Error("Unnacceptable result.", fmt.Sprintf("[%f, %f]", p.lat, p.lng))
//	}
//}
//
//// Enures that a point can be marhalled into JSON
//func TestMarshalJSON(t *testing.T) {
//	p := NewPoint(40.7486, -73.9864)
//	res, err := json.Marshal(p)
//
//	if err != nil {
//		log.Print(err)
//		t.Error("Should not encounter an error when attempting to Marshal a Point to JSON")
//	}
//
//	if string(res) != `{"lat":40.7486,"lng":-73.9864}` {
//		t.Error("Point should correctly Marshal to JSON")
//	}
//}
//
//// Enures that a point can be unmarhalled from JSON
//func TestUnmarshalJSON(t *testing.T) {
//	data := []byte(`{"lat":40.7486,"lng":-73.9864}`)
//	p := &Point{}
//	err := p.UnmarshalJSON(data)
//
//	if err != nil {
//		t.Errorf("Should not encounter an error when attempting to Unmarshal a Point from JSON")
//	}
//
//	if p.Lat() != 40.7486 || p.Lng() != -73.9864 {
//		t.Errorf("Point has mismatched data after Unmarshalling from JSON")
//	}
//}
