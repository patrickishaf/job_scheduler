package app

import (
	"container/heap"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type IndexedPQScheduler struct {
	mu    sync.Mutex
	heap  *indexedHeap
	index map[uuid.UUID]*priorityQueueEntry
}

func NewIndexedPQScheduler() *IndexedPQScheduler {
	h := &indexedHeap{}
	heap.Init(h)
	return &IndexedPQScheduler{
		heap:  h,
		index: make(map[uuid.UUID]*priorityQueueEntry),
	}
}

// EnqueueDueJobs is called by the cron tick every second.
// Idempotent — skips jobs already in the index.
// Cancelled jobs are explicitly rejected.
func (s *IndexedPQScheduler) EnqueueDueJobs(jobs []*Job) map[uuid.UUID]bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	jobIDMap := make(map[uuid.UUID]bool)

	for _, job := range jobs {
		if job.Status == string(PROCESSING_STATUS_CANCELLED) {
			continue
		}

		if job.Status != string(PROCESSING_STATUS_PENDING) {
			continue
		}

		if !job.IsDue() {
			continue
		}

		if _, exists := s.index[job.ID]; exists {
			continue // already queued
		}

		item := &priorityQueueEntry{job: job}
		heap.Push(s.heap, item)
		s.index[job.ID] = item
		job.Status = string(PROCESSING_STATUS_QUEUED)
		jobIDMap[job.ID] = true
	}

	return jobIDMap
}

// Dequeue pulls the next job for a worker.
func (s *IndexedPQScheduler) Dequeue() (*Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.heap.Len() == 0 {
		return nil, false
	}

	item := heap.Pop(s.heap).(*priorityQueueEntry)
	delete(s.index, item.job.ID)

	if item.job.Status == string(PROCESSING_STATUS_CANCELLED) {
		return nil, false
	}

	item.job.Status = string(PROCESSING_STATUS_PROCESSING)
	return item.job, true
}

// RequeueRecurring re-inserts a completed recurring job.
func (s *IndexedPQScheduler) RequeueRecurring(job *Job) {
	job.ScheduledTime = job.NextScheduledTime()
	job.AttemptCount = 0
	job.Error = nil
	job.Status = string(PROCESSING_STATUS_QUEUED)

	s.mu.Lock()
	defer s.mu.Unlock()

	item := &priorityQueueEntry{job: job}
	heap.Push(s.heap, item)
	s.index[job.ID] = item
}

// RequeueFromDLQ handles step 10: client manually retries a dead letter job.
func (s *IndexedPQScheduler) RequeueFromDLQ(job *Job) error {
	if job.Status != string(PROCESSING_STATUS_FAILED) {
		return errors.New("only failed jobs can be requeued from the DLQ")
	}

	job.AttemptCount = 0
	job.RetryCount = 3
	job.Error = nil
	job.Status = string(PROCESSING_STATUS_QUEUED)
	job.ScheduledTime = time.Now()

	s.mu.Lock()
	defer s.mu.Unlock()

	item := &priorityQueueEntry{job: job}
	heap.Push(s.heap, item)
	s.index[job.ID] = item
	return nil
}

// Cancel removes a job from the queue by ID — O(log n)
func (s *IndexedPQScheduler) Cancel(jobID uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.index[jobID]
	if !exists {
		return errors.New("job not found in queue")
	}

	heap.Remove(s.heap, item.index)
	delete(s.index, jobID)
	item.job.Status = string(PROCESSING_STATUS_CANCELLED)
	return nil
}

// UpdatePriority changes a job's priority mid-queue — O(log n)
func (s *IndexedPQScheduler) UpdatePriority(jobID uuid.UUID, newPriority int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.index[jobID]
	if !exists {
		return errors.New("job not found in queue")
	}

	item.job.Priority = newPriority
	heap.Fix(s.heap, item.index)
	return nil
}

func (s *IndexedPQScheduler) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.heap.Len()
}
