package main

import (
    "github.com/nsf/termbox-go"
     //"time"
)

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


func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    // Initial data
    board := KanbanBoard{
        ToDo: []Task{
            {ID: 1, Text: "Design UI"},
            {ID: 2, Text: "Write Tests"},
        },
        InProgress: []Task{
            {ID: 3, Text: "Implement API"},
        },
        Review: []Task{
            {ID: 4, Text: "Fix Bugs"},
        },
        Done: []Task{
            {ID: 5, Text: "Setup Server"},
        },
    }

    // Main loop
loop:
    for {
        termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
        drawBoard(&board)
        termbox.Flush()

        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            switch ev.Key {
            case termbox.KeyCtrlC:
                break loop
            case termbox.KeyArrowLeft, termbox.KeyArrowRight, termbox.KeyArrowUp, termbox.KeyArrowDown:
                // Here you can add logic to navigate or move tasks
            }
        case termbox.EventError:
            panic(ev.Err)
        }
    }
}

func drawBoard(board *KanbanBoard) {
    const colWidth = 30
    const rowHeight = 2
    x, y := 2, 1

    // Draw column headers
    drawText(x, y, "To Do", termbox.ColorWhite, termbox.ColorBlack)
    drawText(x+colWidth, y, "In Progress", termbox.ColorWhite, termbox.ColorBlack)
    drawText(x+2*colWidth, y, "Review", termbox.ColorWhite, termbox.ColorBlack)
    drawText(x+3*colWidth, y, "Done", termbox.ColorWhite, termbox.ColorBlack)

    y += 2 // Move down for tasks

    // Draw tasks for each column
    drawTasks(x, y, board.ToDo, colWidth, rowHeight)
    drawTasks(x+colWidth, y, board.InProgress, colWidth, rowHeight)
    drawTasks(x+2*colWidth, y, board.Review, colWidth, rowHeight)
    drawTasks(x+3*colWidth, y, board.Done, colWidth, rowHeight)
}

func drawText(x, y int, text string, fg, bg termbox.Attribute) {
    for i, ch := range text {
        termbox.SetCell(x+i, y, ch, fg, bg)
    }
}

func drawTasks(x, y int, tasks []Task, colWidth, rowHeight int) {
    for i, task := range tasks {
        if i >= 10 { // Limit to 10 tasks per column for simplicity
            break
        }
        text := task.Text
        if len(text) > colWidth-2 {
            text = text[:colWidth-3] + "..."
        }
        // Draw border
        for j := 0; j < rowHeight; j++ {
            for k := 0; k < colWidth; k++ {
                if j == 0 || j == rowHeight-1 || k == 0 || k == colWidth-1 {
                    termbox.SetCell(x+k, y+j, ' ', termbox.ColorWhite, termbox.ColorBlack)
                }
            }
        }
        // Draw task text
        drawText(x+1, y+1, text, termbox.ColorBlack, termbox.ColorWhite)
        y += rowHeight + 1 // Space between tasks
    }
}


