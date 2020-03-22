---
order: 0
---

# Getting started

Welcome to the quick and simple way to try out the Cosmos SDK. In this tutorial
will be building Hellochain, a "Hello World" blockchain with a basic `greeter`
module.

For this tutorial we are going to first use the [scaffold tool](https://github.com/cosmos/scaffold) to quickly spin up a "basic" application capable of
only bank-like interactions and then add some arbitrary "hello world"
functionality in the form of our `greeter` module. Nothing other than scaffold needs to be downloaded

```bash
When we are finished our directory structure should look something like this
./hellochain
├── go.mod
├── Makefile
app
└──app.go
└──export.go
└──app.go
├── cmd
│   ├── aud
│   └── main.go
|   └── genaccounts.go
│   └── acli
│   └── main.go
└── x
    └── greeter
        ├── client
        │   ├── cli
        │   │   ├── query.go
        │   │   └── tx.go
        ├── internal
            ├── types
            |  ├── msgs.go
            |  └── types.go
            ├─ keeper
              └── querier.go
              ├── keeper.go
        ├── alias.go
        ├── module.go
        ├── handler.go

```

# Scaffold
we will be using the [scaffold tool](https://github.com/cosmos/scaffold) to jumpstart our app and skip much of the boilerplate required.

Install and test scaffold and test it with the following commands.

```bash
it clone git@github.com:cosmos/scaffold.git
cd scaffold
make tools
make install
scaffold --help
```

For a more thorough and in-depth walkthrough of the scaffold tool, check out the Scaffold section of the [scavenge tutorial](../../scavenge/tutorial/03-scaffold.md)

First we will grnerate a bassic blockchain app with the following scaffold command. This boilerplate app  is a functioning Proof of Stake Tendermint BFT blockchain  composed from SDK modules such as auth, bank, staking, slashing and distribution.
the command requires you to provide the github username and repo name from which it will construct your app's path
```bash
$ scaffold  app lvl-1 <github user name> <repo name>

# for example
$ scaffold app lvl-1 hschoenburg hellochain
```

Next initialize a git repository to track your changes.

```bash
cd hellochain
$ git init
```

Next just to test that everything is set up correctly ( but not yet started)  lets try to run our daemon and initialize our chain. 
```bash
$ cd hellochain
go run cmd/aud/ genaccounts.go main.go test-chain-id 
```

You should see a JSON object in STDOUT. This gets saved to genesis.json in your daemon home directory, configured in app/app.go. If an error is thrown, compare your code to the [hellochain repo](https://github.com/cosmos/sdk-tutorials/tree/master/hellochain) minus the `tutorial` dir of course.

Ok now that our basic app daemon is working we can try adding some custom functionality to our chain. in the spirit of keeping things simple. lets make a "greeter" module that enables us to send  short strings, greetings to accounts on the chain. To reduce the risk of minor bugs, lets turn again to the [scaffold tool](https://github.com/cosmos/scaffold) In the cosmos SDK, modules reside in the `x/` directory. Our app should follow this convention. Use the following command to scaffold a module called `greeter` like `scaffold app`, `scaffold module requires a github username and repository name wityh which it constructs include paths. These paths should match the directory structure where you have saved your app (following Go convention)  ( for example my work is saved in $GOPATH/src/github.com/hschoenburg/)

```bash
$ cd hellochain/x
$ scaffold module hschoenburg hellochain greeter
```

Each of these scaffold comands creates files containing unfinished TODOs. From here on this tutorial will walk through turniung those TODOs into a working blockchain application, going through one component at a time, starting with pur module's messages.
```

Want to run this tutorial locally? Install
[Vuepress](https://vuepress.vuejs.org/) and run `vuepress dev` to follow along
at http://localhost:8080
