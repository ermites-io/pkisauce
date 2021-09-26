# PkiSauce Example

A small example, hopefully we will complete as time permits.

## GRPC definitions

We defined an example service definition like you would do for any GRPC service:
```protobuf3
syntax = "proto3";

option go_package = "./;svc";

package main;

message Request {
        int64 cmd = 1;
        string args = 2;
}

message Response {
        int64 ret = 1;
        string msg = 2;
}


service Public {
        rpc CallMe(Request) returns (Response) {}
        rpc Ping(Request) returns (Response) {}
}

service Restricted {
        rpc CallHere(Request) returns (Response) {}
        rpc CallThere(Request) returns (Response) {}
}

service Different {
        rpc CallDifferent(Request) returns (Response) {}
        rpc Shazz(Request) returns (Response) {}
}
```


## PkiSauce definitions

These are our PkiSauce definitions, that we put in the directories.
```
// comments this is our Service A
daemon "servA" path "../cmd/grpc-pkscd" use go-grpc {
        pass "client1" on "Public" rpc any;
        pass "client2" on "Restricted" rpc { "CallThere" };
        pass "client3" on "Different" rpc { "Shazz" };
        pass "client3" on "Public" rpc any;
        pass "servB" on "Porcasse" rpc any;
        pass any on "Public" rpc any;
}

/*
 *  this is our Service B in a different comment
 */
daemon "servB" path "../cmd/grpc-pkscd" use go-grpc {
        pass "client4" on "Restricted" rpc { "CallThere" };
}

/* This is our service C */
daemon "servC" path "../cmd/grpc-pkscd" use go-grpc {
}

/* These are our clients... :) */
daemon "client1" path "../cmd/grpc-pkscc" use go-grpc {}
daemon "client2" path "../cmd/grpc-pkscc" use go-grpc {}
daemon "client3" path "../cmd/grpc-pkscc" use go-grpc {}
daemon "client4" path "../cmd/grpc-pkscc" use go-grpc {}
```



## Relationship Graph

TBD


## Using the GRPC generated code.

Please for now see example `cmd/grpc-pkscd/` for the server and `cmd/grpc-pkscc/` for the client.


## Policy definition granularity / subtleties

TBD
