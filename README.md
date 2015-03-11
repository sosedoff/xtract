# xtract

Extract data from text sources

## Usage

CLI flags and options:

```
Usage of ./xtract:
  -f="": Extraction format: (ipv4, email)
  -ignore="": List of ignore values
  -uniq=false: Return only unique matches
```

Examples:

```
# find all ipv4 addresses from the log
cat auth.log | xtract -f ipv4 -uniq

# find all emails
cat mylog.log | xtract -f email -uniq -ignore email@example.com

# scan all ipv4 addresses excluding localhost from remote log
ssh user@host.com cat /var/log/mylog.log | xtract -f ipv4 -uniq -ignore 127.0.0.1
```

## Install

Get it:

```
go get github.com/sosedoff/xtract
```

Build from source:

```
go build
```

## License

MIT