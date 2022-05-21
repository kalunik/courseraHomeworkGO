package main

import (
	"fmt"
	"sync"
)

func ExecutePipeline(FlowJobs ...job) {

	var in, out chan interface{}
	var wg sync.WaitGroup

	for _, j := range FlowJobs {
		in = out
		out = make(chan interface{}, 1)
		//var out chan interface{}
		//fmt.Println(cap(out), "fds")
		//fmt.Println(k, "||", j)
		wg.Add(1)
		go func(j job, in, out chan interface{}) {
			defer wg.Done()
			//defer close(in)
			defer close(out)
			j(in, out)
		}(j, in, out)
		wg.Wait()
	}
}

func main() {
	//FlowJobs := []job{
	//	job(func(in, out chan interface{}) {
	//		//out <- 2
	//		fmt.Println("in: " /*<-in,*/, " | 1 func executed")
	//	}),
	//	job(func(in, out chan interface{}) {
	//		//out <- 3
	//		fmt.Println("2 func executed. ", "in: " /*, <-in*/)
	//	}),
	//}
	//ExecutePipeline(FlowJobs...)

	FreeFlowJobs := []job{
		job(func(in, out chan interface{}) {
			fmt.Println("first ", "cap", cap(out), "ptr", out)
			out <- "Hello"
		}),
		job(func(in, out chan interface{}) {
			fmt.Println("second", "cap", cap(out), "ptr", out)
			input := fmt.Sprintf("%v", <-in)
			input += " world"
			out <- input
		}),
		job(func(in, out chan interface{}) {
			fmt.Println("third ", "cap", cap(out), "ptr", out)
			fmt.Println("Full string:", <-in)
		}),
	}
	ExecutePipeline(FreeFlowJobs...)
}
