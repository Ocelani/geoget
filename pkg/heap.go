package geoget

import "container/heap"

// RoadMapHeap is a min-heap of *RoadMap.
// It ordenates the heap based on its []*RoadMaps distance.
type RoadMapHeap []*RoadMap

// * PUBLIC * //

// NewRoadMapHeap constructor method of *RoadMapHeap.
func NewRoadMapHeap() *RoadMapHeap {
	h := RoadMapHeap([]*RoadMap{})
	heap.Init(&h)
	return &h
}

// PushRoadMaps just pushed the roadmaps into the heap.
func (h *RoadMapHeap) PushRoadMaps(roadmaps []*RoadMap) {
	for _, m := range roadmaps {
		heap.Push(h, m)
	}
}

// Push in a *RoadMap element into the heap.
// It uses pointer receiver because it modifies the slice content and lenght.
func (h *RoadMapHeap) Push(x interface{}) {
	*h = append(*h, x.(*RoadMap))
}

// Pop out one *RoadMap element from the heap.
// It uses pointer receiver because it modifies the slice content and lenght.
func (h *RoadMapHeap) Pop() interface{} {
	old := *h
	len := len(old)
	*h = old[0 : len-1]
	return old[len-1]
}

func (h RoadMapHeap) Len() int           { return len(h) }
func (h RoadMapHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h RoadMapHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
