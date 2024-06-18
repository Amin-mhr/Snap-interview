package TaskHandler

import "errors"

type TaskService struct {
	repo TaskRepositoryInterface
}

func NewTaskService(repo TaskRepositoryInterface) TaskServiceInterface {
	return &TaskService{repo: repo}
}

func (s *TaskService) Update(task *Task) error {
	if task.Status == "running" {
		task.Status = "done"
		return s.repo.UpdateTask(task)
	}
	return errors.New("task status is invalid")
}

func (s *TaskService) Create(task *Task) error {
	task.Status = "not started"
	task.Name = "task"
	return s.repo.CreateTask(task)
}

func (s *TaskService) TaskList() ([]*Task, error) {
	return s.repo.GetTasks()
}
