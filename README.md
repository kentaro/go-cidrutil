# go-cidrutil

A Go package to retrieve all IP addresses from a CIDR.

## Usage

```go
p := NewParser()
network, err := p.Parse("192.168.0.0/30")
/*
	=> &Network{
		CIDR: "192.168.0.0/30",
		Ones: 30,
		Bits: 32,
		Network: "192.168.0.0",
		Hosts: []string{
			"192.168.0.1",
			"192.168.0.2",
		},
		Broadcast: "192.168.0.3",
    }
*/
```

## Author

[Kentaro Kuribayashi](https://kentarok.org)

## License

MIT
