package action

import (
	"fmt"
	"strings"

	"github.com/ekara-platform/engine/ansible"
	"github.com/ekara-platform/engine/util"
	"github.com/ekara-platform/model"
)

type (
	// Hook represents tasks to be executed linked to an ekara life cycle event
	hookContext struct {
		action    string
		target    model.Describable
		hookOnwer string
		hookName  string
		baseParam ansible.BaseParam
		envVar    ansible.EnvVars
		buffer    ansible.Buffer
	}
)

//RunHookBefore Runs the hooks defined to be executed before a task.
func runHookBefore(rC *runtimeContext, r *StepResults, h model.Hook, ctx hookContext, cl Cleanup) StepResult {
	return runHookTasks(rC, r, h.Before, ctx, cl)
}

//RunHookAfter Runs the hooks defined to be executed after a task.
func runHookAfter(rC *runtimeContext, r *StepResults, h model.Hook, ctx hookContext, cl Cleanup) StepResult {
	return runHookTasks(rC, r, h.After, ctx, cl)
}

func runHookTasks(rC *runtimeContext, r *StepResults, tasks []model.TaskRef, ctx hookContext, cl Cleanup) StepResult {
	for i, hook := range tasks {
		repName := fmt.Sprintf("%s_%s_hook_%s_%s_%s_%d", ctx.action, ctx.target.DescName(), ctx.hookOnwer, ctx.hookName, hook.HookLocation, i)
		stepResult := InitHookStepResult(fmt.Sprintf("Running %s tasks", hook.HookLocation), ctx.target, cl)

		ef, ko := createChildExchangeFolder(rC.lC.Ef().Input, repName, &stepResult)
		if ko {
			return stepResult
		}

		t, err := hook.Resolve()
		if err != nil {
			FailsOnCode(&stepResult, err, fmt.Sprintf("An error occurred resolving the task"), nil)
			return stepResult
		}

		bp := ctx.baseParam.Copy()
		bp.AddNamedMap("params", t.Parameters)

		if ko := saveBaseParams(bp, ef.Input, &stepResult); ko {
			return stepResult
		}

		exv := ansible.BuildExtraVars("", ef.Input, ef.Output, ctx.buffer)
		r.Add(runTask(rC, t, exv, ctx.envVar))
		r.Add(consumeHookResult(rC, ctx.target, repName, ef))
	}
}

func consumeHookResult(rC *runtimeContext, target model.Describable, repName string, ef util.ExchangeFolder) StepResult {
	stepResult := InitCodeStepResult("Consuming the hook result", target, NoCleanUpRequired)
	efOut := ef.Output
	buffer, err := ansible.GetBuffer(efOut, rC.lC.Log(), "hook:"+repName)
	if err != nil {
		FailsOnCode(&stepResult, err, fmt.Sprintf("An error occurred getting the buffer"), nil)
		return stepResult
	}
	// Keep a reference on the buffer based on the output folder
	rC.buffer[efOut.Path()] = buffer
	return stepResult
}

func runTask(rC *runtimeContext, task model.Task, exv ansible.ExtraVars, env ansible.EnvVars) StepResult {
	stepResult := InitCodeStepResult("Running task", task, NoCleanUpRequired)

	// Update the runtime info
	rC.cM.TemplateContext().RunTimeInfo.SetTarget(task)
	defer rC.cM.TemplateContext().RunTimeInfo.Clear()

	// Make the task usable
	usable, err := rC.cM.Use(task)
	if err != nil {
		FailsOnCode(&stepResult, err, "An error occurred getting the usable task", nil)
		return stepResult
	}
	defer usable.Release()

	// Execute the playbook
	code, err := rC.aM.Execute(usable, task.Playbook, exv, env)
	if err != nil {
		pfd := playBookFailureDetail{
			Playbook:  task.Playbook,
			Component: task.ComponentName(),
			Code:      code,
		}
		FailsOnPlaybook(&stepResult, err, "An error occurred executing the playbook", pfd)
		return stepResult
	}

	return stepResult
}

func folderAsMessage(s string) string {
	r := strings.Replace(s, "_", " ", -1)
	return strings.Title(r[:1]) + r[1:]
}
