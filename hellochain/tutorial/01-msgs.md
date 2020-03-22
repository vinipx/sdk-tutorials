---
order: 1
---

# Msgs

`Msgs` define your application's state transitions. They are encoded into protobuf and passed
around the network packaged in `Txs`. Messages are "owned" by a single module,
meaning they are routed to one and only one of your application's modules. Messages update their modulessubset of the chain
state. We will give our `greeter` module just one Message, `MsgGreet`, to keep
things simple. `MsgGreet` stores the addresses of the sender and reciever as
well as the body of the "greeting". Evry Message must implement the sdk.Msg interface. We will include this line to verify the interface at compile time.
```golang
var _ sdk.Msg = &MsgGreet{}
```

up[datew] the scaffold in `x/greeter/internal/types/msg.go` to contain the following

<<< @/hellochain/x/greeter/internal/types/msg.go
