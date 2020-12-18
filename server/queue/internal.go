package queue

import (
	"github.com/enriquebris/goconcurrentqueue"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"time"
)

type internalQueue struct {
	queue *goconcurrentqueue.FIFO
}

func newInternalQueue() Queue {
	queue := goconcurrentqueue.NewFIFO()
	return &internalQueue{queue}
}

func (q *internalQueue) Enqueue(archive *domain.Archive) error {
	logger := common.GetLogger()
	logger.Printf("enqueue - %v", archive)
	return q.queue.Enqueue(archive)
}

func (q *internalQueue) Channel() (chan *domain.Archive, error) {
	logger := common.GetLogger()
	ch := make(chan *domain.Archive)
	go func() {
		for {
			elem, err := q.queue.DequeueOrWaitForNextElement()
			if err != nil {
				logger.Printf("dequeue error - %v", err)
			}
			logger.Printf("received - %v", elem)
			archive := elem.(*domain.Archive)
			ch <- archive
			time.Sleep(time.Second)
		}
	}()
	return ch, nil
}

func (q *internalQueue) Close() {
	// nothing..
}
