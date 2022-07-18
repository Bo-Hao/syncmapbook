package mapbook

import "sync"

type AskBook struct {
	Book sync.Map
}

type BidBook struct {
	Book sync.Map
}
