---
order: 0
---

# Getting started

Welcome to the quick and simple way to try out the Cosmos SDK. In this tutorial
will be building Hellochain, a "Hello World" blockchain with a basic `greeter`
module.

For this tutorial we are going to first use th [scaffold tool](https://github.com/cosmos/scaffold) to quivkly spin up aa "blank" application capable of
only bank-like interactions and then add some arbitrary "hello world"
functionality in the form of our `greeter` module. Nothing other than scaffold needs to be downloaded

#TOOO make sure this is right
```bash
./hellochain
├── go.mod
├── Makefile
├── app.go
├── cmd
│   ├── hccli
│   │   └── main.go
│   └── hcd
│       └── main.go
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
For a more thorough and in depth walkthrough of using the scaffold tool. Check out the [scavenge tutorial](../../scavenge/tutorial/03-scaffold.md)#TODO check this link

Use the following command to create the home directory for your "hello world" blockchain app containing the minimum necessary code (specified by the `lvl-1` arg) for a functioning PoS Tendermint blockchain app.
the command requires you to provide a github username and repo name from which it will construct your app's path
```bash
$ scaffold  app lvl-1 <github user name> <repo name>

$ scaffold app lvl-1 hschoenburg hellochain
```

Next iitialize your app as a go module and initialize a git repository to track your changes.

```bash
cd hellochain
$ git init
$ go mod init
```

#TODO
demostrate dsome basic commands to start the chain, etc.

Ok now we can try adding some custom functionsality to our chain. in the spirit of keeping things simple. lets make a "greeter" module that enables us to send greetings (short strings) to accounts on the chain. Tokeep it simple and reduce the risk of minor bugs, lets turn again to the [scaffold tool](https://github.com/cosmos/scaffold) In the cosmos SDK, modules reside in the `x/` directory. Our app should follow this conve=ntion. Use the following command to scaffold a module called `greeter1` like `scaffold app`, `scaffold module requires a github username and repositrory name.
```bash
$ cd hellochain/x
$ scaffold module schoenburg hellochain greeter

Each of these scaffold comands creates files containing unfinished TODOs. This tutorial will walk through turniung those TODOs into a working blockchain application.
```

Want to run this tutorial locally? Install
[Vuepress](https://vuepress.vuejs.org/) and run `vuepress dev` to follow along
at http://localhost:8080
