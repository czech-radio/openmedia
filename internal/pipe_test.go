package internal

import (
	"path/filepath"
	"testing"
)

// One chan result
func TestConsumeResult(t *testing.T) {
	dir := filepath.Join(TEST_DATA_DIR_SRC, "rundowns_additional")
	// var resultChan chan Result = make(chan Result)
	resultChan := ProduceResult(dir)
	ConsumeResult(resultChan)
}

// Two chan result
func TestConsumeResultErr(t *testing.T) {
	dir := filepath.Join(TEST_DATA_DIR_SRC, "rundowns_additional")
	results := ProduceResultErr(dir)
	ConsumeResultErr(results)
	results.WG.Wait()
}
