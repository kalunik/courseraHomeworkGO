package main

import (
	"fmt"
	"sync"
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
	var (
		crcMd,
		crc string
	)
	for v := range in {
		data := fmt.Sprint(v)
		Md := DataSignerMd5(data)
		crcMd := DataSignerCrc32(Md)
		crc := DataSignerCrc32(data)
		fmt.Println(data, " SingleHash data ", data)
		fmt.Println(data, " SingleHash md5(data) ", Md)
		fmt.Println(data, " SingleHash crc32(md5(data)) ", crcMd)
		fmt.Println(data, " SingleHash crc32(data) ", crc)
		fmt.Printf(" SingleHash result ", crc, "~", crcMd)
	}
	out <- fmt.Sprintf("%s~%s", crc, crcMd) // out must be in range, so then PANIC happen

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

	testExpected := "1173136728138862632818075107442090076184424490584241521304_1696913515191343735512658979631549563179965036907783101867_27225454331033649287118297354036464389062965355426795162684_29568666068035183841425683795340791879727309630931025356555_3994492081516972096677631278379039212655368881548151736_4958044192186797981418233587017209679042592862002427381542_4958044192186797981418233587017209679042592862002427381542"
	testResult := "NOT_SET"

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
				_ = fmt.Errorf("cant convert result data to string")
			}
			testResult = data
		}),
	}

	ExecutePipeline(hashSignJobs...)

	if testExpected != testResult {
		_ = fmt.Errorf("results not match\nGot: %v\nExpected: %v", testResult, testExpected)
	}
}
