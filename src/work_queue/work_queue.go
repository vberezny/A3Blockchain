package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs         chan Worker
	Results      chan interface{}
	StopRequests chan int
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{})
	q.StopRequests = make(chan int, nWorkers)
	q.NumWorkers = nWorkers
	for i := uint(0); i < nWorkers; i++ {
		go q.worker()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	running := true
	// Run tasks from the Jobs channel, unless we have been asked to stop.
	for running {
		tasks := <-queue.Jobs
		queue.Results <- tasks.Run()
		if len(queue.StopRequests) > 0 {
			<-queue.StopRequests
			running = false
			return
		}
	}
}

// put the work into the Jobs channel so a worker can find it and start the task.
func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
}

// tell workers to stop processing tasks.
func (queue WorkQueue) Shutdown() {
	for i := uint(0); i < queue.NumWorkers; i++ {
		queue.StopRequests <- 1
	}
}
