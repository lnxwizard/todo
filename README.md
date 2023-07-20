# todo
A simple to-do list app that you can use from the terminal.

# About
## Usage
```shell
todo [command]
```

## Examples
```shell
todo add Take a walk
```

```shell
todo list
```

```shell
todo complete 2
```

```shell
todo delete 5
```

## Available Commands
- `add`: Add a new to-do.
- `delete`: Delete a to-do by index.
- `complete`: Mark a to-do as completed by index.
- `list`: List all to-dos.

## Available Flags
- `-h`, `--help`: Print help message for todo.
- `-v`, `--version`: Print version.

### Available Flags for the `list` Command
- `-d`, `--done`: List completed to-do's.

# Documentation
Documentation about available commands, sub-commands, flags and more is in the [`/docs`](/docs) folder.
