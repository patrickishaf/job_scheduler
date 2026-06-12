package app

// :::::::::::::::::::::: job heap ::::::::::::::::::::::::::
type JobHeap []*Job

func (this JobHeap) Len() int {
	return len(this)
}

func (this JobHeap) Less(i, j int) bool {
	return this[i].HasLessPriorityThan(this[j])
}

func (this JobHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (h *JobHeap) Push(x any) {
	*h = append(*h, x.(*Job))
}

func (h *JobHeap) Pop() any {
	old := *h
	n := len(old)
	job := old[n-1]
	old[n-1] = nil
	*h = old[:n-1]
	return job
}

// :::::::::::::::::::::: indexed heap ::::::::::::::::::::::::::
type indexedHeap []*priorityQueueEntry

func (h indexedHeap) Len() int {
	return len(h)
}

func (h indexedHeap) Less(i, j int) bool {
	return h[i].job.HasLessPriorityThan(h[j].job)
}
func (h indexedHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *indexedHeap) Push(x any) {
	item := x.(*priorityQueueEntry)
	item.index = len(*h)
	*h = append(*h, item)
}

func (h *indexedHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[:n-1]
	return item
}
