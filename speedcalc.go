package fes

import (
	"math"

	"github.com/umahmood/haversine"
)

const secondsInHour = 3600

func calcKMAndKPH(a, b Point) (float64, float64) {
	deltaT := math.Abs(float64(b.Timestamp) - float64(a.Timestamp))
	_, km := haversine.Distance(haversine.Coord{Lat: a.Lat, Lon: a.Lng}, haversine.Coord{Lat: b.Lat, Lon: b.Lng})
	speed := km / deltaT * secondsInHour
	return km, speed
}
