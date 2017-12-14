# Lottery Go

Simple CLI and library to check if your Spanish Christmas Lottery numbers have been awarded.

Examples:

```bash
$ lottery-go check -h
Checks if number has been awarded

Usage:
  lottery-go check number [flags]

Flags:
  -h, --help   help for check

Global Flags:
  -v, --verbose   prints verbose information during command execution
```

```
$ lottery-go check 1337
Error: raffle hasn't started yet
Usage:
  lottery-go check number [flags]

Flags:
  -h, --help   help for check

Global Flags:
  -v, --verbose   prints verbose information during command execution

raffle hasn't started yet
```

```go
func main() {
	p, err := lottery.CheckNumber(n)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
```
