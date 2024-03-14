package internal

import (
	"testing"
)

func TestNewNextLink(t *testing.T) {
	payload := LinkPayload{}
	nlink := new(LinkedRow)
	for i := 0; i < 4; i++ {
		payload.Index = i
		nlink = nlink.NewNextLink(payload)
	}
	// NOTE: Cannot solve start inside NewNextLink. Start mast be changed after sequnece initialized. Why? Inside loop the start will be shifted by one for each iteration.
	*nlink.Start = nlink.PrevL
	PrintLinks("TEST", nlink)
}

func TestNewLinkSequence(t *testing.T) {
	payload := LinkPayload{}
	nlink := NewLinkSequence(payload)
	for i := 0; i < 4; i++ {
		payload.Index = i
		nlink = nlink.NextLinkAdd(payload)
	}
	PrintLinks("TEST2", nlink)
	// fmt.Println(*nlink.PrevL.PrevL.PrevStart)
	// *nlink.Start = nlink.PrevL.PrevL.PrevL
	// PrintLinks("TEST3", nlink)
}
