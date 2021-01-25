# HIBP

**HIBP** stand for **H**ave**IB**een**P**wned.

It is a simple tool to check if your password has been pwned and found in some data leak.

This unofficial tool uses the [haveibeenpwned](haveibeenpwned.com) API to protect your privacy, using a mathematical property called [*k*-anonymity](https://en.wikipedia.org/wiki/K-anonymity). Here's is explained how the API maintains your privacy: [haveibeenpwned-privacy](https://www.troyhunt.com/ive-just-launched-pwned-passwords-version-2/#cloudflareprivacyandkanonymity)

## Usage

Use this tool is pretty easy.

```bash
echo "MySuperSecurePassword" | hibp check
```

or

```bash
hibp check "MySuperSecurePassword"
```

### Tip

To avoid logging your password into your history, consider to add the following line to your `.bashrc` file:

```bash
export HISTCONTROL='ignoreboth:erasedups'
```

## Installation

``` bash
go get -v github.com/alegrey91/hibp
```

