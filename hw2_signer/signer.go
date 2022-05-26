package main

import (
	"fmt"
	"sort"
	"strings"
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
	for v := range in {
		data := fmt.Sprint(v)
		Md := DataSignerMd5(data)
		crcMd := DataSignerCrc32(Md)
		crc := DataSignerCrc32(data)
		fmt.Println(data, "SingleHash data", data)
		fmt.Println(data, "SingleHash md5(data)", Md)
		fmt.Println(data, "SingleHash crc32(md5(data))", crcMd)
		fmt.Println(data, "SingleHash crc32(data)", crc)
		fmt.Printf("%s SingleHash result %s~%s\n", data, crc, crcMd)
		out <- fmt.Sprintf("%s~%s", crc, crcMd) // out must be in range, so then PANIC happen
	}
}

var MultiHash = func(in, out chan interface{}) {
	var (
		tmp, crc string
		res      []string
	)
	for v := range in {
		data := fmt.Sprint(v)
		for th := 0; th < 6; th++ {
			crc = DataSignerCrc32(fmt.Sprintf("%d%s", th, data))
			fmt.Println(data, "MultiHash: crc32(th+step1)) ", th, crc)
			tmp += crc
		}
		res = append(res, tmp)
		fmt.Println(data, "MultiHash result:", tmp)
		tmp = ""
	}
	out <- res
}

var CombineResults = func(in, out chan interface{}) {
	v := <-in
	data := v.([]string)

	sort.Strings(data)
	out <- strings.Join(data, "_")
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

	testExpected := "27225454331033649287118297354036464389062965355426795162684_29568666068035183841425683795340791879727309630931025356555"
	testResult := "NOT_SET"

	inputData := []int{0, 2}

	hashSignJobs := []job{
		job(func(in, out chan interface{}) {
			for _, fibNum := range inputData {
				out <- fibNum
			}
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			fmt.Println("!!", dataRaw)
			data, ok := dataRaw.(string)
			if !ok {
				_ = fmt.Errorf("cant convert result data to string")
			}
			testResult = data
		}),
	}

	ExecutePipeline(hashSignJobs...)

	if testExpected != testResult {
		fmt.Printf("results not match\nGot: %v\nExpected: %v", testResult, testExpected)
	}
}
