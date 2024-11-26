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
  "unique": true,
  "sort": "sort",

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

<details>
<summary>sources</summary>

>  An array of source objects that specify the paths to search and their respective depth levels.
>
> Each source object should contain:
> - **`path`**: The directory path to search.
> - **`depth`**: The depth level for searching within the specified path.

</details>

<details>
<summary>expand_output (optional, defaults to `true`)</summary>

> Determines whether the output should be expanded to show additional details. Set to `false` to display only the basic information.

</details>

<details>
<summary>selector (optional, defaults to `fzf`)</summary>

> Specifies the tool used for displaying projects. Available options are:
> - `fzf`: A command-line fuzzy finder.
> - `fzy`: A faster alternative to `fzf`.

</details>

<details>
<summary>unique (optional, defaults to `false`)</summary>

> When set to `true`, the output will only display unique values. Note that enabling this option may slightly impact performance.

</details>

<details>
<summary>sort (optional, defaults to ``)</summary>

> Defines the order in which the output is sorted. Available options are:
> - `asc`: Sorts the output in ascending order.
> - `desc`: Sorts the output in descending order.
>
> Enabling sorting may also have a slight impact on performance.

</details>

For CLI options, run `gosp --help`.
