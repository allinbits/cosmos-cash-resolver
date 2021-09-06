# Cosmos Cash Resolver

This is a DID universal resolver driver for the Cosmos Cash DID module 


## Usage

To run the server use:

```sh
➜ cosmos-cash-resolver --help
Usage of cosmos-cash-resolver:
  -grpc-server string
    	The target grpc server address in the format of host:port (default "localhost:9090")
  -listen string
    	The REST server listen address in the format of host:port (default "0.0.0.0:2109")
  -mrps int
    	Max-Requests-Per-Seconds: define the throttle limit in requests per seconds (default 10)
```

### Configuration

The resolver can be also configured using environment variables:

- `GRPC_SERVER_ADDRESS` - target grpc server address in the format of host:port
- `LISTEN` - listen address in the format of host:port 
- `MRPS` - max requests per seconds, define the throttle limit in requests per seconds


### Universal resolver driver 

Cosmos Cash Resolver implements a [universal resolver](https://github.com/decentralized-identity/universal-resolver) compatible REST API

The configuration for this resolver are the following:

```
{
    "pattern": "^(did:cosmos:.+)$",
    "url": "http://uni-resolver-driver-did-uport:8081/",
    "testIdentifiers": [
        "did:cosmos:cosmoscash-testnet:123456789",
        "did:cosmos:key:cosmos1u7clngyucn867fm2za0s869yvln9aur8zjujxe"
    ]
}
```
