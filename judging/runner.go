package judging

import (
	"github.com/BayviewComputerClub/smoothie-runner/adapters"
	pb "github.com/BayviewComputerClub/smoothie-runner/protocol"
	"github.com/BayviewComputerClub/smoothie-runner/shared"
	"github.com/BayviewComputerClub/smoothie-runner/util"
	"github.com/rs/xid"
	"os"
	"syscall"
	"unsafe"
)

func emptyTcr() *pb.TestCaseResult {
	return &pb.TestCaseResult{
		BatchNumber: 0,
		CaseNumber:  0,
		Result:      "",
		ResultInfo:  "",
		Time:        0,
		MemUsage:    0,
	}
}


func TestSolution(req *pb.TestSolutionRequest, res chan shared.JudgeStatus, cancelled *bool) {

	// create judgesession object
	session := shared.JudgeSession{
		Workspace:       shared.TESTING_DIR + "/" + xid.New().String(),
		Code:            req.Solution.Code,
		Language:        req.Solution.Language,
		OriginalRequest: req,
	}

	// remove workspace when exit
	defer os.RemoveAll(session.Workspace)

	// create session workspace
	err := os.Mkdir(session.Workspace, 0755)
	if err != nil {
		panic(err)
	}

	// attempt to compile user submitted code
	runCommand, err := adapters.CompileAndGetRunCommand(session)
	if err != nil {

		// send compile error back
		res <- shared.JudgeStatus{
			Err: err,
			Res: pb.TestSolutionResponse{
				TestCaseResult:   emptyTcr(),
				CompletedTesting: true,
				CompileError:     shared.OUTCOME_CE + ": " + err.Error(),
			},
		}

		return
	}

	// get exec command pointers
	f, err := os.Open(runCommand.Path)
	if err != nil {
		util.Warn("commandfileopen: " + err.Error())
		res <- shared.JudgeStatus{
			Err: err,
			Res: pb.TestSolutionResponse{
				TestCaseResult:   emptyTcr(),
				CompletedTesting: true,
				CompileError:     shared.OUTCOME_ISE,
			},
		}

		return
	}
	defer f.Close()

	commandArgs, err := syscall.SlicePtrFromStrings(append(runCommand.Args, "NULL"))
	if err != nil {
		util.Warn("commandbyteparse: " + err.Error())
		res <- shared.JudgeStatus{
			Err: err,
			Res: pb.TestSolutionResponse{
				TestCaseResult:   emptyTcr(),
				CompletedTesting: true,
				CompileError:     shared.OUTCOME_ISE,
			},
		}

		return
	}

	// loop over test batches and cases
	for _, batch := range req.Solution.Problem.TestBatches {
		for _, batchCase := range batch.Cases {
			if *cancelled { // exit if cancelled
				return
			}

			batchRes := make(chan pb.TestCaseResult)

			// do judging
			gradingSession := GradeSession{
				JudgingSession: &session,
				Problem:        req.Solution.Problem,
				Solution:       req.Solution,
				CurrentBatch:   batchCase,
				Stderr:         "",
				ExitCode:       0,
				StreamResult:   batchRes,
				StreamDone:     make(chan CaseReturn),
				Command:        runCommand,
				ExecCommand:    f.Fd(),
				ExecArgs:       uintptr(unsafe.Pointer(&commandArgs)),
			}
			go gradingSession.StartJudging()

			// wait for case result
			result := <-batchRes

			// send result
			res <- shared.JudgeStatus{
				Err: nil,
				Res: pb.TestSolutionResponse{
					TestCaseResult:   &result,
					CompletedTesting: false,
					CompileError:     "",
				},
			}

			// exit if whole batch fails
			if result.Result != shared.OUTCOME_AC && !req.TestBatchEvenIfFailed {
				break
			}
		}
	}

	// return successful judging
	res <- shared.JudgeStatus{
		Err: nil,
		Res: pb.TestSolutionResponse{
			TestCaseResult:   emptyTcr(),
			CompletedTesting: true,
			CompileError:     "",
		},
	}

}
