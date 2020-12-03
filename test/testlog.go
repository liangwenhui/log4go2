package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	testLog()

}

func testLog() {
	iLoopTime := flag.Int("loop", 10, "loop")
	flag.Parse()
	var iLoopTimes = *iLoopTime
	var test_times = 500
	pid := os.Getpid()
	fmt.Printf("pid=%d Times=%dx\n", pid, iLoopTimes)
	start := time.Now()
	for j := 0; j < iLoopTimes; j++ {
		for x := 0; x < 20; x++ {
			for i := j; i < test_times; i++ {
				//fmt.Fprintf(writer,AiRandomFmt[i], AiRandomInt[i], AiRandomInt[test_times+i], AsRandomStr[i], AfRandomFloat[i])
				fmt.Println(AiRandomFmt[i])
			}
		}
	}
	since := time.Since(start)
	fmt.Println("循环耗时：", since)
	fmt.Printf("Press any key to exit...\n")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
