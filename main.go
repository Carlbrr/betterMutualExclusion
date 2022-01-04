package main

import (
	"sync"

	"github.com/Carlbrr/disys_BetterMutualExclusion/proto"
)

type STATE int

var wait *sync.WaitGroup

func init() {
	wait = &sync.WaitGroup{}
}

type Node struct {
	name         string
	id           int
	state        STATE
	timestamp    int
	ports        []string
	replyCounter int
	queue        customQueue
	protoNode    proto.Node
	mutex        sync.Locker
	proto.UnimplementedExclusionServiceServer
}

const (
	RELEASED STATE = iota
	WANTED
	HELD
)
