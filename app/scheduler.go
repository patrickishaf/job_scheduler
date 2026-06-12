package app

type scheduler interface {
	EnqueueDueJobs(jobs []*Job)
	Dequeue() (*Job, bool)
	Cancel(jobID string) error
	UpdatePriority(jobID string, newPriority int) error
	RequeueRecurring(job *Job)
	Len() int
}
