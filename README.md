# tureng-cli (TURENG CLI APP)

## REQUIREMENTS (IF YOU ARE USING LINUX, YOU CAN USE THE BINARY BUT FOR COMPILE ;) 
go1.17.6 linux/amd64

## FOR INSTALL

```bash
$git clone https://github.com/bur4kgazi/tureng-cli.git
$cd tureng-cli
$go mod tidy
$go build -o tureng main.go 
$./tureng --help
```

## USAGE

1. lang string => which language do you translate (default "tr-en")
2. ret int     => return word length (default 2)
3. url         => show url
4. w string    => word for searching

```bash
$./tureng -w hello -lang fr-en
1 : bonjour
2 : salut

Done :D

```
