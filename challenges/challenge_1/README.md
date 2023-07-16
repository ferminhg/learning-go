<p align="center">
    <img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-1.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center">
  Challenge #1. Package visibility, structs and data structures
</h1>

In this first challenge we are going to introduce the business problem in which we will be working during the entire set
of challenges. Furthermore, we will create our first, and main, module to set up our project.

Later, we are going to model our business requirements and even implement the repository pattern with some in-memory 
data structures. There we Go! 🚀

## The business problem

We are in charge of creating another (another?) marketplace. This will be the definitive one (come on, trust me). And we
need to develop the use cases related with posting and fetching ads. For the sake of simplicity there won't be complex
searching requirements apart from searching by ad identifier or a simple "find a few of them" use case.

So, as it stands right now (and we made product people to swear that these requirements aren't going to change) what we 
need is:
- [ ] The capability of posting an ad
- [ ] The capability of finding an ad by ID
- [ ] The capability of finding a set of 5 (as maximum) ads to compose a listing

Regarding the ad, it is composed by:
* id (UUID)
* title
* description
* price
* date and time of publication

## How to start

The first step is to create our first, and main, Go module. To do that, we need to run a command in our favorite CLI at 
the root of the project (replace with your GitHub user before running it):

````bash
go mod init github.mpi-internal.com/{githubUser}/learning-go
````

This command will take care of create the module pointing to your remote repository.

Once our module has been created, we only need an entrypoint for our new Go application. Create a new file, at the root
of the project, named `main.go` with the follow content:

````go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
````
With this we are totally ready! You can run your new application that will print out a simple greeting message.

## Instructions

* To complete this challenge you will be required to implement the three use cases we mentioned earlier. To do that,
we are going to need, also, to implement a proper modeling of the `Ad` and an in-memory repository (which can be based
on arrays, maps or the data structure you want).

* Keep in mind that we want to practice with the package visibility and how we can interact between different packages, so
try to split all this logic across different packages. If you feel lost, take a look at [1] to get an overview of how a project layout looks like 
in Go, but it is not so important at this point.

* For the sake of simplicity, we are coding only the happy paths, don't worry (yet) about possible errors in our use cases
(i.e. what happens when we post an ad with an id that is currently used).

* Finally, to test these features, by the moment, we are going to use the main function of `main.go` to run them and log
the results through the standard output.

## Resources

[1] Non-official project structure standard for Go projects: https://github.com/golang-standards/project-layout#go-directories
