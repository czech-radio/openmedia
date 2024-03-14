package internal

import (
	"fmt"
	"log/slog"
	"testing"
)

func (l *LinkedRow) NewNextLink(
	payload LinkPayload) *LinkedRow {
	newRow := new(LinkedRow)
	if l == nil || !l.Initialized {
		// if l == nil {
		// Not initialized
		slog.Debug("initializing first link in sequence")
		newRow.FirstL = newRow
		count := 1
		newRow.RowsCount = &count
		// newRow.Start = nil
		// fmt.Println("krax", *newRow.Start)
		// *newRow.Start = newRow
		newRow.Start = &newRow
		newRow.End = &newRow
		newRow.Payload = payload
	} else {
		// Initialized
		slog.Debug("initializing new link in sequence")
		newRow.Start = l.Start
		// *newRow.Start = *l.PrevL.Start
		newRow.FirstL = l.FirstL
		l.NextL = newRow
		newRow.End = l.End
		*l.End = newRow
		newRow.PrevL = l
	}
	newRow.Payload = payload
	newRow.Initialized = true
	return newRow
}

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
	// nlink := NewLinkSequence(payload)
	var nlink *LinkedRow
	for i := 0; i < 4; i++ {
		payload := LinkPayload{}
		// payload.Index = i
		payload.IndexStr = fmt.Sprintf("%d_ahoj", i)
		nlink = nlink.NextLinkAdd(payload)
	}
	PrintLinks("TEST2", nlink)
	// fmt.Println(*nlink.PrevL.PrevL.PrevStart)
	// *nlink.Start = nlink.PrevL.PrevL.PrevL
	// PrintLinks("TEST3", nlink)
}
