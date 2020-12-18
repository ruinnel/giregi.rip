package queue

import (
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
)

type Queue interface {
	Enqueue(archive *domain.Archive) error
	Channel() (chan *domain.Archive, error)
	Close()
}

func NewQueue(config *common.Config) Queue {
	if config.Platform == common.PLATFORM_SERVER {
		return newRabbitQueue(config)
	} else {
		return newInternalQueue()
	}
}
