## Build

```sh
make

# building for Windows, the binary will have .exe
GOOS=windows make
```

## Usage

```sh
# cd into the binary location
cd ./bin/darwin_amd64

# call the https://${url} using TLS 1.2 and 1.3
tls-issue-repro -url ${url}

# call the https://${url} using TLS 1.2
tls-issue-repro -url ${url} -withTLS12MaxVersion
```
