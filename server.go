package main

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"math"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/distance", getDistance)

	m.Run()
}

func getDistance(Response render.Render, Request *http.Request) {
	input := Request.URL.Query()
	lat1 := input.Get("lat1")
	lon1 := input.Get("lon1")
	lat2 := input.Get("lat2")
	lon2 := input.Get("lon2")

	from := NewCoords(lat1, lon1)
	to := NewCoords(lat2, lon2)
	dist := from.Haversine(to)
	Response.JSON(200, map[string]interface{}{"distance": dist})
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
