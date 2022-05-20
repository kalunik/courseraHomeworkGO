package main

import (
	"fmt"
	"sync"
)

func ExecutePipeline(FlowJobs ...job) {

	var in, out chan interface{}
	var wg sync.WaitGroup

	for k, j := range FlowJobs {
		in = out
		out = make(chan interface{})
		fmt.Println(k, "||", j)
		wg.Add(1)
		go func(j job, in, out chan interface{}) {
			j(in, out)
			defer wg.Done() //without
			defer close(out)
		}(j, in, out)
		wg.Wait()
	}
}

func main() {
	FlowJobs := []job{
		job(func(in, out chan interface{}) {
			//out <- 2
			fmt.Println("in: " /*<-in,*/, " | 1 func executed")
		}),
		job(func(in, out chan interface{}) {
			//out <- 3
			fmt.Println("2 func executed. ", "in: " /*, <-in*/)
		}),
	}
	ExecutePipeline(FlowJobs...)
}
