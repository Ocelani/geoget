package geoget

import (
	"github.com/jasonwinn/geocoder"
)

// Route object stores information about one route of
// onde City to another City, as well with its distance and time.
type Route struct {
	From     *City
	To       *City
	Time     string
	Distance float64
}

// * PUBLIC * //

// NewRoute constructor method of *Route.
func NewRoute(from *City, to *City) *Route {
	r := &Route{From: from, To: to}
	r.setDirections()
	return r
}

// Log this *Route info.
func (r *Route) Log() {
	zlog.Info().
		Str("from", r.From.Name).
		Str("to", r.To.Name).
		Float64("distance", r.Distance).
		Send()
}

// * PRIVATE * //

// setDirections collects this *Route data and saves into this object.
func (r *Route) setDirections() {
	var (
		from = r.From.Name
		to   = r.To.Name
	)
	direct := geocoder.NewDirections(from, []string{to})
	direct.RouteType = "shortest"

	// Get the geocode data
	result, err := direct.Get()
	if err != nil {
		zlog.Panic().
			Str("from", from).
			Str("to", to).
			Err(err).
			Send()
	}
	res := result.Route

	// Saves data into this *Route
	r.Time = res.FormattedTime
	r.Distance = res.Distance
}
