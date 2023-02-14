package pkg

import "time"

type Task struct {
	ID         int
	Title      string
	Descrition string
	CreatedAt  time.Time
	DueDate    time.Time
	Status     string
	Priority   string
}

type TaskList struct {
	Tasks []Task
}

func (t *Task) AddTask(title, description string, dueDate time.Time, priority string) {
	t.ID = generateID()
	t.Title = title
	t.Descrition = description
	t.DueDate = dueDate
	t.Priority = priority
	t.CreatedAt = time.Now()
	t.Status = "pending"
}

func (t *Task) UpdateTask(title, description string, dueDate time.Time, priority string) {
	t.Title = title
	t.Descrition = description
	t.DueDate = dueDate
	t.Priority = priority
}

func (t *Task) DeleteTask() {
	t.Status = "deleted"
}

func (t *Task) CompleteTask() {
	t.Status = "completed"
}

func (t *Task) GetTask() Task {
	return *t
}

func (tl *TaskList) AddTaskToList(task Task) {
	tl.Tasks = append(tl.Tasks, task)
}

func (tl *TaskList) GetTaskList() []Task {
	return tl.Tasks
}

func (tl *TaskList) GetTaskByID(id int) Task {
	for _, task := range tl.Tasks {
		if task.ID == id {
			return task
		}
	}
	return Task{}
}

func (tl *TaskList) DeleteTaskByID(id int) {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].DeleteTask()
		}
	}
}

func (tl *TaskList) CompleteTaskByID(id int) {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].CompleteTask()
		}
	}
}

func generateID() int {
	return 1
}
