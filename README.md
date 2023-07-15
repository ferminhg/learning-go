<p align="center">
    <img alt="Two gophers reading something" src="img/gopher.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center">
  Learning Go
</h1>

## Introduction

This repository contains a series of challenges to learn Go in an incremental way. Each challenge will introduce
new concepts such as struct, testing, data persistence, etc. Furthermore, each challenge will be built on top of the
previous ones, so it is not recommended to try Challenge #2 without completing Challenge #1.

By the end of this series of challenges, we will have a minimal project that will expose a REST API, persist data in
a database and consume some domain events from Kafka in a docker environment.

Disclaimer: This is not intended to teach Go, but to allow us, developers, to have a safe pet project in which
we can validate and put into practice our learnings regarding Go and its ecosystem. So look at this repository as a
coding dojo 🥷.

To take a look at other resources to learn go, please go to: 
https://confluence.mpi-internal.com/display/LHP/Backend+onboarding+resources

## Challenges

As it was mentioned earlier, here you can find a series of incremental challenges to practice coding with Go and its
ecosystem.  
<ins>We consider challenges 1-4 to be mandatory for a basic understanding, challenges 5-6 to be highly recommended and the rest to be nice to do if time allows.</ins>  
In the `challenges` directory you can find a complete list of the challenges and their instructions. At the
moment the available challenges are:

#### Challenge #1. Visibility in Go, structs and data structures

We are going to work on how Go handles visibility by creating our first module and packages. 
Furthermore, we will define our first structs and data structures (arrays, maps and slices).

#### Challenge #2. Error handling and testing

Let's practice with error handling. Then we are going to add first test suites to validate 
the pieces of code which we have been working on until now.

#### Challenge #3. Building an HTTP API

In this third challenge we are going to create our first HTTP API with Gin and to validate it with some fresh new 
contract tests.

#### Challenge #4. Dockerizing our app and using a real database

Let's use go routines to see if the ad was posted also on other platforms

#### Challenge #5. Go routines

In this fifth challenge we are going to deal with goroutines to solve a simple problem.

#### Challenge #6. Introducing Apache Kafka: consumers & producers

In this sixth challenge we are going to deal with Kafka's consumers and producers using the third party library 
`shopify/sarama`

#### Challenge #7. Integration testing

Let's add some integration tests for what we implemented!

#### Challenge #8. TDD approach

In this challenge we'll use TDD approach in order to delete an add

#### Challenge #9. Clean code

Let's take a look back and see where's room for improvements


## How to start

To start with this series of challenges, the minimal requirements are:

### 1. Fork this repository

The way of go with these challenges is by forking this repository and working in your own version of the repository.
This way, all of us can have tons of solutions to learn from!

### 2. Install Go

Follow instructions from https://go.dev/doc/install. 

On macOS if you are currently using Homebrew is even easier, just type in your terminal:
```bash
brew install go
```

### 3. Set up your favourite IDE

* For Visual Studio follow these [instructions](https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code)

* For IntelliJ just install this [plugin](https://plugins.jetbrains.com/plugin/9568-go?_ga=2.122569868.664457569.1681124920-637112615.1649069055&_gl=1%2A1upp818%2A_ga%2ANjM3MTEyNjE1LjE2NDkwNjkwNTU.%2A_ga_9J976DJZ68%2AMTY4MTEyNDkyMC41My4wLjE2ODExMjQ5MjAuNjAuMC4w).

Furthermore, JetBrains has its own IDE for Go: [GoLand](https://www.jetbrains.com/go/promo/). However, if you are already 
familiar with IntelliJ, the recommendation should be to continue with IntelliJ.

### 4. Have fun and share!

Most of the time, there will be different ways of solving a challenge, the purpose of this repository is double: to 
practice with Go's ecosystem and to set up interesting discussions between us, so we encourage you to share your solutions
or your doubts with the community (at slack: #lbc-backend-guild) to learn from each other ❤️

That's all! You are completely ready to dive into the first challenge 🚀
