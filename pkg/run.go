package geoget

import (
	"container/heap"
	"os"

	"github.com/jasonwinn/geocoder"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	_topN = 5
	_lvl  = zerolog.DebugLevel
	_zlog = startLogger()
)

func Run(locations []string) {
	data := HandleData(locations)

	h := NewRoadMapHeap()
	h.PushRoadMaps(data)

	for i := 1; i <= _topN; i++ {
		heap.Pop(h).(*RoadMap).Log(i)
	}
}

func PrintResult(heapCH chan *RoadMap, doneCH chan bool) {
	for i := 1; i <= _topN; i++ {
		roadmap := <-heapCH
		roadmap.Log(i)
	}
}

func HandleData(locations []string) []*RoadMap {
	doneCH := make(chan bool)
	defer close(doneCH)

	cities := SetCities(locations)

	go func() {
		for _, c := range cities {
			go func(c *City) {
				doneCH <- c.SetRoutes(cities)
			}(c)
		}
	}()

	for range cities {
		<-doneCH
	}

	return Permute(cities)
}

// SetCities that really exists.
func SetCities(locations []string) []*City {
	cityCH := make(chan *City)
	defer close(cityCH)

	// Busca as infos e coordenadas de cada cidade
	go func() {
		for _, l := range locations {
			go func(l string) {
				_, err := geocoder.FullGeocode(l)

				if err != nil {
					_zlog.Panic().Err(err).Send()
				}

				cityCH <- NewCity(l)
			}(l)
		}
	}()

	cities := []*City{}

	for range locations {
		cities = append(cities, <-cityCH)
	}

	return cities
}

// Permute calls f with each permutation of cities.
func Permute(cities []*City) []*RoadMap {
	roadCH := make(chan *RoadMap)
	defer close(roadCH)

	go permute(roadCH, sliceToMap(cities), 0)

	roadmap := []*RoadMap{}

	f := factorial(len(cities))

	for len(roadmap) < f {
		roadmap = append(roadmap, <-roadCH)
	}

	return roadmap
}

// Permute the values.
func permute(roadCH chan<- *RoadMap, c map[int]*City, i int) {
	if i > len(c) {
		roadCH <- NewRoadMap(c)
		return
	}
	permute(roadCH, c, i+1)

	for j := i + 1; j < len(c); j++ {
		c[i], c[j] = c[j], c[i]
		permute(roadCH, c, i+1)
		c[i], c[j] = c[j], c[i]
	}
}

func sliceToMap(slice []*City) map[int]*City {
	m := map[int]*City{}
	for i, c := range slice {
		m[i] = c
	}
	return m
}

func factorial(n int) int {
	var fact = 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

// startLogger just initializes the logger.
func startLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(_lvl)
	return log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
