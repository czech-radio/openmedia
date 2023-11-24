package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// NON-BLOCKING RESULTS PIPING
// // ONE CHANNEL
type Result struct {
	Message string
	Error   error
}

func ProduceResult(directory_path string) chan Result {
	resultChan := make(chan Result)
	go func() {
		dirs, err := os.ReadDir(directory_path)
		if err != nil {
			result := Result{}
			result.Error = err
			resultChan <- result
			close(resultChan)
		}
		for i, dir := range dirs {
			Sleeper(133, "ms")
			result := Result{}
			if i%3 == 0 {
				result.Error = fmt.Errorf("file not readable %s", dir)
				resultChan <- result
			} else {
				result.Message = filepath.Join(directory_path, dir.Name())
				resultChan <- result
			}
		}
		close(resultChan)
	}()
	return resultChan
}

func ConsumeResult(resultChan <-chan Result) {
	for {
		Sleeper(167, "ms")
		result, ok := <-resultChan
		if !ok && len(resultChan) == 0 {
			break
		}
		if result.Error != nil {
			fmt.Println(result.Error)
		} else {
			fmt.Println(result.Message)
		}
	}
	fmt.Println("finished")
}

// // TWO CHANNELS WITH WAIT GROUP
type ResultErr struct {
	ChanMsg chan string
	ChanErr chan error
	WG      *sync.WaitGroup
}

func NewChanResults() ResultErr {
	var chr ResultErr
	chr.ChanMsg = make(chan string)
	chr.ChanErr = make(chan error)
	chr.WG = new(sync.WaitGroup)
	return chr
}

func ProduceResultErr(directory_path string) *ResultErr {
	results := NewChanResults()
	results.WG.Add(1)
	go func() {
		dirs, err := os.ReadDir(directory_path)
		if err != nil {
			results.ChanErr <- err
		}
		for i, dir := range dirs {
			if i%3 == 0 {
				results.ChanErr <- fmt.Errorf("file not readable %s", dir)
			} else {
				results.ChanMsg <- filepath.Join(directory_path, dir.Name())
			}
		}
		close(results.ChanMsg)
		results.WG.Done()
	}()
	return &results
}

func ConsumeResultErr(inputs *ResultErr) {
	inputs.WG.Add(1)
	go func() {
	consume:
		for {
			select {
			case msg, ok := <-inputs.ChanMsg:
				if !ok && len(inputs.ChanMsg) == 0 {
					fmt.Println("finished")
					inputs.WG.Done()
					break consume
				}
				fmt.Println(msg)
				Sleeper(100, "ms")
			case err := <-inputs.ChanErr:
				fmt.Println(err)
			}
		}
	}()
}
