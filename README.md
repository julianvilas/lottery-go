# Lottery Go

Simple CLI and library to check if your Spanish Christmas Lottery numbers have been awarded.

## CLI

It has two available commands:

```bash
$ lottery-go
Checks Spain Christmas Lottery Results using EL PAIS API

Usage:
  lottery-go [command]

Available Commands:
  check       Checks if number has been awarded
  help        Help about any command
  monitor     monitors your lottery numbers and notifies you

Flags:
  -v, --verbose   prints verbose information during command execution

Use "lottery-go [command] --help" for more information about a command.
```

The command `check` is for checking a number against the API just once:

```bash
$ lottery-go check 11111 -v
[{Num:11111 Prize:0 Timestamp:1513241029 Status:0 Error:0}]
```

While the command `monitor` is for checking a list of numbers until the raffle has finished and the official results are available. In addition, when a number you are checking has been awarded, it notifies you by email.

```bash
$ lottery-go monitor test@example.com 11111 22222 -v
2017/12/18 23:14:14 checking [11111 22222] numbers, notifying test@example.com
2017/12/18 23:15:14 checking numbers again
2017/12/18 23:15:14 [{Num:11111 Prize:0 Timestamp:1513241029 Status:0 Error:0} {Num:22222 Prize:0 Timestamp:1513241029 Status:0 Error:0}]
2017/12/18 23:15:14 raffle has not started yet. Skipping...
```

In order to send the email, the command is using [AWS SES](https://aws.amazon.com/ses/). You need to setup your AWS credentials as specified in the aws-sdk-go [documentation](https://github.com/aws/aws-sdk-go#configuring-credentials)

## Library

The library function CheckNumbers accept a list of numbers as a variadic list of arguments and returns an array with the parsed response of the EL PAIS API.

```go
var n int = 11111 

func main() {
	p, err := lottery.CheckNumbers(n)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)
}
```
