package TaskHandler

type Task struct {
	Name       string
	Status     string
	TaskNumber int
}

type TaskServiceInterface interface {
	Update(*Task) error
	Create(*Task) error
	TaskList() ([]*Task, error)
}
