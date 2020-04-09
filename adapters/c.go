package adapters

import (
	"errors"
	pb "github.com/BayviewComputerClub/smoothie-runner/protocol/runner"
	"github.com/BayviewComputerClub/smoothie-runner/shared"
	"io/ioutil"
	"os/exec"
	"strings"
)

type C11Adapter struct {}

func (adapter C11Adapter) GetName() string {
	return "c11"
}

func (adapter C11Adapter) JudgeFinished(tcr *pb.TestCaseResult) {}

func (adapter C11Adapter) Compile(session *shared.JudgeSession) (*exec.Cmd, error) {
	err := ioutil.WriteFile(session.Workspace + "/main.c", []byte(session.Code), 0644)
	if err != nil {
		return nil, err
	}

	output, err := exec.Command("gcc", "-std=c11", session.Workspace+"/main.c", "-o", session.Workspace+"/main").CombinedOutput()
	if err != nil {
		return nil, errors.New(strings.ReplaceAll(string(output), session.Workspace+"/main.c", ""))
	}

	c := exec.Command("./main")
	c.Dir = session.Workspace
	return c, nil
}
