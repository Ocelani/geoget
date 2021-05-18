package geoget

import (
	"container/heap"
	"os"

	"github.com/jasonwinn/geocoder"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	_topN = 5 // Rank sample
	_lvl  = 1 // Log "info" level
)

// Logger
var zlog = startLogger()

// * PUBLIC * //

// Run executes the application algorithm.
func Run(locations []string) {
	cities := FindCities(locations) // Get cities data
	FindRoutes(cities)              // Get routes data and saves in "cities" var
	roadmap := Permute(cities)      // Get every possible route
	h := NewRoadMapHeap()           // Initializes the heap of roadmaps
	h.PushRoadMaps(roadmap)         // Push roadmap data into the heap
	for i := 1; i <= _topN; i++ {   // Prints the ordered rank data sample
		heap.Pop(h).(*RoadMap).Log(i)
	}
}

// FindCities collects data from locations and stores into an array of *City.
func FindCities(locations []string) []*City {
	cityCH := make(chan *City)
	defer close(cityCH)

	go func() {
		// Search for locations geo refferences
		for _, l := range locations {
			go func(l string) {

				_, err := geocoder.FullGeocode(l)
				if err != nil {
					zlog.Panic().Err(err).Send()
				}

				cityCH <- NewCity(l) // Send data in channel
			}(l)
		}
	}()

	cities := []*City{}

	// Waits untill receive all cities data in channel
	for range locations {
		cities = append(cities, <-cityCH)
	}

	return cities
}

// FindRoutes searches for routes from each city to the other cities.
// It doesn't return any value, but it saves in each city pointer.
func FindRoutes(cities []*City) {
	doneCH := make(chan bool)
	defer close(doneCH)

	go func() {
		// For each city, it searches and saves on city pointer
		for _, c := range cities {
			go func(c *City) {
				doneCH <- c.SetRoutes(cities)
			}(c)
		}
	}()

	for range cities {
		<-doneCH
	}
}

// * PRIVATE * //

// startLogger just initializes the logger of this app.
func startLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(_lvl)
	return log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
