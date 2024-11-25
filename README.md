# gosp
A simple tool for quickly selecting projects.

<img alt="Demo" src="examples/demo.gif" width="600" />

## Installation
```sh
go install github.com/gabefiori/gosp/cmd/gosp@latest
```

Once the installation is complete, you can use the `gosp` command along with other commands in your shell.

### Examples with `cd`:

<details>
<summary>Bash</summary>

> Add to your `~/.bashrc` file:
>
> ```sh
> alias fp='cd "$(gosp)"'
> ```

</details>

<details>
<summary>Zsh</summary>

> Add to your `~/.zshrc` file:
>
> ```sh
> alias fp='cd "$(gosp)"'
> ```

</details>

<details>
<summary>Fish</summary>

> Add to your `~/config.fish` file:
>
> ```fish
> alias fp "cd (gosp)"
> ```

</details>

### Using with tmux
You can utilize this [script](/scripts/gosp-tmux.sh), which enables you to easily attach to or switch between Tmux sessions using the `gosp` command for selection.

<details>
<summary>Install</summary>

>```sh
>sudo wget -O /usr/local/bin/tms https://raw.githubusercontent.com/gabefiori/gosp/refs/heads/main/scripts/gosp-tmux.sh
>sudo chmod +x /usr/local/bin/tms
>```

</details>

## Configuration
Create a configuration file at `~/.config/gosp/config.json`:

```json
{
  "expand_output": true,
  "selector": "fzf",
  "sources": [
    {
      "path": "~/your/path",
      "depth": 1
    },
    {
      "path": "/home/you/your_other/path",
      "depth": 3
    }
  ]
}
```

> - `"expand_output"` is optional and defaults to `true`.
> - `"selector"` is optional and defaults to `fzf`. Available options are `fzf` and `fzy`.

## CLI Options
```sh
--config file, -c file      Load configuration from file (default: "~/.config/gosp/config.json")
--selector value, -s value  Selector for displaying the projects (available: "fzf", "fzy") (default: "fzf")
--expand-output, --eo       Expand output (default: true)
--measure, -m               Measure performance (time taken and number of items) (default: false)
--help, -h                  show help
```
