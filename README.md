# inspect-test
Two simple programs to test if HTTP traffic on non-standard ports
is checked by your service provider (and blocked).  The test programs 
are written in Go.  It works on Windows/Mac/Linux.

## Setup
Compile both of the programs.
```
git clone https://github.com/theborak/inspect-test.git
cd inspect-test
go build edge.go
go build reverse_tun.go
```

On a machine (virtual or otherwise) that is available on the Internet,
run the edge program and specify a non-standard HTTP port.
```
./edge -port 9000
```

On a separate machine behind your service providers router (home network), 
run the reverse_tun program on a machine.
```
./reverse_tun -port 9000 -host <ip or public machine>>
```

## How it works
Once the edge program receives a connection from the reverse_tun program
it will begin writing HTTP requests to the client (establishing a reverse tunnel).
The reverse_tun program will write back simple HTTP 200 responses.

### The Purpose
These two programs can be used to determine if your service provider is 
inspecting traffic beyond examining ports for incoming traffic. 
