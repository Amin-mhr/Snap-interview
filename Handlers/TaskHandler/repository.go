package TaskHandler

import "errors"

type TaskRepositoryInterface interface {
	UpdateTask(task *Task) error
	CreateTask(task *Task) error
	GetTasks() ([]*Task, error)
}

type TaskRepository struct {
	tasks []*Task
}

func NewTaskRepository() TaskRepositoryInterface {
	repo := &TaskRepository{tasks: make([]*Task, 0)}
	for i := 0; i < 10; i++ {
		repo.tasks = append(repo.tasks, &Task{
			Name:       "task",
			Status:     "running",
			TaskNumber: i,
		})
	}
	return repo
}

func (r *TaskRepository) UpdateTask(task *Task) error {
	for _, t := range r.tasks {
		if t.TaskNumber == task.TaskNumber {
			t.Status = task.Status
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *TaskRepository) CreateTask(task *Task) error {
	task.TaskNumber = len(r.tasks)
	r.tasks = append(r.tasks, task)
	return nil
}

func (r *TaskRepository) GetTasks() ([]*Task, error) {
	return r.tasks, nil
}
