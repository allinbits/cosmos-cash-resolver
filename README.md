# Cosmos Cash Resolver

This is a DID universal resolver driver for the Cosmos Cash DID module 


## Usage

To run the server use:

```sh
➜ cosmos-cash-resolver --help
Usage of /tmp/go-build2385902778/b001/exe/main:
  -grpc-server-address string
    	The target grpc server address in the format of host:port (default "localhost:9090")
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

### Endpoints

The Cosmos Cash Resolver exposes the following endpoints

#### Universal resolver endpoint

The universal resolver endpoint aims to be compatible to the [universal resolver](https://github.com/decentralized-identity/universal-resolver) REST API endpoint

```curl
> GET /identifier/did:cosmos:net:cash:bob HTTP/1.1
> Host: localhost:2109
> User-Agent: curl/7.76.1
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/did+ld+json
< Vary: Origin
< Date: Mon, 11 Oct 2021 18:01:05 GMT
< Content-Length: 1028
< 

{
  "didDocument": {
    "@context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:net:cash:bob",
    "verificationMethod": [
      {
        "controller": "did:cosmos:net:cash:bob",
        "id": "did:cosmos:key:cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a#cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a",
        "publicKeyMultibase": "F022e7e57627fce23fb684879ae1d486ec9c81f76962213943ca8f31c6cbfd39ef0",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [
      {
        "id": "agent",
        "type": "DIDCommMessaging",
        "serviceEndpoint": "http://localhost:8091"
      }
    ],
    "authentication": [
      "did:cosmos:key:cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a#cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a"
    ],
    "keyAgreement": [
      "did:cosmos:key:cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a#cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a"
    ]
  },
  "didDocumentMetadata": {
    "versionId": "e0ff046939a72993ccd838fc7e38750ec281c9a83f6c682ab914fb67bc0383a6",
    "created": "2021-10-11T14:09:27.51467081Z",
    "updated": "2021-10-11T14:09:37.609840653Z"
  },
  "didResolutionMetadata": {
    "did": {
      "method": "cosmos",
      "methodSpecificId": "net:cash:bob"
    }
  }
}

```

#### Hyperledger Aries specific endpoint 

A Hyperledger Aries specific endpoint is available at the following path

> Note: this endpoint is mostly used for R&D purposes

```curl
> GET /identifier/aries/did:cosmos:net:cash:bob HTTP/1.1
> Host: localhost:2109
> User-Agent: curl/7.76.1
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/did+ld+json
< Vary: Origin
< Date: Mon, 11 Oct 2021 18:03:14 GMT
< Content-Length: 832
< 

{
  "didDocument": {
    "@context": [
      "https://www.w3.org/ns/did/v1"
    ],
    "id": "did:cosmos:net:cash:bob",
    "verificationMethod": [
      {
        "controller": "did:cosmos:net:cash:bob",
        "id": "did:cosmos:key:cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a#cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a",
        "publicKeyHex": "022e7e57627fce23fb684879ae1d486ec9c81f76962213943ca8f31c6cbfd39ef0",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [
      {
        "id": "agent",
        "type": "DIDCommMessaging",
        "serviceEndpoint": "http://localhost:8091"
      }
    ],
    "authentication": [
      "did:cosmos:key:cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a#cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a"
    ],
    "keyAgreement": [
      "did:cosmos:key:cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a#cosmos1svxs3j3xz9at0ku5estaffa86qypce5k4gsq3a"
    ],
    "created": "2021-10-11T14:09:27.51467081Z",
    "updated": "2021-10-11T14:09:37.609840653Z"
  }
}
```
