<p align="center">
    <img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-5.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center" style="text-align: center;">
  Challenge #5. Go routines
</h1>

In this 5th challenge we are going to practice with goroutines. Apparently we are developing a new feature for posting
ads in an easier way. To do that, our Machine Learning team is developing some different ML models to produce a nice
ad's description given a title. We want to test which one of these models performs better.

## Instructions

First of all, don't worry about ML models, it's just an excuse to play with goroutines 😄. We can imagine these ML
models as a services that return some random or event static result with a delay to simulate some processing or some 
network handling.

So, we have been told that our company have developed 3 different ML models that given a title they return to us a proper description. In fact,
these models does not return a simple string, but it returns also a value between 0 and 1. This value indicates the confidence 
level that we should have on its description.

Ideally, this feature should be exposed through our HTTP API. We can imagine a POST endpoint that accepts a payload with
the title and returns a response with the different descriptions we generate with our ML models.

````http request
POST http://localhost:8080/description-generator
````
````json
{
  "title": "iPhone 13 Pro 64gb almost new"
}
````
And the response:

````json
{
  "descriptions": [
    {
      "description": "Like-new iPhone 13 Pro 64 available! Meticulously cared for, exceptional performance, and an unbeatable price. Immerse yourself in the 6.1\" Super Retina XDR display, capture stunning moments with the triple-camera system, and enjoy the security of Face ID. Don't miss this chance to own a device that's been cherished and pampered!",
      "confidence": 0.8
    },
    {
      "description": "Get your hands on an iPhone 13 Pro 64 in pristine condition! Carefully maintained with exceptional performance and an unbeatable price. Immerse yourself in the stunning 6.1\" Super Retina XDR display, capture unforgettable moments with the advanced triple-camera system, and enjoy the security of Face ID. Don't miss this opportunity to own a meticulously cared for device!",
      "confidence": 0.75
    }
  ]
}
````
However, think that this challenge is about Go routines, so if you don't have time enough for completing it, prioritize 
the following exercises and forget about the HTTP API.

So, once the background has been set, let's get our hands a bit dirty. First of all, we will need to simulate these ML models.
As we mentioned, we can simply create a function that returns a static/random string and that generates a random value to apply
some delay and to expose a confidence level:

````go
func GenerateDescription(_ string) Description {
	random := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(random))

	return Description{
		value:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam vel justo in nunc fringilla bibendum",
		confidence: float32(random) / 1000.0,
	}
}
````
The proposed exercises are:
1. Implement a service that invokes the 3 ML models (or invoke three times the same one) to get 3 different descriptions
sequentially and return them.
2. Use goroutines to invoke the 3 ML models in a concurrent way.
3. We need to speed up this new feature, so we want to set a timeout of 350ms to every ML model to generate a description. 
We need at least one valid description in time to return a successful response. Tip: Try to implement the timeout feature with Contexts.

## Resources
1. https://go.dev/tour/concurrency/1
2. https://go.dev/doc/effective_go#goroutines
3. https://go.dev/blog/context
4. Handling timeouts with contexts: https://medium.com/geekculture/timeout-context-in-go-e88af0abd08d
