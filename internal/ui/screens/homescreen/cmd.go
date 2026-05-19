package homescreen

import (
	"time"

	types "github.com/alex-305/tuido/internal/types"
	tea "github.com/charmbracelet/bubbletea"
)

func (h *HomeScreen) deleteTaskCmd(task *types.Task) tea.Cmd {
	return func() tea.Msg {
		// TODO: delete task
		return ActionCompletedMsg{nil}
	}
}

func (h *HomeScreen) completeTaskCmd(task *types.Task) tea.Cmd {
	return func() tea.Msg {
		// TODO: complete task
		return ActionCompletedMsg{nil}
	}
}

func (h *HomeScreen) decompleteTaskCmd(task *types.Task) tea.Cmd {
	return func() tea.Msg {
		// TODO: decompleteTask
		return ActionCompletedMsg{nil}
	}
}

func (h HomeScreen) fetchCompletedTasksCmd(projectIDs []string) tea.Cmd {
	return func() tea.Msg {
		//now := time.Now()
		// tasks, err := h.ctx.APIClient.ListCompletedTasks(projectIDs, types.TickTickTime(now.AddDate(0, 0, -4000)), types.TickTickTime(now))
		//if err != nil {
		//	return CompletedTaskListMsg{tasks: tasks, err: err}
		//}
		return CompletedTaskListMsg{tasks: []*types.Task{}, err: nil}
	}
}

func (h *HomeScreen) fetchActiveTasksCmd(projectID string) tea.Cmd {
	return func() tea.Msg {
		//tasks, err := h.ctx.APIClient.ListTasks(projectID)
		//if err != nil {
		//return ActiveTaskListMsg{tasks: tasks, err: err}
		//}
		return ActiveTaskListMsg{tasks: []*types.Task{}, err: nil}
	}
}

func (h *HomeScreen) fetchProjectsCmd() tea.Cmd {
	return func() tea.Msg {
		//projects, err := h.ctx.APIClient.ListProjects()
		//if err != nil {
		//return ProjectsLoadedMsg{projects: projects, err: err}
		//}
		return ProjectsLoadedMsg{projects: []*types.Project{}, err: nil}
	}
}

func (h *HomeScreen) showLoadingCmd() tea.Cmd {
	return tea.Tick(time.Second, func(_ time.Time) tea.Msg {
		return ShowLoadingMsg{}
	})
}
