package main

import (
	"math"
	"net/http"
)

func main() {
}

type Coords struct {
	lat float64
	lon float64
}

const (
	EARTH_RADIUS = 6371
)

func NewCoords(lat float64, lon float64) *Coords {
	return &Coords{lat: lat, lon: lon}
}

func (p *Coords) Haversine(p2 *Coords) float64 {
	dLat := (p2.lat - p.lat) * (math.Pi / 180.0)
	dLon := (p2.lon - p.lon) * (math.Pi / 180.0)

	lat1 := p.lat * (math.Pi / 180.0)
	lat2 := p2.lat * (math.Pi / 180.0)

	a1 := math.Sin(dLat/2) * math.Sin(dLat/2)
	a2 := math.Sin(dLon/2) * math.Sin(dLon/2) * math.Cos(lat1) * math.Cos(lat2)

	a := a1 + a2

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EARTH_RADIUS * c
}
