package signal_test

import (
	. "github.com/hiromaily/golibs/signal"
	//lg "github.com/hiromaily/golibs/log"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
	"testing"
	"time"

	tu "github.com/hiromaily/golibs/testutil"
)

const (
	TimeOut = 2
)

var fileName string

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[Signal]")
	fileName = fmt.Sprintf("/tmp/childprocess-%d", time.Now().Unix())

	//build test/childprocess.go
	curPath, _ := os.Getwd()
	codePath := fmt.Sprintf("%s/test/childprocess.go", curPath)
	err := exec.Command("go", []string{"build", "-o", fileName, codePath}...).Run()

	if err != nil {
		panic(err)
	}
}

func teardown() {
	//remove binary of childprocess
	err := exec.Command("rm", []string{"-f", fileName}...).Run()
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestSignal(t *testing.T) {
	var procAttr os.ProcAttr
	procAttr.Files = []*os.File{nil, os.Stdout, os.Stderr}

	cases := []syscall.Signal{
		syscall.SIGINT, //os.Interrupt
		syscall.SIGTERM,
	}

	for _, sig := range cases {
		//`go test` command should not be stopped by interrupt.
		//That's why it run child process as receiver of signal and then child procss try to be done by EnableHandling()
		p, err := os.StartProcess(fileName, []string{fileName, "-time", strconv.Itoa(TimeOut)}, &procAttr)
		if err != nil {
			t.Fatalf("error occuerred error in os.StartProcess(: %v", err.Error())
			return
		}

		time.Sleep(1 * time.Second)
		err = p.Signal(sig)

		if err != nil {
			t.Fatalf("Signal(os.Interrupt) occuerred error: %v", err.Error())
			return
		}

		tm := time.Now()
		state, _ := p.Wait()
		elapsed := time.Since(tm)
		fmt.Printf("elapsed: %v", elapsed)

		//Child process should be exited.
		if !state.Exited() {
			t.Error("process should be done by Interrupt")
		}
		//Child process should be run until timeout is gone.
		//if elapsed.Truncate(1*time.Second) != TimeOut*time.Second {
		//	t.Errorf("process should be done by %d(s) but it tooks %d(s)", TimeOut, elapsed)
		//}
	}
}

func TestSignalManually(t *testing.T) {
	tu.SkipLog(t)
	wg := &sync.WaitGroup{}

	go StartSignal()
	fmt.Println("Input Ctrl + c")

	wg.Add(1)
	go func() {
		count := 0
		for {
			fmt.Println(count)
			count++
		}
	}()

	wg.Wait()
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkSignal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
