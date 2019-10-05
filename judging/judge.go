package judging

import (
	pb "github.com/BayviewComputerClub/smoothie-runner/protocol"
	"github.com/BayviewComputerClub/smoothie-runner/shared"
	"github.com/BayviewComputerClub/smoothie-runner/util"
	"golang.org/x/sys/unix"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type CaseReturn struct {
	Result     string
	ResultInfo string
}

/*
 * Tasks:
 * setrlimit to prevent fork() in processes (use golang.org/x/sys)
 * set nice value to be low
 * run process as noperm user (or just run the whole program with no perm)
 * use ptrace unix calls
 * can use syscall credentials
 * make sure thread is locked
 */


func judgeStderrListener(reader *io.ReadCloser, done chan CaseReturn) {
	str, err := ioutil.ReadAll(*reader)

	if err != nil { // should terminate peacefully
		util.Warn("Stderr: " + err.Error()) // TODO
	} else {
		done <- CaseReturn{
			Result:     shared.OUTCOME_RTE,
			ResultInfo: string(str),
		}
	}
}

func judgeStdinFeeder(writer *os.File, done chan CaseReturn, feed *string) {
	//println(*feed)
	_, err := writer.WriteString(*feed)
	if err != nil {
		done <- CaseReturn{
			Result: shared.OUTCOME_RTE,
		}
		util.Warn("Stdin: " + err.Error()) // TODO
		return
	}
}

func judgeCheckTimeout(c *exec.Cmd, d time.Duration, done chan CaseReturn) {
	time.Sleep(d)
	if util.IsPidRunning(c.Process.Pid) {
		done <- CaseReturn {Result: shared.OUTCOME_TLE}
	}
}

func judgeCase(c *exec.Cmd, batchCase *pb.ProblemBatchCase, result chan pb.TestCaseResult) {
	runtime.LockOSThread() // https://github.com/golang/go/issues/7699
	defer runtime.UnlockOSThread()

	done := make(chan CaseReturn)

	t := time.Now()

	// enable ptrace
	c.SysProcAttr = &unix.SysProcAttr{Ptrace: true}

	// initialize pipes

	/*stderrPipe, err := c.StderrPipe()
	if err != nil {
		util.Warn("stderrpipeinit: " + err.Error())
		result <- pb.TestCaseResult{Result: shared.OUTCOME_ISE, ResultInfo: err.Error()}
		return
	}*/

	inputBuff, inStream, err := os.Pipe()
	if err != nil {
		panic(err) // TODO
	}
	c.Stdin = inStream
	defer inputBuff.Close()
	defer inStream.Close()

	outputBuff, outStream, err := os.Pipe()
	if err != nil {
		panic(err) // TODO
	}
	c.Stdout = outStream
	defer outputBuff.Close()
	defer outStream.Close()

	// start process
	err = c.Start()
	if err != nil {
		util.Warn("RTE: " + err.Error())
		result <- pb.TestCaseResult{Result: shared.OUTCOME_ISE, ResultInfo: err.Error()}
		return
	}

	// start time watch (convert to seconds)
	go judgeCheckTimeout(c, time.Duration(batchCase.TimeLimit) * time.Second, done)

	c.Wait() // pause execution on first instruction

	// start listener pipes
	go judgeStdoutListener(c, outputBuff, done, &batchCase.ExpectedAnswer)
	//go judgeStderrListener(&stderrPipe, done)

	// feed input to process
	go judgeStdinFeeder(inStream, done, &batchCase.Input)

	// sandbox has to hog the thread, so move receiving to another one
	go func() {
		// wait for judging to finish
		response := <-done

		//fmt.Println(response.Result + " " + response.ResultInfo) // TODO

		if util.IsPidRunning(c.Process.Pid) {
			err = c.Process.Kill()
			if err != nil  && err.Error() != "os: process already finished" {
				util.Warn("pkill fail: " + err.Error())
			}
		}

		result <- pb.TestCaseResult{
			Result:     response.Result,
			ResultInfo: response.ResultInfo,
			Time:       time.Since(t).Seconds(),
			MemUsage:   0, // TODO
		}
	}()

	// start sandboxing
	// must run on this thread because all ptrace calls have to come from one thread
	sandboxProcess(&c.Process.Pid, done)
}
