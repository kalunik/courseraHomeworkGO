package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

const hashVariance = 6

func ExecutePipeline(FlowJobs ...job) {
	in := make(chan interface{})
	var wg sync.WaitGroup

	for _, j := range FlowJobs {
		wg.Add(1)
		out := make(chan interface{})

		go func(j job, in chan interface{}) {
			defer wg.Done()
			defer close(out)
			j(in, out)
		}(j, in)

		in = out
	}
	wg.Wait()
}

func workerCrc32(data string) chan string {
	crc := make(chan string)
	go func() {
		defer close(crc)
		crc <- DataSignerCrc32(data)
	}()
	return crc
}

func printerSingleHash(data string, Md string, crc string, crcMd string) {
	fmt.Println(data, "SingleHash data", data)
	fmt.Println(data, "SingleHash md5(data)", Md)
	fmt.Println(data, "SingleHash crc32(md5(data))", crcMd)
	fmt.Println(data, "SingleHash crc32(data)", crc)
	fmt.Printf("%s SingleHash result %s~%s\n", data, crc, crcMd)
}

var SingleHash = func(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for v := range in {
		data := fmt.Sprint(v)
		Md := DataSignerMd5(data) //it'll overheat if put in `go`
		wg.Add(1)
		go func(data string, Md string) {
			defer wg.Done()
			crc := workerCrc32(data)
			crcMd := workerCrc32(Md)

			crcPrint := <-crc
			crcMdPrint := <-crcMd
			//printerSingleHash(data, Md, crcPrint, crcMdPrint)

			out <- fmt.Sprintf("%s~%s", crcPrint, crcMdPrint)
		}(data, Md)
	}
	wg.Wait()
}

var MultiHash = func(in, out chan interface{}) {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		res []string
	)
	for v := range in {
		wg.Add(1)

		go func(data interface{}) {
			defer wg.Done()

			crcBuf := make([]chan string, 0, hashVariance)
			for th := 0; th < hashVariance; th++ {
				crc := workerCrc32(fmt.Sprintf("%d%s", th, data))
				crcBuf = append(crcBuf, crc)
			}

			manyHashes := make([]string, 0, hashVariance)
			for _, ch := range crcBuf { // change '_' to 'th', if u want print
				manyHashes = append(manyHashes, <-ch)
				//fmt.Println(data.(string), "MultiHash: crc32(th+step1)) ", th, manyHashes[th])
			}
			mu.Lock()
			res = append(res, strings.Join(manyHashes, ""))
			mu.Unlock()

			//fmt.Printf("%s MultiHash result: %s\n\n", data, strings.Join(manyHashes, ""))
		}(v)
	}
	wg.Wait()
	out <- res
}

var CombineResults = func(in, out chan interface{}) {
	v := <-in
	data := v.([]string)

	sort.Strings(data)
	//fmt.Println("CombineResults ", strings.Join(data, "_"))
	out <- strings.Join(data, "_")
}

func main() {
	testExpected := "29568666068035183841425683795340791879727309630931025356555_4958044192186797981418233587017209679042592862002427381542"
	testResult := "NOT_SET"

	inputData := []int{0, 1}

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
			fmt.Println("\033[1;34m", "Result", "\033[0m", dataRaw)
			data, ok := dataRaw.(string)
			if !ok {
				fmt.Printf("cant convert result data to string")
			}
			testResult = data
		}),
	}
	start := time.Now()

	ExecutePipeline(hashSignJobs...)

	end := time.Since(start)
	fmt.Println("\033[1;34m", "Time", "\033[0m", end)

	if testExpected != testResult {
		fmt.Printf("results not match\nGot: %v\nExpected: %v", testResult, testExpected)
	}
}
