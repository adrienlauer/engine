package ansible

import (
	"bufio"
	"github.com/lagoon-platform/engine/component"
	"github.com/lagoon-platform/engine/util"
	"github.com/lagoon-platform/model"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

type AnsibleManager interface {
	Execute(component model.Component, playbook string, extraVars ExtraVars, envars EnvVars) int
}

type ExecutionReport struct {
}

type context struct {
	logger           *log.Logger
	componentManager component.ComponentManager
}

func CreateAnsibleManager(logger *log.Logger, componentManager component.ComponentManager) AnsibleManager {
	return &context{
		logger:           logger,
		componentManager: componentManager}
}

func (ctx context) Execute(component model.Component, playbook string, extraVars ExtraVars, envars EnvVars) int {
	// Path of the component where the playbook is supposed to be located
	path := ctx.componentManager.ComponentPath(component.Id)

	playBookPath := filepath.Join(path, playbook)
	if _, err := os.Stat(playBookPath); os.IsNotExist(err) {
		ctx.logger.Fatalf(util.ERROR_MISSING, playBookPath)
	} else {
		ctx.logger.Printf(util.LOG_STARTING_PLAYBOOK, playBookPath)
	}
	ctx.logger.Printf(util.LOG_LAUNCHING_PLAYBOOK, playBookPath)

	var args = []string{playbook}
	moduleDirectories := ctx.componentManager.MatchingDirectories(util.ComponentModuleFolder)
	if len(moduleDirectories) > 0 {
		ctx.logger.Printf("Detected %d modules directories for launch: %s", len(moduleDirectories), moduleDirectories)
		args = append(args, "--module-path", strings.Join(moduleDirectories, ":"))
	} else {
		ctx.logger.Printf("No module directory detected for launch")
	}

	inventoryDirectories := ctx.componentManager.MatchingDirectories(util.InventoryModuleFolder)
	if len(inventoryDirectories) > 0 {
		ctx.logger.Printf("Detected %d inventory directories for launch: %s", len(moduleDirectories), moduleDirectories)
		args = append(args, "--inventory", strings.Join(inventoryDirectories, ":"))
	} else {
		ctx.logger.Printf("No inventory directory detected for launch")
	}

	if extraVars.Bool {
		ctx.logger.Printf("Playbook extra-vars: %s", extraVars.String())
		args = append(args, "--extra-vars", extraVars.String())
	} else {
		ctx.logger.Printf("No extra-vars")
	}

	cmd := exec.Command("ansible-playbook", args...)
	cmd.Dir = path
	cmd.Env = os.Environ()
	for k, v := range envars.Content {
		cmd.Env = append(cmd.Env, k+"="+v)
	}

	errReader, err := cmd.StderrPipe()
	if err != nil {
		ctx.logger.Fatal(err)
	}
	logPipe(errReader, ctx.logger)

	outReader, err := cmd.StdoutPipe()
	if err != nil {
		ctx.logger.Fatal(err)
	}
	logPipe(outReader, ctx.logger)

	err = cmd.Start()
	if err != nil {
		ctx.logger.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		e, ok := err.(*exec.ExitError)
		if ok {
			s := e.Sys().(syscall.WaitStatus)
			code := s.ExitStatus()
			ctx.logger.Printf("---> ANSIBLE RETURNED ERROR : %v\n", ReturnedError(code))
			return code
			// TODO write report here
		} else {
			ctx.logger.Fatal(err)
		}
	}
	return 0
}

// logPipe logs the given pipe, reader/closer on the given logger
func logPipe(rc io.ReadCloser, l *log.Logger) {
	s := bufio.NewScanner(rc)
	go func() {
		for s.Scan() {
			l.Printf("%s\n", s.Text())
		}
	}()
}
