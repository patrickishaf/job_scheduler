package app

import (
	"container/heap"
	"errors"
	"sync"

	"github.com/patrickishaf/job_scheduler/internal/common"
)

type IndexedPQScheduler struct {
	mu    sync.Mutex
	heap  *indexedHeap
	index map[string]*priorityQueueEntry
}

func CreateIndexedPQScheduler() *IndexedPQScheduler {
	h := &indexedHeap{}
	heap.Init(h)
	return &IndexedPQScheduler{
		heap:  h,
		index: make(map[string]*priorityQueueEntry),
	}
}

// EnqueueDueJobs is called by the cron tick.
// Skips jobs already in the queue (idempotent).
func (s *IndexedPQScheduler) EnqueueDueJobs(jobs []*Job) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, job := range jobs {
		if !job.IsDue() {
			continue
		}
		if _, exists := s.index[job.ID.String()]; exists {
			continue
		}

		item := &priorityQueueEntry{job: job}
		heap.Push(s.heap, item)
		s.index[job.ID.String()] = item
		job.Status = string(PROCESSING_STATUS_QUEUED)
	}
}

func (s *IndexedPQScheduler) Dequeue() (*Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.heap.Len() == 0 {
		return nil, false
	}

	item := heap.Pop(s.heap).(*priorityQueueEntry)
	delete(s.index, item.job.ID.String())
	item.job.Status = string(PROCESSING_STATUS_PROCESSING)
	return item.job, true
}

func (s *IndexedPQScheduler) Cancel(jobID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.index[jobID]
	if !exists {
		return errors.New(common.ErrJobNotFoundInQueue)
	}

	heap.Remove(s.heap, item.index)
	delete(s.index, jobID)
	item.job.Status = string(PROCESSING_STATUS_PENDING)
	return nil
}

// UpdatePriority changes a job's priority mid-queue — O(log n).
// Useful if you want to escalate a retried job.
func (s *IndexedPQScheduler) UpdatePriority(jobID string, newPriority int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.index[jobID]
	if !exists {
		return errors.New(common.ErrJobNotFoundInQueue)
	}

	item.job.Priority = newPriority
	heap.Fix(s.heap, item.index)
	return nil
}

// RequeueRecurring re-inserts a recurring job after completion.
func (s *IndexedPQScheduler) RequeueRecurring(job *Job) {
	s.mu.Lock()
	defer s.mu.Unlock()

	job.ScheduledTime = job.NextScheduledTime()
	job.AttemptCount = 0
	job.Status = string(PROCESSING_STATUS_QUEUED)

	item := &priorityQueueEntry{job: job}
	heap.Push(s.heap, item)
	s.index[job.ID.String()] = item
}

func (s *IndexedPQScheduler) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.heap.Len()
}
