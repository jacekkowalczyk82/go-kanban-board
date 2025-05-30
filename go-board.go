package main

import (
    //"github.com/nsf/termbox-go"
    //"time"
    "fmt"
    "os"
	"gopkg.in/yaml.v3"
    "time"

)

var CONST_TIME_FORMAT = "2006-01-02 15:04:05"
var CONST_DATE_FORMAT = "2006-01-02"

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

func getTimeNowString() string {
    currentTime := time.Now()

	//currentTimeUnixSeconds := currentTime.Unix()
    return currentTime.Format(CONST_TIME_FORMAT)

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
    fmt.Println("    go-board.exe --print            # prints the board")
    fmt.Println("    go-board.exe --help             # shows this help")

	fmt.Println("\nAll aguments you provided: ")
	fmt.Println(os.Args)

}

// ReadConfig reads config.yaml if it exists and returns its contents as a map
func ReadConfig(filePath string) (map[string]interface{}, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s does not exist", filePath)
	}

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// Unmarshal YAML into map
	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML: %v", err)
	}

	return config, nil
}


// WriteConfig writes a map to a YAML file
func WriteConfig(filePath string, config map[string]interface{}) error {
	// Marshal the map to YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshaling to YAML: %v", err)
	}

	// Write to file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filePath, err)
	}

	return nil
}

func testConfigFile() {
    // Example configuration data
	config := map[string]interface{}{
		"server": map[string]interface{}{
			"host": "localhost",
			"port": 8080,
		},
		"database": map[string]interface{}{
			"name": "mydb",
			"user": "admin",
		},
		"debug": true,
	}

	// Write to config.yaml
	err := WriteConfig("test-config.yaml", config)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Successfully wrote config to test-config.yaml")
}


func writeConfigFile(boardFilePath string, modificatioDateTime string) {
    // Example configuration data
	config := map[string]interface{}{
        "kanban-board-file": "",
        "archived-kanban-board-file": "_archived.json",
        "board-config-modification-time": "",
		"dummy-server": map[string]interface{}{
			"host": "localhost",
			"port": 8080,
		},
		"dummy-database": map[string]interface{}{
			"name": "mydb",
			"user": "admin",
		},
		"debug": true,
	}

    config["kanban-board-file"] = boardFilePath;
    config["archived-kanban-board-file"] = boardFilePath + "_archived.json";
    config["board-config-modification-time"] = modificatioDateTime;

	// Write to config.yaml
	err := WriteConfig("config.yaml", config)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Successfully wrote config to config.yaml")
}


func testReadConfigFile() {
    // Example configuration data
	

	// Read to config.yaml
    // var config := map[string]interface{}
	config, err := ReadConfig("test-config.yaml")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

    fmt.Println("config: ", config["server"])
    fmt.Println("config: ", config["database"])
	fmt.Println("Successfully read config from test-config.yaml")
}


func main() {
    
    testConfigFile()
    testReadConfigFile()

    writeConfigFile("my-kanban-board.json", getTimeNowString())
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
    correctParams:=false

	if len(os.Args) == 2 {
		// 1 arguments
		
		operation := os.Args[1];
		
		
        fmt.Println("Arguments OK: ",argsWithoutProgramName)

        fmt.Println("Arguments OK: ",operation )
        correctParams = true;
        
    } else if len(os.Args) == 3 { 
		//  2  arguments
		
		operation := os.Args[1];
		taskId := os.Args[2];
		
		
        fmt.Println("Arguments OK: ",argsWithoutProgramName)

        fmt.Println("Arguments OK: ",operation, taskId)

        correctParams = true;
        
    } else if len(os.Args) == 4 { 
        // 3  arguments
        
        operation := os.Args[1];
        taskId := os.Args[2];
        taskDescription := os.Args[3];
        
        fmt.Println("Arguments OK: ",argsWithoutProgramName)

        fmt.Println("Arguments OK: ",operation, taskId, taskDescription)

        correctParams = true;
            
    } 
    
    
    if correctParams {
        fmt.Println("Arguments OK: ")

        config, err := ReadConfig("config.yaml")
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }
        fmt.Println("config: ", config["server"])
        fmt.Println("config: ", config["database"])
	    fmt.Println("Successfully read config from test-config.yaml")

        // Print the config map
        fmt.Printf("Config contents: %+v\n", config)
    } else {
        ShowUsage()
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


