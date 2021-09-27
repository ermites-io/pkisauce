# PkiSauce

Ephemeral Build Time TLS PKI saucing for your intra services GRPC (or not) communications.

## Description

A simple attempt to avoid deploying complex PKI mechanisms, certificates signing, deployment, synchronisation,
protecting the CA, etc..  while keeping a consistent and relatively "safe" to use PKI for your internal needs.

The idea is to build a PKI and leverage the common CI/CD dev habits to our advantage, allowing to integrate intra communication low overhead transport security.

While we're using it daily on our projects and we consider it safe for our needs, it is still sort of BETA and a WORK IN PROGRESS, it is NOT polished and not all features are baked in yet.
Obviously, we welcome contributions & constructive ideas, that does NOT MEAN _every idea or PR_ will be included or planned.

## How it works

Simply, you define your services/elements/daemon and their relationships / interdependence, we define one entity as "daemon" (we're unix folk yes).
You then say which one should access which grpc service/RPC call of another defined "daemon". 
`pkisauce` will generate the one time use PKI, which includes: the CA,cert/key, services cert/key and the policy code in an interceptor format for the output language supported.

Example `myservices.conf`:
```
daemon "servA" path "../cmd/grpc-pkscd" use go-grpc {
	pass "client1" on "Public" rpc any;
	pass "client2" on "Restricted" rpc { "CallThere" };
}

daemon "client1" path "../cmd/grpc-pkscc" use go-grpc {}
daemon "client2" path "../cmd/grpc-pkscc" use go-grpc {}
```

Run `pkisauce myservices.conf` to generate the appropriate helpers for your services in golang

see the [example](example) directory for an example on how to use it.

## Requirements

* golang 1.17+
* gRPC services to use that on (but actually you could build a pem based pki only if you wish)
* the go.mod stuff


## Install

```shell
go install ermites.io/pkisauce@latest
```

```shell
git clone...
cd pkisauce
go build -a
```

## Features

* mTLS 1.3+ communication.
* no CA to handle (CA is destroyed after the build).
* service communication consistency (avoid version conflicts between builds).
* GRPC service/call granularity.
* Generate TLS host only policies for non GRPC services.
* TLS certificates/keys are using Curve25519.
* TLS certificates/keys are generated & embedded in the generated code.
* TLS certificates/keys are unpacked at runtime.
* Golang GRPC policies are generated in unary interceptor code ready to be embedded your services.
* Golang GRPC call granularity, daemonA can only call daemonB/ServiceX/Call1.
* PEM generation support.
* no additionnal dependencies in the generated code (beside GRPC obviously).


# Usage

This will prepare one single Ephemeral PKI composed of pkisauce config files 
and generate TLS/GRPC helpers to apply the policy defined in those pkisauce conf files.

```shell
pkisauce [options] <conf1> <conf2> <conf3> ... <confN> 
```

all the files provided are considered part of the SAME PKI.

# TODO

- Graph generation (dot, gephy, etc..) for document purposes.
- TLS only policies should NOT include grpc dependencies.
- minimize dependencies.
- certificates properties configuration support (key type?)
- certificate user controlled data  (add user specific data in the generated certs)
- proper debug template of this Ephemeral PKI (list other member of the PKI, only one detailed, etc..)
- java/grpc helpers templates/generator
- c++/grpc helpers templates/generator
- C helpers templates/generator (tls only)
- GRPC persistent connection load balancing included.
- more...
