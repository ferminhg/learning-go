<p align="center">
    <img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-2.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center" style="text-align: center;">
  Challenge #2. Error handling and testing
</h1>

In this second challenge we are going to practice testing and error handling.

On one hand we are going to implement some tests suites to cover the code we wrote in the first challenge. We are going
to get a little help from `stretchr/testify` which is a toolkit with commont **assertions** and **mocks** that
complements the standard library.

On the other hand, we will start handling our first errors with Go, and we will discuss which situations should be
handled
as errors and which not.

## Instructions

### Testing

As we mentioned above in the introduction, to practice testing with the standard library and `stretchr/testify` we are
going to revisit the previous challenge and to cover with unit tests both, the in-memory repository and the use cases:

- [ ] Let's start with the repository which is a simple unit test suite since it has no collaborators and no mocks are needed.
- [ ] Then, to test our use cases, typically, we are going to need to mock the repository to reduce the complexity of the test suite. Take a look at the resources section to find some useful documentation about building mocks with `stretchr/testify`.

### Error handling

To practice with error handling, three scenarios are proposed:

- [ ] Our product team comes with a new business rule: an ad cannot contain a description longer than 50 characters. How can we deal with that? Where should we place this logic and how an error could be handled there? Tip: since it is a business rule, it should be placed as near to the business as possible.
- [ ] Let's talk about the `find an ad by ID` use case. What happens if given an ID there is no Ad in the repository to return? Is this an exceptional (error) flow? Should be handled as an error? If not, how can we handle this scenario?
- [ ] Let's pretend we are not implementing an in-memory repository, but a repository based on some classic database (i.e. PostgreSQL). Eventually, every single method of the repository could return an error due to connectivity issues, or timeouts, etc. So, let's model our repository to be able of returning errors on those scenarios and see how can we handle  them from the use cases.

## Resources
1. Testify documentation for mock and asssert packages: https://github.com/stretchr/testify
2. Library used in Leboncoin : https://medium.com/yemeksepeti-teknoloji/mocking-an-interface-using-mockery-in-go-afbcb83cc773
3. Mocks in Go. Tests with Testify Mock: https://dev.to/salesforceeng/mocks-in-go-tests-with-testify-mock-6pd
4. Go by example: Errors: https://gobyexample.com/errors
5. Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors
6. Effective error handling in Golang: https://earthly.dev/blog/golang-errors/
