package app

import (
	"container/heap"
	"errors"
	"sync"

	"github.com/patrickishaf/job_scheduler/internal/common"
)

type HeapScheduler struct {
	mu   sync.Mutex
	heap *JobHeap
}

func CreateHeapScheduler() *HeapScheduler {
	h := &JobHeap{}
	heap.Init(h)
	return &HeapScheduler{heap: h}
}

// EnqueueDueJobs is called by the cron tick. It receives pending jobs fetched
// from the DB and pushes only the ones that are due into the heap.
func (s *HeapScheduler) EnqueueDueJobs(jobs []*Job) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, job := range jobs {
		if job.IsDue() {
			job.Status = string(PROCESSING_STATUS_QUEUED)
			heap.Push(s.heap, job)
		}
	}
}

func (s *HeapScheduler) Dequeue() (*Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.heap.Len() == 0 {
		return nil, false
	}

	job := heap.Pop(s.heap).(*Job)
	job.Status = string(PROCESSING_STATUS_PROCESSING)
	return job, true
}

func (s *HeapScheduler) RequeueRecurring(job *Job) {
	s.mu.Lock()
	defer s.mu.Unlock()

	job.ScheduledTime = job.NextRun()
	job.AttemptCount = 0
	job.Status = string(PROCESSING_STATUS_PENDING)
	// TODO: Make job to go back to the db as pending
	// Job goes back to DB as pending; cron tick will re-enqueue when due.
	// If you want immediate re-insert (next interval is short), push directly:
	heap.Push(s.heap, job)
}

func (s *HeapScheduler) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.heap.Len()
}

func (s *HeapScheduler) Cancel(jobID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, job := range *s.heap {
		if job.ID.String() == jobID {
			heap.Remove(s.heap, i)
			job.Status = string(PROCESSING_STATUS_FAILED)
			return nil
		}
	}

	return errors.New(common.ErrJobNotFoundInQueue)
}

func (s *HeapScheduler) UpdatePriority(jobID string, newPriority int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, job := range *s.heap {
		if job.ID.String() == jobID {
			job.Priority = newPriority
			heap.Fix(s.heap, i)
			return nil
		}
	}

	return errors.New(common.ErrJobNotFoundInQueue)
}
