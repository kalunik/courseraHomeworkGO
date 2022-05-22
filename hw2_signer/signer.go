package main

import (
	"fmt"
	"sync"
	"time"
)

func ExecutePipeline(FlowJobs ...job) {
	in := make(chan interface{})
	var wg sync.WaitGroup

	for _, j := range FlowJobs {
		wg.Add(1)
		out := make(chan interface{})

		go func(j job, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			j(in, out)
		}(j, in, out)

		in = out
	}
	wg.Wait()
}

var SingleHash = func(in, out chan interface{}) {
	dataRaw := <-in
	data := dataRaw.(string)
	crc := DataSignerCrc32(data)
	crcMd := DataSignerCrc32(DataSignerMd5(data))
	fmt.Println()

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

	/*	FreeFlowJobs := []job{
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
		ExecutePipeline(FreeFlowJobs...)*/
	inputData := []int{0, 1}

	hashSignJobs := []job{
		job(func(in, out chan interface{}) {
			for _, fibNum := range inputData {
				out <- fibNum
			}
		}),
		job(SingleHash),
		/*job(MultiHash),
		job(CombineResults),*/
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			data, ok := dataRaw.(string)
			if !ok {
				t.Error("cant convert result data to string")
			}
			testResult = data
		}),
	}

	start := time.Now()

	ExecutePipeline(hashSignJobs...)
}
