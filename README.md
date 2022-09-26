# Shopping List CLI

A command line interface for [Shopping List](https://github.com/hidakatsuya/shopping_list).

## Status

Under development

## Usage

```
export SHOPPING_LIST_URL=<url for your shopping-list app>
export SHOPPING_LIST_API_KEY=<your api key>
```
```
$ shopping_list-cli add milk
Successfully added!
```

See the output of the `--help` for details:
```
$ shopping_list-cli --help
A command line interface for Shopping List.

Usage:
  shopping_list-cli [command]

Available Commands:
  add         Add an item to your shopping list
  help        Help about any command

Flags:
  -h, --help   help for shopping_list-cli

Use "shopping_list-cli [command] --help" for more information about a command.
```
