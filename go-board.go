package main

import (
    //"github.com/nsf/termbox-go"
    //"time"
    "fmt"
    "os"

)

var debugEnabled bool = false;
var version string = "0.1";

type Task struct {
    ID   int
    Text string
}

type KanbanBoard struct {
    ToDo        []Task
    InProgress  []Task
    Review      []Task
    Done        []Task
}


func ShowUsage() {
	fmt.Println("\n")
	fmt.Println("go-board version: " + version + " - Application prints kanban board on the terminal.")

	fmt.Println("Usage:")
    fmt.Println("    go-board.exe --configure KanbanBoard <path to the Kanban board file>  # creates config files ")
	fmt.Println("    go-board.exe --new <ID> <TASK description but not too long>  # adds a new task to the ToDo lane")
    fmt.Println("    go-board.exe --inprogress <ID>  # moves the task with id to the InProgress lane")
    fmt.Println("    go-board.exe --inreview <ID>    # moves the task with id to the InReview lane")
    fmt.Println("    go-board.exe --onhold <ID>      # moves the task with id to the OnHole lane")
    fmt.Println("    go-board.exe --done <ID>        # moves the task with id to the Done lane")
    fmt.Println("    go-board.exe --archive <ID>     # moves the task with id to archived board")
    fmt.Println("    go-board.exe --help             # shows this help")

	fmt.Println("\nAll aguments you provided: ")
	fmt.Println(os.Args)

}


func main() {
    

    // // Initial data
    // board := KanbanBoard{
    //     ToDo: []Task{
    //         {ID: 1, Text: "Design UI"},
    //         {ID: 2, Text: "Write Tests"},
    //     },
    //     InProgress: []Task{
    //         {ID: 3, Text: "Implement API"},
    //     },
    //     Review: []Task{
    //         {ID: 4, Text: "Fix Bugs"},
    //     },
    //     Done: []Task{
    //         {ID: 5, Text: "Setup Server"},
    //     },
    // }

    fmt.Println("Go-board - Version: " + version)

	argsWithProgramName := os.Args
	argsWithoutProgramName := os.Args[1:]
	fmt.Println("Application and arguments: ",argsWithProgramName)
	fmt.Println("only arguments: ",argsWithoutProgramName)
	fmt.Println("\n");
	fmt.Println("-------------------------------------------------------");
	if len(os.Args) == 3 || len(os.Args) == 4 { 
		// 2 or 3  arguments
		
		operation := os.Args[1];
		taskId := os.Args[2];
		taskDescription := os.Args[3];
		
        fmt.Println("Arguments OK: ",argsWithoutProgramName)

        fmt.Println("Arguments OK: ",operation, taskId, taskDescription)
        
    }



}

// func drawBoard(board *KanbanBoard) {
//     const colWidth = 30
//     const rowHeight = 2
//     x, y := 2, 1

//     // Draw column headers
//     drawText(x, y, "To Do", termbox.ColorWhite, termbox.ColorBlack)
//     drawText(x+colWidth, y, "In Progress", termbox.ColorWhite, termbox.ColorBlack)
//     drawText(x+2*colWidth, y, "Review", termbox.ColorWhite, termbox.ColorBlack)
//     drawText(x+3*colWidth, y, "Done", termbox.ColorWhite, termbox.ColorBlack)

//     y += 2 // Move down for tasks

//     // Draw tasks for each column
//     drawTasks(x, y, board.ToDo, colWidth, rowHeight)
//     drawTasks(x+colWidth, y, board.InProgress, colWidth, rowHeight)
//     drawTasks(x+2*colWidth, y, board.Review, colWidth, rowHeight)
//     drawTasks(x+3*colWidth, y, board.Done, colWidth, rowHeight)
// }

// func drawText(x, y int, text string, fg, bg termbox.Attribute) {
//     for i, ch := range text {
//         termbox.SetCell(x+i, y, ch, fg, bg)
//     }
// }

// func drawTasks(x, y int, tasks []Task, colWidth, rowHeight int) {
//     for i, task := range tasks {
//         if i >= 10 { // Limit to 10 tasks per column for simplicity
//             break
//         }
//         text := task.Text
//         if len(text) > colWidth-2 {
//             text = text[:colWidth-3] + "..."
//         }
//         // Draw border
//         for j := 0; j < rowHeight; j++ {
//             for k := 0; k < colWidth; k++ {
//                 if j == 0 || j == rowHeight-1 || k == 0 || k == colWidth-1 {
//                     termbox.SetCell(x+k, y+j, ' ', termbox.ColorWhite, termbox.ColorBlack)
//                 }
//             }
//         }
//         // Draw task text
//         drawText(x+1, y+1, text, termbox.ColorBlack, termbox.ColorWhite)
//         y += rowHeight + 1 // Space between tasks
//     }
// }


