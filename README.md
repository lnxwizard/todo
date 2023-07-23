# todo
A simple to-do list app that you can use from the terminal.

# About
## Usage
```shell
todo [command]
```

## Examples
```shell
# Adding new to-do
todo add Take a walk
```

```shell
# Listing all to-do's
todo list
```

```shell
# Marking a to-do as completed
todo complete 2
```

```shell
# Marking a to-do as important
todo mark important 8
```

```shell
# Deleting a to-do
todo delete 5
```

```shell
# Unmarking a to-do as completed
todo unmark done 1
```

```shell
# Unmarking a to-do as important
todo unmark important 9
```

## Available Commands
- `add`: Add a new to-do.
- `complete`: Mark a to-do as completed by index.
- `delete`: Delete a to-do by index.
- `list`: List all to-dos.
- `mark`: Mark your to-do.
- `unmark`: Unmark your to-do.

## Available Flags
- `-h`, `--help`: Print help message for todo.
- `-v`, `--version`: Print version.

# Documentation
Documentation about available commands, sub-commands, flags and more is in the [`/docs`](/docs) folder.
