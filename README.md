# Task Tracker CLI

A simple command-line task tracker built in Go. Manage your tasks from the terminal with persistent.

---

## Features

- Add tasks with auto-generated IDs
- List all tasks or filter by status (`todo`, `in-progress`, `done`)
- Update task descriptions
- Delete tasks by ID
- Mark tasks as in-progress or done
- Persistent local storage via `tasks.json`
- Clear error messages and proper exit codes

---

## Installation

### Prerequisites

- Go 1.21 or higher

### Build from Source

```bash
git clone https://github.com/OdaloV/tasktracker.git
cd tasktracker
go build -o task-cli
```
---

## Usage

```bash
task-cli <command> [arguments]
```

### Available Commands

| Command | Description | Example |
|---|---|---|
| `add <description>` | Add a new task | `task-cli add "Learn Go"` |
| `list` | List all tasks | `task-cli list` |
| `list todo` | List tasks with status `todo` | `task-cli list todo` |
| `list in-progress` | List tasks with status `in-progress` | `task-cli list in-progress` |
| `list done` | List tasks with status `done` | `task-cli list done` |
| `update <id> <description>` | Update a task's description | `task-cli update 1 "Learn Go and goroutines"` |
| `delete <id>` | Delete a task | `task-cli delete 1` |
| `mark-in-progress <id>` | Mark a task as in-progress | `task-cli mark-in-progress 1` |
| `mark-done <id>` | Mark a task as done | `task-cli mark-done 1` |

---

## Examples

### Adding Tasks

```bash
$ task-cli add "Learn Go programming"
Task added successfully (ID: 1)

$ task-cli add "Build a CLI tool"
Task added successfully (ID: 2)

```

### Listing Tasks

```bash
# List all tasks
$ task-cli list
[1] Learn Go programming - todo (updated: 2026-06-21 10:30:15)
[2] Build a CLI tool - todo (updated: 2026-06-21 10:35:22)

# Filter by status
$ task-cli list todo
[1] Learn Go programming - todo (updated: 2026-06-21 10:30:15)

$ task-cli list in-progress
No tasks with status 'in-progress' found.
```

### Updating a Task

```bash
$ task-cli update 1 "Learn Go programming with goroutines"
Task 1 updated successfully

$ task-cli list
[1] Learn Go programming with goroutines - todo (updated: 2026-06-21 10:45:30)
[2] Build a CLI tool - todo (updated: 2026-06-21 10:35:22)
```

### Marking Task Status

```bash
$ task-cli mark-in-progress 1
Task 1 marked as 'in-progress'

$ task-cli mark-done 2
Task 2 marked as 'done'

$ task-cli list in-progress
[1] Learn Go programming with goroutines - in-progress (updated: 2026-06-21 10:50:15)

$ task-cli list done
[2] Build a CLI tool - done (updated: 2026-06-21 10:55:30)
```

### Deleting a Task

```bash
$ task-cli delete 3
Task 3 deleted successfully

# Trying to delete a task that doesn't exist
$ task-cli delete 99
Task with ID 99 not found
```

---

### Exit Codes

- `0` — Success
- `1` — Error occurred

---

## Data Storage

Tasks are stored in `tasks.json` in the same directory where you run the CLI. The file is created automatically on first use.

### JSON Format

```json
[
  {
    "id": 1,
    "description": "Learn Go programming with goroutines",
    "status": "in-progress",
    "created_at": "2026-06-21T10:30:15.123456Z",
    "updated_at": "2026-06-21T10:50:15.456789Z"
  }
]
```

### Output Format

```
[id] description - status (updated: YYYY-MM-DD HH:MM:SS)
```

---

## Project Structure

```
tasktracker/
├── main.go          # Command routing and entry point
├── operations.go    # CRUD operations (add, update, delete, mark)
├── list.go          # List and display logic
├── storage.go       # File I/O (load, save)
├── models.go        # Task struct definition
├── simple_test.go   # Unit tests
├── tasks.json       # Data storage (auto-generated)
├── go.mod           # Module definition
└── README.md        # Documentation
```

---

## Testing

```bash
# Run all tests
go test -v

```

---

## Dependencies

None. Uses Go standard library only: `encoding/json`, `fmt`, `os`, `strconv`, `strings`, `time`, `testing`.

---

## Quick Reference

```bash
task-cli add "description"              # Add a task
task-cli list                           # List all tasks
task-cli list todo                      # List todo tasks
task-cli list in-progress               # List in-progress tasks
task-cli list done                      # List done tasks
task-cli update <id> "new description"  # Update a task
task-cli delete <id>                    # Delete a task
task-cli mark-in-progress <id>          # Mark as in-progress
task-cli mark-done <id>                 # Mark as done
```
