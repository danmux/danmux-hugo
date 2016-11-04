---
title: "What Golang Is and Is Not"
description: "Like all of us Golang has its problems, but sometimes those problems are creations of our own expectations."
date: "2016-08-17"
categories: 
    - "golang"
    - "go"
    - "tech"
    - "testing"
    - "development"
    - "engineering"
    - "quality"
---

We are all products of our own histories, and I suspect many routes to Go have been made less enjoyable by misguided expectations. The journey from when a budding developer first ‘hello worlded’ to now may have made Go’s more subtle strengths less obvious to them.

<!--more-->

Go is least of all about the language in and of itself, but rather about the broader things affected by it, more so than other languages I have used over the years.

Many recent arrivals to Go have preconceptions that have been proved wrong and caused disappointment, this appears to happen most often when looking from a more purist computer science based language design point of view.

Go is an engineering tool, in a much broader sense. To appreciate it properly I think you have to have spent a decent amount of time responsible for the full lifecycle. If all you ever do is write and commit code then much of Go will be lost on you.

The Go Hype, or Not.
--------------------
I have heard more than once that people have been disappointed by go given the hype surrounding it. I don’t remember anything really hype like 5 years ago when go was in its infancy, and I must say that I am not really aware of anything specific now. Of course any new language that is becoming better adopted, in particular as quickly as Go, will naturally attract attention.

There are articles out there that sing Go’s praises, perhaps a little too highly, but I don’t remember many of them explicitly making any claims about the language itself being the reason why Go is so good.

Go is Not An Innovative Language
--------------------------------
“Innovative” is such an overused and abused word that it has lost a lot of power and meaning. All innovation is contextual and to use the word without context is foolhardy. In the context of language design Go was never an innovative language, nor was it presented as such, or anyone dishonest in representing it that way.

As a language Go was always explicitly a return to simplicity, and in many ways naivety, for sound reasons.

“There is nothing new under the sun” rings true in all languages since the 80’s. Virtually everything we see in language design now that someone says is “innovative” has been explored in some form before. Go is certainly no exception, but remember it never claimed to be state of the art.

Regarding the language being youthful, of course it is, but the intention is not for the language itself to ‘mature’: no more complexity is going to be added, or at least it’s very unlikely. It is not ‘missing’ comprehensions, or inheritance, or generics, they are **omitted** (and I pray, always will be). In some way, in the context of the current fashion of returning to more functional languages, or the evolution of good old languages to include more functional paradigms (I’m looking at you Javascript and Python for two examples) then in a tenuous convoluted way Go has ‘innovated’ by avoiding that trend.

Go is an Innovative Thing
-------------------------
It is hard to define what the ‘thing’ is, but it is quite a broad thing, I can’t fully say it is an approach, or a belief, or even ‘patterns and practices’, though that last phrase feels closest. This is still about the best read on the subject: [Go at Google: Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article)

Possibly Go’s greatest ‘innovation’ is to eschew making software engineering an overly academic process in daily practice, and to focus on improving the tools, speed, reliability and pure pleasure in **delivering** and running things of value.

Thats not to say that Go encourages you hack away without a sound foundation, quite the opposite, I think Go almost *requires* you to have a good grasp of the fundamentals, to be an effective Go programmer. I would say that the simplicity and imperative style that Go encourages, tends to demand a greater underlying computer science knowledge than many other languages may expect. Having recently watched [Mind the Gap (GopherCon 2016)](https://www.youtube.com/watch?v=ClPIeuL9HnI) I think Katrina Owen echo’s this belief.

### Unifying
I think one reasonable way of categorising all the things that make up delivering software products is into: *human, operational and technology* factors, in that order of value. Go helps address some problems in all three areas but it is its influence over human and operational factors where it sets itself apart from other systems.

Even as technologists we can’t help moving things into the ‘human’ domains such as emotion and personalisation, most compilers don’t care about many of the things humans care about, but the language naturally becomes a very human thing.

Go has learned from the experience of fractured communities and continual in-fighting amongst teams and has attempted to avoid debates that continue to rage in other languages that are 20+ years old. The early focus on idioms helped that. This approach comes from years of experience delivering in teams at scale, where the language is one small factor, which has caused an inappropriate amount of time wasted on a tiny fraction of the whole value. As an example: curly brace positioning is one of the most trivial things possible, and yet still many hours are wasted on it.

### Paradox of Choice
In all languages there are always some basic primitives and data types that relate closely to the machine instruction set, which in turns map well to the hardware (see [Note 1.](/posts/what_golang_isnt#note-1)) Ultimately all other higher order data structures in all languages are composed of arrays, references, and structs. Trees, heaps, sets, queues and everything else effectively only manipulate arrays of structs/primitives or self referencing structs, thats it, simple, or it should be.

In the Go language at its heart that simplicity is encouraged, we are only offered some basics. To start with we are given primitives, structs and arrays, then because it is unavoidably useful we have a dynamic array, a `Slice` which is an embellished array to allow dynamic resizing. Finally, in certain problem domains the power and flexibility of a hash-map is also unavoidable, therefore Go provides a `Map` built in. These are given some special treatment, simply because it is very useful and they are special, (providing the same treatment to function returns would add low value complexity)

There is a subtlety to providing only this subset. The fact they map well to lower layers imbues an immediate sense of being more intimate with the CPU, which certainly for ‘older’ engineers feels refreshing, and at a minimum for younger engineers tends to influence design decisions towards simplicity. Having said that the desire to construct complex implementations of data structures - even if used only once - is ever present when a new arrival to Go finds their favourite container is ‘missing’.

To provide any solution in Go that needs a dynamic data structure you can choose between hand rolled linked structures or a `Slice` or `Map` (or compose with them). As they are quite different the choice is normally obvious. Contrast this to the choice between map, set, hashset, bag etc etc, or rolling your own in a language that makes this a lot harder. Often the author actually only uses a subset of the functionality of those data structures. In these cases the choice becomes much less simple, indeed often a point of confusion and contention and can be the cause of further low-value conversations.

A Go programmer takes a slice or map and mixes in a few functions to provide the structure they need. For example The Go standard library has provided a minimal `container` package with a `heap` (which is just an interface), a `list` (doubly linked), and a `ring` (which is a closed doubly linked list). To implement a heap - you need to provide a builtin to implement the storage with an array being the typical choice.

Each one of those packages has no more than around 200 lines of code. Those few lines of code are very readable, the behaviour understandable, and the performance predictable; being a function of the performance of the well understood builtin and the users own implementation code.

This removal of choice and focus on reusing the two builtins, drives a readability, clarity and consistency amongst Go programmers, not afforded in other languages.

In other languages an iterator (one of the often complained about omissions) necessarily abstracts that which is being contained, and often insists on a broad interface some of which then remains unused, and adds some cognitive load, sometimes unnecessarily. Not providing an iterator and not providing many containers implementing iterable, or whatever other system, avoids needing a whole swathe of knowledge, indirection, discussion and misunderstanding, at little practical cost.

The effect of the omission in real world code results in minimal extra work at code creation time, for great gains in the rest of the lifecycle. Custom data structures can be composed from the well understood builtins, rolled in under 100 lines of code and can can exist close to the place they are used (yes repeated!). The effect of this approach on readability, maintainability, decoupling, removing seemingly endless low value conversations, and when push comes to shove the ability to understand performance characteristics and then tune them, adds so much more value to the whole lifecycle, than the cost of the omission.

These are subtle yet important factors that attempt to address some of the human and operational complications.

More Than a Language
--------------------
Without going into all of the things that the Go ecosystem brings to the table on top of the language design, what should be clear from a shallow familiarity with the tooling is that Go has focussed on providing answers to some of the more difficult aspects of actually getting code that is both stable and agile into production. These things that were part of Go from the beginning have had to evolve over decades in other languages, often in fractured directions - again adding the paradox of choice. For example Go’s dependency management attempts to solve a thorny problem, and while `go get` in particular is going through some teething pains, its inclusion from day one is illustrative of Go’s intention.

It is this focus on the operational aspects of development, so early on in Go’s evolution that emphasises the reason Go was created, and is commonly overlooked in favour of low value critiques of the language itself.

* * *

Note 1.
-------
In our Von Neumann / Harvard world we have basically three hardwired data structures, a register, a stack, and addressable memory, these are mapped via the instruction set to: register operations; effectively push and pop, moving data to and from memory addresses, and in CISC’s contiguous memory operations, even in RISC’s loop primitives are optimised for contiguous ranges. Ultimately these in turn map through compilers to value variables (which indirect the decision to use registers, the stack, or an addressable value), and reference variables, which contain the value of an address of the value, and slightly higher up the conceptual scale: arrays. Compilers also compose these fundamental variables into primitive data types: ints, floats, and arrays into strings etc. Finally, also through well managed contiguous memory, namely ‘packing’, common primitives are grouped into ‘structs’. There are subtle variations on these themes, particularly with respect to structs or objects, involving further indirection (think v-tables etc), but that is the crux of it.


