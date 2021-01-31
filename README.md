# HIBP

**HIBP** stand for **H**ave**IB**een**P**wned.

It is a simple client to communicate with [haveibeenpwned](haveibeenpwned.com) APIs.

This interact (for now) with the **breach** and the **password** endpoints.

## Usage

```
HIBP stand for HaveIBeenPwned.
It is a simple client to communicate with haveibeenpwned (haveibeenpwned.com) APIs.

Usage:
  hibp [command]

Available Commands:
  breach      Check if a specific site has been compromised.
  help        Help about any command
  password    Check if your password has been found in some data set.

Flags:
      --config string   config file (default is $HOME/.hibp.yaml)
  -h, --help            help for hibp
  -t, --toggle          Help message for toggle

Use "hibp [command] --help" for more information about a command.
```

This actually supports passing arguments from **stdin**.

### Tip

To avoid logging your password into your history, consider to add the following line to your `.bashrc` file:

```bash
export HISTCONTROL='ignoreboth:erasedups'
```

## Installation

``` bash
go get -v github.com/alegrey91/hibp && go build
```

