# Task CLI

A simple CLI task manager written in Go.
Allows you to add, update, delete, mark, and list tasks, storing them in a JSON file.

This project is meant as a learning exercise, focusing on:

- clean architecture
- separation of domain and infrastructure
- idiomatic Go
- proper error handling
- minimal dependencies

## Features

- Add new tasks
- Update task descriptions
- Delete tasks
- Change task status (backlog / ready / wip / done)
- List all tasks
- Store tasks in JSON
- No database required

## Requirements

- Go 1.20+

## Build

``` bash
go build -o task-cli
```

This will create a binary named task-cli.

## Usage

### Add a new task

``` bash
./task-cli add "Buy groceries"
```

### Update a task

``` bash
./task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a task

``` bash
./task-cli delete 1
```

### Change task status

``` bash
./task-cli mark 1 backlog
./task-cli mark 1 ready
./task-cli mark 1 wip
./task-cli mark 1 done
```

### List all tasks

``` bash
./task-cli list
```

## Task Statuses

| Status  | Meaning          |
|---------|------------------|
| backlog | task created     |
| ready   | ready to start   |
| wip     | work in progress |
| done    | completed        |

## Data Storage

All tasks are saved to `data.json`.  
The file is created automatically on first run.

## Backlog

- Filter tasks by status (e.g., list done, list wip)
- Support multiple task files or workspaces
- Add due dates / reminders for tasks
- Implement search by keyword in task descriptions
- Improve CLI UX: colored output, better table formatting
- Add undo/redo functionality for task operations
- Unit and integration tests for all commands
- Optional persistent storage using SQLite or BoltDB
- Configurable JSON storage path via environment variable or flag
- Auto-sort tasks by status or creation date
