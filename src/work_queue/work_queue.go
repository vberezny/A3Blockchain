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
	// TODO: initialize struct; start nWorkers workers as goroutines
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
		// TODO: listen on the .Jobs channel for incoming tasks
		tasks := <-queue.Jobs
		// TODO: run tasks by calling .Run()
		// TODO: send the return value back on Results channel
		queue.Results <- tasks.Run()
		// TODO: exit (return) when a signal is sent on StopRequests
		if len(queue.StopRequests) > 0 {
			<-queue.StopRequests
			running = false
			return
		}
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// TODO: tell workers to stop processing tasks.
	for i := uint(0); i < queue.NumWorkers; i++ {
		queue.StopRequests <- 1
	}
}
