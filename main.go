package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Task struct{
	Description string
	Completed bool
}
// A constant string to hold the file name
const todofile = "tasks.txt"

// Function to display all the tasks in the To-Do list.
func DisplayTasks(tasks []Task) {
	if(len(tasks)==0){
		fmt.Println("Your To Do list is empty")
		return
	}
	for i,task := range tasks{
      prefix := "[]"
	  if task.Completed{
		prefix = "[x]"
	  }
	  fmt.Printf("%d. %s %s",i+1,prefix,task.Description)
	}
}

// A function to read the tasks from the tasks.txt file
func ReadTasks(filename string) ([]Task,error){
	tasks :=[]Task{}
   file,err := os.Open(filename)
   if(err!=nil){
	fmt.Println("Error Occured")
	return nil,err
   }
   defer file.Close()
   scanner := bufio.NewScanner(file)
   for scanner.Scan(){
	line := scanner.Text()
	if line == "" {
		continue;
	}
	parts := strings.Split(line,";")
	if len(parts) !=2 {
		fmt.Println("File structure incompatible")
		return nil,errors.New("invalid format task")
	}
	completed,_:= strconv.ParseBool(parts[1])
	tasks = append(tasks, Task{Description: parts[0],Completed: completed})
   }
   if err := scanner.Err();err!=nil{
	return nil,err
   }
  return tasks,nil
}

// A function to add a new task to the TO-DO list
func addTasks(tasks []Task,description string)[]Task{
  newTask :=  Task{Description:description,Completed:false}
  return append(tasks,newTask)
}
// A function to mark the completion of a task
func completeTask(tasks []Task,index int)([]Task,error){
 if index<1 || index >len(tasks){
	return tasks, errors.New("invallid task index")
 }
 tasks[index-1].Completed = true
 return tasks,nil
}
func removeTask(tasks []Task,index int) ([]Task,error){
	if index<1 || index >len(tasks) {
		return nil ,errors.New("invalid task index")
	}
	return append(tasks[:index-1],tasks[index:]...),nil
}
// A function to save all the tasks in TO_DO list into the file.
func saveTasks(tasks []Task,filename string) error{
	file,err:=os.OpenFile(filename,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
	if err!=nil{
		return errors.New("problem with the file")
	}
	defer file.Close()
	for _,task := range tasks{
       completed := "false"
	   if task.Completed{
		completed = "true"
	   }
	   _,err = fmt.Fprintf(file,"%s;%s\n",task.Description,completed)
	   if(err!=nil){
		return err
	   }
	}
	return nil
}
func main(){
tasks,err:= ReadTasks(todofile)
if err!=nil{
  fmt.Println("Error loading tasks")
  return
}
scanner:=bufio.NewScanner(os.Stdin)
for {
	DisplayTasks(tasks)
    fmt.Println("\nChoose an action:")
    fmt.Println("1. Add task")
    fmt.Println("2. Complete task")
    fmt.Println("3. Remove task")
    fmt.Println("4. Exit")
	fmt.Print(">")
	scanner.Scan()
	choice:=scanner.Text()
	switch choice {
	case "1": fmt.Println("Enter the task description:")
	scanner.Scan()
	description:= scanner.Text()
	tasks = addTasks(tasks,description)
	case "2":
		fmt.Println("Enter task number to complete")
		scanner.Scan()
		index,err:= strconv.Atoi(scanner.Text())
		if err!=nil{
			fmt.Println("Invalid input:", err)
            continue
		}
	  tasks ,err= completeTask(tasks,index)
	  if err != nil {
		fmt.Println(err)
		}
	case "3":
		fmt.Print("Enter task number to remove: ")
        scanner.Scan()
        index, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Invalid input:", err)
            continue
        }
        tasks, err = removeTask(tasks, index)
        if err != nil {
           fmt.Println(err)
        }
	case "4":
		fmt.Println("Exiting...")
        // Save tasks to file
        if err := saveTasks(tasks, "tasks.txt"); err != nil {
           fmt.Println("Error saving tasks:", err)
        }
        return
    default:
        fmt.Println("Invalid choice.")
	}
}
}