# ipv4mix
`ipv4mix` creates various mixed representations of a specified `IPv4` address. It outputs the `IP` in different formats, including `decimal`, `hexadecimal`, `octal`, and their combinations, providing a versatile view of the address.

For more info see also:
- [inet_aton](https://linux.die.net/man/3/inet_aton);
- [RFC-3936 Uniform Resource Identifier (URI)](https://datatracker.ietf.org/doc/html/rfc3986) (section [Rare IP Address Formats](https://datatracker.ietf.org/doc/html/rfc3986#page-45))

---
## Table of Contents

- [Table of Contents](#table-of-contents)
- [Installation](#installation)
- [Usage](#usage)
- [CTF](#ctf)
- [Contributing](#contributing)

## Installation

Using `go`:

```shell
go install github.com/cmd-tools/ipv4mix@latest
```

Using `homebrew`:

```shell
brew tap cmd-tools/homebrew-tap
brew install ipv4mix
```

Using `docker`:

```shell
docker pull cmdtoolsowner/ipv4mix
```

## Usage

```shell
ipv4mix <IPV4 dot separated>
```

Example
```shell
ipv4mix 8.8.9.9
8.8.9.9
0x8.0x8.0x9.0x9
010.010.011.011
8.0x8.9.9
8.8.0x9.9
8.8.9.0x9
8.010.9.9
8.8.011.9
8.8.9.011
0x8.8.9.9
0x8.0x8.9.9
0x8.8.0x9.9
0x8.8.9.0x9
010.010.9.9
8.0x8.9.0x9
8.0x8.0x9.9
134744329
0x8080909
01002004411
8.8.2313
8.0x8.2313
8.010.2313
8.526601
0x8.526601
010.526601
010.010.2313
010.010.0x909
010.010.04411
```

Using `docker`:

```shell
docker run -i cmdtoolsowner/ipv4mix 8.8.8.8
```

It can be used in combination with other network related tools like `ping`, `curl`, etc.:

```shell
ipv4mix main.go 8.8.8.8 | xargs -I {} ping -c 1 {}
ipv4mix main.go 1.2.3.4 | xargs -I {} curl {}/some-path
```

## CTF
Highlighting [RFC-3936 Uniform Resource Identifier (URI)](https://datatracker.ietf.org/doc/html/rfc3986) (section [Rare IP Address Formats](https://datatracker.ietf.org/doc/html/rfc3986#page-45))

>   These additional IP address formats are not allowed in the URI syntax
    due to differences between platform implementations. However, they
    can become a security concern if an application attempts to filter
    access to resources based on the IP address in string literal format.
    If this filtering is performed, literals should be converted to
    numeric form and filtered based on the numeric value, and not on a
    prefix or suffix of the string form

## Contributing
You want to contribute to this project? Wow, thanks! So please just fork it and send a pull request.
