# PkiSauce

Ephemeral Build Time TLS PKI saucing for your intra services GRPC (or not) communications.

## Description

A simple attempt to avoid deploying complex PKI mechanisms, certificates signing, deployment, synchronisation,
protecting the CA, etc..  while keeping a consistent and relatively "safe" to use PKI for your internal needs.

The idea is to build a PKI and leverage the common CI/CD dev habits to our advantage, allowing to integrate intra communication low overhead transport security.


## How it works

Simply, you define the relationships between services/elements of your infrastructure that we define as "daemon" (we're unix folk yes).
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

While we're using it to fullfill our needs during our work and we consider it bring more gains than issues,
it is still a BETA public release, not all features we need & use in some deployed version are baked in yet.

## Instal

```shell
go install ermites.io/pkisauce
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

# TODO

- TLS only policies should NOT include grpc dependencies.
- minimize dependencies.
- certificates properties configuration support (key type?)
- certificate user controlled data  (add user specific data in the generated certs)
- proper debug template of this Ephemeral PKI (list other member of the PKI, only one detailled, etc..)
- java/grpc helpers templates/generator
- c++/grpc helpers templates/generator
- C helpers templates/generator (tls only)
- more...
