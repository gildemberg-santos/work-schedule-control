package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gildemberg-santos/work-schedule-control/pkg"
)

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	tasklist := pkg.TaskList{}

	task := pkg.Task{}
	fmt.Print("Digite o título da tarefa: ")
	title, _ := consoleReader.ReadString('\n')

	fmt.Print("Digite a descrição da tarefa: ")
	description, _ := consoleReader.ReadString('\n')

	fmt.Print("Prioridade da tarefa: ")
	priority, _ := consoleReader.ReadString('\n')

	task.AddTask(title, description, time.Now(), priority)
	tasklist.Tasks = append(tasklist.Tasks, task)

	fmt.Println(tasklist)
}
