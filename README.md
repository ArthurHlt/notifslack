# Notifslack

Send notification to a slack's channel in command line

## Installation

### On *nix system

You can install this via the command-line with either `curl` or `wget`.

#### via curl

```bash
$ sh -c "$(curl -fsSL https://raw.github.com/ArthurHlt/notifslack/master/bin/install.sh)"
```

#### via wget

```bash
$ sh -c "$(wget https://raw.github.com/ArthurHlt/notifslack/master/bin/install.sh -O -)"
```

### On windows

You can install it by downloading the `.exe` corresponding to your cpu from releases page: https://github.com/ArthurHlt/notifslack/releases .
Alternatively, if you have terminal interpreting shell you can also use command line script above, it will download file in your current working dir.

### From go command line

Simply run in terminal:

```bash
$ go get github.com/ArthurHlt/notifslack
```

## Usage

```
NAME:
   notifslack - Send notification to slack

USAGE:
   notifslack --url https://hooks.slack.com/services/XXXX [global options] "my text to send"

VERSION:
   1.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value                   Required. The webhook URL as provided by Slack. Usually in the form: https://hooks.slack.com/services/XXXX
   --channel value, -c value     Optional. Override channel to send message to. #channel and @user forms are allowed.
   --username value, -u value    Optional. Override name of the sender of the message.
   --icon-url value, -i value    Optional. Override icon by providing URL to the image.
   --icon-emoji value, -e value  Optional. Override icon by providing emoji code (e.g. :ghost:).
   --insecure, -k                Ignore certificate validation
   --help, -h                    show help
   --version, -v                 print the version
```