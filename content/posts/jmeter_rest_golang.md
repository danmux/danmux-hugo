---
title: "Jmeter is a Poor Choice for REST and Golang"
description: "Its surprising that Jmeter is clunky when it comes to RESTful based API testing."
date: "2014-07-06"
categories: 
    - "tech"
    - "jmeter"
    - "golang"
    - "testing"
    - "performance"
---

Its essential we have good integration tests and performance tests on our restful api, particularly now that many of the moving parts will migrate to microservices written in Go.

Trying to use Jmeter both to validate responses and apply reasonable load has been troublesome.

<!--more-->

## JSON Requests are Hard

Jmeter is crap at constructing anything but static JSON.

Just a quick search uncovers [Beanshell](http://beanshell.org/) - and it looks like hassle.

[This alternative](http://www.ubik-ingenierie.com/blog/extract-JSON-content-efficiently-with-jmeter-using-JSON-path-syntax-with-ubik-load-pack/) doesn't look too much better.


## JSON Responses are Hard

Applying meaningful assertions to the responses in Jmeter is also a bit of a ball ache - regex is the default. There is no json parsing out of the box.

Plugins improve things somewhat, but modelling a flow of a couple of requests with some shared session awareness is another load of hassle.


## Performance is Questionable

My own experiments and the thread below shows how a single jmeter instance is probably not quick enough to test the performance of a Go based web server, without setting up a few instances - but i spose we will have to do that in production even if we do find a fast tool.

<!-- Place this tag in your head or just before your close body tag. -->
<script type="text/javascript" src="https://apis.google.com/js/plusone.js"></script>

<!-- Place this tag where you want the widget to render. -->
<div class="g-post" data-href="https://plus.google.com/101114877505962271216/posts/PeZk8FY3PWY"></div>

## Alternatives

A quick search shows...

[Resty](https://code.google.com/p/restty/) - not tried it, not sure it does performance testing.

Others via stack overflow etc. all appear to be pretty much GUI based or GUI only

## Next Steps

Do we build our own tests?

Is there a more JSON friendly tool that can assert, pass on responses to the next step and load the service in parallel? I couldnt find anything.

[Fancy discussing on HN](https://news.ycombinator.com/item?id=7995111)
