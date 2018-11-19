package ansible

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/ekara-platform/engine/component"
	"github.com/ekara-platform/engine/util"
	"github.com/ekara-platform/model"
)

type AnsibleManager interface {
	Execute(component model.Component, playbook string, extraVars ExtraVars, envars EnvVars, inventories string) (error, int)
	Contains(component model.Component, playbook string) bool
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

func (ctx context) Contains(component model.Component, file string) bool {
	path := ctx.componentManager.ComponentPath(component.Id)
	playbookPath := util.JoinPaths(path, file)
	if _, err := os.Stat(playbookPath); !os.IsNotExist(err) {
		return true
	}
	return false
}

func (ctx context) Execute(component model.Component, playbook string, extraVars ExtraVars, envars EnvVars, inventories string) (error, int) {
	// Path of the component where the playbook is supposed to be located
	path := ctx.componentManager.ComponentPath(component.Id)

	playBookPath := filepath.Join(path, playbook)
	if _, err := os.Stat(playBookPath); os.IsNotExist(err) {
		return err, 0
	} else {
		ctx.logger.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - -")
		ctx.logger.Println("* * * * * A N S I B L E - - P L A Y B O O K  * * * * ")
		ctx.logger.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - -")
		ctx.logger.Printf(util.LOG_STARTING_PLAYBOOK, playBookPath)
		ctx.logger.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - -")
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
	if inventories != "" {
		ctx.logger.Printf("launched with inventories :%s", inventories)
		for _, v := range strings.Split(inventories, " ") {
			args = append(args, v)
		}
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
		return err, 0
	}
	logPipe(errReader, ctx.logger)

	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return err, 0
	}
	logPipe(outReader, ctx.logger)

	err = cmd.Start()
	if err != nil {
		return err, 0
	}

	err = cmd.Wait()
	if err != nil {
		e, ok := err.(*exec.ExitError)
		if ok {
			s := e.Sys().(syscall.WaitStatus)
			code := s.ExitStatus()
			ctx.logger.Printf("Ansible returned error code : %v\n", ReturnedError(code))
			return err, code
		} else {
			return err, 0
		}
	}
	return nil, 0
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