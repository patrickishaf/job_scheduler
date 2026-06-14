package app

import (
	"container/heap"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

// --- Starvation Prevention ---
// A job that has been pending for longer than this threshold
// gets its effective priority boosted by 1 level (e.g. Low→Medium)
const starvationThreshold = 5 * time.Minute

type HeapScheduler struct {
	mu   sync.Mutex
	heap *JobHeap
}

func NewHeapScheduler() *HeapScheduler {
	h := &JobHeap{}
	heap.Init(h)
	return &HeapScheduler{heap: h}
}

// EnqueueDueJobs is called by the cron tick every second.
// Only pending jobs whose ScheduledTime has passed enter the heap.
// Cancelled jobs are explicitly rejected.
func (s *HeapScheduler) EnqueueDueJobs(jobs []*Job) map[uuid.UUID]bool {
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
		job.Status = string(PROCESSING_STATUS_QUEUED)
		heap.Push(s.heap, job)
		jobIDMap[job.ID] = true
	}

	log.Printf("job ID map: %v", jobIDMap)

	return jobIDMap
}

// Dequeue is called by the worker to pull the next job.
// Returns false if the queue is empty.
func (s *HeapScheduler) Dequeue() (*Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.heap.Len() == 0 {
		return nil, false
	}

	job := heap.Pop(s.heap).(*Job)

	// Guard: if job was cancelled while sitting in the queue, discard it
	if job.Status == string(PROCESSING_STATUS_CANCELLED) {
		return nil, false
	}

	job.Status = string(PROCESSING_STATUS_PROCESSING)
	return job, true
}

// RequeueRecurring re-inserts a recurring job after successful completion.
// The next run time is computed from the interval and pushed back as pending.
// The cron tick will pick it up when it is due.
func (s *HeapScheduler) RequeueRecurring(job *Job) {
	job.ScheduledTime = job.NextScheduledTime()
	job.AttemptCount = 0
	job.Error = nil
	job.Status = string(PROCESSING_STATUS_PENDING)
	// Return to DB as pending — cron tick re-enqueues when due.
	// For short intervals, push directly into heap instead:
	s.mu.Lock()
	defer s.mu.Unlock()
	heap.Push(s.heap, job)
}

// RequeueFromDLQ handles step 10: client manually retries a dead letter job.
// The job is reset and re-enters the queue as pending.
func (s *HeapScheduler) RequeueFromDLQ(job *Job) error {
	if job.Status != string(PROCESSING_STATUS_FAILED) {
		return errors.New("only failed jobs can be requeued from the DLQ")
	}

	job.AttemptCount = 0
	job.RetryCount = 3
	job.Error = nil
	job.Status = string(PROCESSING_STATUS_PENDING)
	job.ScheduledTime = time.Now()

	s.mu.Lock()
	defer s.mu.Unlock()
	heap.Push(s.heap, job)
	return nil
}

// Cancel removes a job from the heap by ID — O(n) scan + O(log n) remove
func (s *HeapScheduler) Cancel(jobID uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, job := range *s.heap {
		if job.ID == jobID {
			heap.Remove(s.heap, i)
			job.Status = string(PROCESSING_STATUS_CANCELLED)
			return nil
		}
	}
	return errors.New("job not found in queue")
}

// UpdatePriority changes a job's priority mid-heap — O(n) scan + O(log n) fix
func (s *HeapScheduler) UpdatePriority(jobID uuid.UUID, newPriority int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, job := range *s.heap {
		if job.ID == jobID {
			job.Priority = newPriority
			heap.Fix(s.heap, i)
			return nil
		}
	}
	return errors.New("job not found in queue")
}

func (s *HeapScheduler) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.heap.Len()
}
