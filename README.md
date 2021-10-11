# Cosmos Cash Resolver

This is a DID universal resolver driver for the Cosmos Cash DID module 


## Usage

To run the server use:

```sh
➜ cosmos-cash-resolver --help
Usage of /tmp/go-build2385902778/b001/exe/main:
  -grpc-server-address string
    	The target grpc server address in the format of host:port (default "localhost:9090")
  -legacy-format
    	Apply legacy format to DID document content (false by default)
  -listen-address string
    	The REST server listen address in the format of host:port (default "0.0.0.0:2109")
  -mrps int
    	Max-Requests-Per-Seconds: define the throttle limit in requests per seconds (default 10)
```

### Configuration

The resolver can be also configured using environment variables:

- `GRPC_SERVER_ADDRESS` - target grpc server address in the format of host:port
- `LISTEN` - listen address in the format of host:port 
- `MRPS` - max requests per seconds, define the throttle limit in requests per seconds
- `LEGACY_FORMAT` - apply legacy format to DID document content (false by default)


#### Legacy Format

The legacy format option will rewrite the did verification methods `pubKeyMultibase` to `pubKeyHex` if there is the `F` prefix. 
This is useful since the current version of the aries-framework-go does not support yet the `pubKeyMultibase` format.

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
