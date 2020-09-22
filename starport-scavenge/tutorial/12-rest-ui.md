---
order: 12
---

# Implementing REST and a UI (Bonus)

As of now, we have a functioning game through the CLI, but what if we wanted to implement a simple user interface for our app?

We can build off the existing UI scaffolded by `starport app`, as well as add some functions to our `rest` folder.

Let's start with our REST endpoints in `./x/scavenge/client/rest`.

## REST endpoints

In this folder, we have our `query<Type>.go` and `tx<Type>.go` files, as well as our `rest.go` file which registers the endpoints. These files will follow the same structure as our `../cli/query<Type>.go` and `../cli/tx<Type>.go` files, and should look as follows:


#### `queryCommit.go`
<<< @/starport-scavenge/scavenge/x/scavenge/client/rest/queryCommit.go

#### `txCommit.go`
<<< @/starport-scavenge/scavenge/x/scavenge/client/rest/txCommit.go

#### `queryScavenge.go`
<<< @/starport-scavenge/scavenge/x/scavenge/client/rest/queryScavenge.go

#### `txScavenge.go`
<<< @/starport-scavenge/scavenge/x/scavenge/client/rest/txScavenge.go

Once we have these files, we need to register them in `rest.go`.

#### `rest.go`
<<< @/starport-scavenge/scavenge/x/scavenge/client/rest/rest.go


Our rest endpoints are complete!

## User interface

When visiting our REST interface, we'll have a barebones UI that lets us view and create commits. We need to tinker with the UI to modify the fields used when creating a new type, as well as adding a hashing mechanism, and adding a section that allows us to reveal a scavenge.

#### `./vue/src/store/app.js`

In this file, we can modify the fields needed when creating a new type.

<<< @/starport-scavenge/scavenge/vue/src/store/app.js

### Adding a new section to page

Since we need to be able to reveal a scavenge, as well as have some functionality to hash a scavenge on our front end application. We can implement this by adding new components `src/components/HashString.vue` and `src/components/RevealSolutionLayout.vue` as follows:

#### `HashString.vue`

In this component, we'll be using a library `crypto-js` which contains an implementation of a hashing function, so we also need to run `npm i --save cryptojs` in our `./vue` directory.

<<< @/starport-scavenge/scavenge/vue/src/components/HashString.vue

#### `RevealSolutionLayout.vue`

<<< @/starport-scavenge/scavenge/vue/src/components/RevealSolutionLayout.vue

Once this is done, we need to add the new components `<hash-string />` and `<reveal-solution-layout />` to our `views/Index.vue` file.

#### `Index.vue`
<<< @/starport-scavenge/scavenge/vue/src/views/Index.vue


Next, we can interact with our application via a UI!

