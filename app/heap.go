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

func (this *JobHeap) Push(x any) {
	*this = append(*this, x.(*Job))
}

func (this *JobHeap) Pop() any {
	old := *this
	n := len(old)
	job := old[n-1]
	old[n-1] = nil
	*this = old[:n-1]
	return job
}

type indexedHeap []*priorityQueueEntry

func (this indexedHeap) Len() int {
	return len(this)
}

func (this indexedHeap) Less(i, j int) bool {
	return this[i].job.HasLessPriorityThan(this[j].job)
}
func (this indexedHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *indexedHeap) Push(x any) {
	item := x.(*priorityQueueEntry)
	item.index = len(*this)
	*this = append(*this, item)
}

func (this *indexedHeap) Pop() any {
	old := *this
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*this = old[:n-1]
	return item
}
