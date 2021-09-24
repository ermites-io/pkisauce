# PKI Sauce

Ephemeral One Time Use/Build Time TLS PKI saucing for your intra services GRPC communications.

# Description

An idea to avoid deploying complex PKI mechanisms, certificates signin, deployment, synchronisation,
protecting the CA, etc..  while keeping a consistent and relatively "safe" to use PKI for your internal needs.

Simply, you define the relationships between elements of your infrastructure, that we describe as "daemon" (we're unix folk yes)
and say which one should access which Service/RPC call of another defined "daemon".

Example:
```
daemon "servA" path "../cmd/grpc-pkscd" use go-grpc {
	pass "client1" on "Public" rpc any;
	pass "client2" on "Restricted" rpc { "CallThere" };
}

daemon "client1" path "../cmd/grpc-pkscc" use go-grpc {}
daemon "client2" path "../cmd/grpc-pkscc" use go-grpc {}
```

and run pkisauce with that description file to generate the appropriate helpers for your GRPC services usign golang

While we're using it to fullfill our needs during our work and we consider it bring more gains than issues,
it is still a BETA public release, not all features we need & use in some deployed version are baked in yet.

# Features

* mTLS 1.3+ communication.
* no CA to handle (CA is destroyed after the build).
* service communication consistency (avoid version overlap).
* GRPC Service/Call granularity.
* Generate TLS host only policies for non GRPC services.
* TLS certificates/keys are using Curve25519.
* TLS certificates/keys are generated & embedded in the generated code.
* TLS certificates/keys are unpacked at runtime.
* GRPC policies are generated using an interceptor and embedded in the generated code.
* PEM generation support
* no additionnal dependencies in the generated code (beside GRPC obviously)

# Building

$ go build -a

# Usage

This will prepare one single Ephemeral PKI composed of pkisauce config files 
and generate TLS/GRPC helpers to apply the policy defined in those pkisauce conf files.

$ pkisauce <conf1> <conf2> <conf3> ... <confN> 

# How it works

TBD

# TODO

- TLS only policies should NOT include grpc dependencies.
- minimize dependencies.
- certificates properties configuration support (key type?)
- certificate user controlled data  (add user specific data in the generated certs)
- proper debug template of this Ephemeral PKI (list other member of the PKI, only one detailled, etc..)
- java helpers template
- c++ helpers templates
- C helpers templates
- more...
