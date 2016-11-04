---
title: "The Cult of Go Test"
description: "You are using plain Go test then you have drunk too much Go Kool-aid."
date: "2016-10-30"
categories: 
    - "golang"
    - "go"
    - "tech"
    - "testing"
    - "development"
    - "engineering"
    - "quality"
---

A favourite test helper library, with some simple test assertion functions clearly has some value. But this post puts forward some useable concrete arguments why they are normally just not worth it.

<!--more-->

After five years with Go and the last three in a (now) 100% Go team I've worked with around 30 Go developers - not a huge amount, but not insignificant. The one thing that new people to the team challenge most is the lack of their favorite test helpers. People generally are OK that we don't need a whole framework, but the small simple assertions? Why are they so bad?  

tl;dr
-----

They are not so bad, but they come at a cost, defer to avoid them.

Three Reactions
---------------

This came up again recently and as I prepared to put forward the case against assert libs a learned colleague reminded me that this is still a case of bike-shedding. So if like us you have bigger problems, park this, and focus on them. However I'm putting pen to paper so I can point people at it in future, and it may help others. At first this recent tweet from @KentBeck first felt relevant:

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Don&#39;t spend more time discussing a reversible change than it would take to make (&amp; potentially reverse) the change</p>&mdash; Kent Beck (@KentBeck) <a href="https://twitter.com/KentBeck/status/792911449249026048">October 31, 2016</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

But then actually once you have `assert.Equals` dotted throughout the tests the cost of removing it becomes unaffordable. Once you commit to one of the assertion libs it becomes hard to reverse that change. In that respect it is an expensive decision.

Perhaps it was somewhat lucky that when I started writing Go code that no assertion libs existed, so I was forced to learn to live without them (I had been working with python and Java most recently prior to that) I felt the same surprise then as people do now! Years ago when `testify` was young, I immediately started using it. It was only then as I had to learn how it worked, and work round its bugs and surprises, that I first came to think that perhaps my last 10+ years of unit testing had been overly complicated. (I expect there are many developers who don't remember any of the xUnits existing - when each new project needed time to write the test harness or at least copying in some previous [seedwork](http://www.martinfowler.com/bliki/Seedwork.html) as Fowler called it. At the time jUnit felt like a godsend.

This repeating conversation with engineers new to the team (or new to the stdlib only approach) could be grouped into one of three catagories of response:

1. **Interesting** - I've always felt a bit uneasy about the need for all this extra stuff - lets give this a go.

2. **Really** - I'm pretty sure its going to be worse without the test lib/framework (dsl) im used to - but lets see if im right.

3. **What!** - thats just ignoring all the other testing stuff that happened in other things and is almost certainly a retrograde step.

There is a subtly different class of response for people who have been using Go for a while and have already chosen one of the test helper libraries, and have not yet *tried* plain go test. When I say *'tried'* I mean actually stuck with it for a few months - not wrote some plain tests in the first few weeks and immediately added a familiar looking assertion lib. 
   
In all three cases most people will give it some open minded time and form an opinion based on their now real world experience. In general people start to at least get comfortable with, if not actually value, the approach. During this time when a bit of effort is expended on learning a few of these tips the adoption is improved.

In the worse case even after working with plain go test, some developers really struggle to understand or accept why they should not use their choice of helpers, occasionally attempting to sway opinion by challenging the intelligence and integrity of the team with accusations of cult like behaviour (cargo or otherwise). "Kool-aid" gets mentioned more than once. It will take a greater depth of argument than that to challenge this new (or retrograde, depending on your stance) well considered approach. From having spoken to team members and other teams and at conferences and from lots of stuff online it is clear that I am not alone in having thought about this a lot.

This poll is possibly not great quality, and almost certainly somewhat self selecting but illustrates that at least some folk also prefer plain go test according to this [Straw Poll](http://www.strawpoll.me/1716206/r)

<iframe src="http://www.strawpoll.me/embed_1/1716206/r" style="width:680px;height:541px;border:0;">Loading poll...</iframe>

Not Such a Big thing 
--------------------
(but it feels like it is fundamental!)

It really isn't such a big deal in any case, either way is OK, each team should make the call, but once it has been made it should be kept consistent. The mix of both plain go and assertions libs - clearly dilutes the benefits of consistency, particularly if multiple helper libs are added.

Without doubt assertions can reduce verbosity in the tests, but as the code samples below will demonstrate when the stdlib tests are well written there is not that much in it. Clearly typing effort is not the final arbiter (by any means) in assessing the best approach, but it an important factor. There are other good, if subtle, reasons to stick with plain go test even if it may take more effort. 

Before looking at some specific examples of how stdlib tests compete well with assertion libraries (on various factors including effort) here are the bare minimum things to have read that start to explain the position.

https://golang.org/doc/faq#testing_framework

https://golang.org/doc/faq#assertions

It's worth hearing the reason for Blake Mizerany (of Sinatra fame) to do this..

![Blake sees the light](/img/just-use-stdlib.png "just use stdlib")

Explained in this presentation...

<iframe width="560" height="315" src="https://www.youtube.com/embed/yi5A3cK1LNA" frameborder="0" allowfullscreen></iframe>

Real World Example
------------------

Here is some real world test code that tends to favour an assertion library as there are only two fixtures to test a single function.

The test was initially written with [testify](https://github.com/stretchr/testify) :

<script src="https://gist.github.com/danmux/cbe74e643538bb0ff9c2bf78511d630e.js"></script>

...A compact 25 lines and 529 chars typed.

This was then rewritten in plain go test in a very imperative style...

<script src="https://gist.github.com/danmux/64b4c47b2b25676adb7b3c18ac6193ac.js"></script>

...41 lines, but only 646 Chars typed - just 120 chars more than the assert lib.

Then even though there are only two fixtures a table test and deep equals was tried...

<script src="https://gist.github.com/danmux/8e4e727ac36dcada592c882cd2384e9a.js"></script>

...35 lines and 643 chars typed - yay we saved 3 chars! of course the table approach starts to pay off with more fixtures.

The same test with a small local helper assert function...

<script src="https://gist.github.com/danmux/c805910f0727698581696f6e715843c0.js"></script> 

...Only 30 lines but more typing - the assert lib uses the correct line number, in this test our failures would all come from line 7 so we have to add the `valid` and `invalid` words so it is clear where the failure is from thereby adding to the character count. Hence: 708 chars typed, without them it would be closer to 650 chars.

You may end up writing or dare I say C&P-ing this and similar helpers many times for good readability, and with little harm.  

Written again with the comparisons factored into its own helper...

<script src="https://gist.github.com/danmux/e7fe6c833a784a94200f463ac197ca29.js"></script>

...Now it is down to 28 lines and 637 chars, the lowest line and char count of all the contenders.

In all 'none assertion' cases the number of characters needed is more, but trivially so - and are a constant offset, not linear, adding more fixtures does not grow the delta between assertion and none assertion (assuming that the assertion based test also migrates to using tables as well). The delta can grow as more things are being tested.

All the above tests are imperfect and can be pared down or improved but it is enough to provide a comparison to discuss.

Although as the number of fixtures increase the table and larger local helper may be the best approach for this simple case the implicit test is probably the best because the failures are reported at the line numbers of the `t.Error` so for the few extra chars the helpers are not worth it.  

The Implications
----------------     
An essential function of tests is to help document the thing under test, some approaches to testing can reduce this documentation effect, but in this case the assertion lib does not overly abstract or otherwise hide the real function under test - so as documentation it is ok.

It does need concerted effort to structure and write clear tests in go test to accomplish a comparable succinctness - but this is a good thing.  

The assertion style is (very slightly) more concise.

But at what cost.

### Indirection and Gotchas


The functional indirection is also present in our local test helper cases but the helper is *local* - its right there in the code and is a few very easily understood lines, and the arguments are typed.

There is another semantic indirection in the assertion lib, something of a mini DSL to learn:    

The `Equals` and `Nil`, `NotNil` are in another package and we make assumptions based on their naming

I think it is a mistake to remove type safety from a unit test (I feel *somewhat* differently about testing some across the wire API's)

You have to know and think about the `Equals` having type checking entirely removed. The following should never both pass, but they do:

{{< highlight go >}}
assert.Equal(t, iban.Bban, "123")
assert.Equal(t, iban.Bban, 123)
{{< /highlight >}}

You may well know of another equality assertion in your favorite library that does type checking as well - but I expect that is only at runtime. There is considerable value in failing during compilation (more on the DSL later).  

Pointer equality is another unnecessarily introduced gotcha:

e.g. this testify issue

https://github.com/stretchr/testify/issues/296

> Hi Rob,
>
> It should be applicable to all pointers.
> 
> `NotEqual` is comparing the values the pointers point to, rather than the
> pointer addresses.
> 
> We must make that clear in the docs.


This issue was closed with **clearer documentation**. So to understand how and when to use NotEqual you have to carefully read the documentation or inspect the code. The fundamentals of the issue captured more succinctly in this snippet.

https://play.golang.org/p/QA6WK4aNfA

Something naturally avoided in explicit comparison. 

These `Equals` and friends add another thing to learn for us and for every new engineer for ever more, and it is just another thing to be tripped up by, and then it does not solve the testing debate completely. In fact it creates its own new debate: the "I prefer x over y lib/framework" debate. Which DSL is best?    

### The (not so) Mini [DSL](https://en.wikipedia.org/wiki/Domain-specific_language)

Equality never looked so complicated. What follows are some public API functions from testify:

{{< highlight go >}}
// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) bool {
{{< /highlight >}}

(btw what does the comment "This function does no assertion of any kind." mean)

Then there are the following...

{{< highlight go >}}

func ObjectsAreEqualValues(expected, actual interface{}) bool {

func EqualValues(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {

func Exactly(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {

func Equal(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {

{{< /highlight >}}

and then why is this ... 

{{< highlight go >}}

func NotEqual(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
{{< /highlight >}}

different from ` !Equal(...` ?

Testify is one of the more simple libraries - with a smaller DSL to learn, but there is still a reasonable amount of explicit knowledge needed to not get tripped up. 

Take a look at another larger one... [GoConvey](http://goconvey.co/) who prodly announce 

> "Expressive DSL"

What will happen when you write:

`So(x, ShouldAlmostEqual, 2)`

or

`So(y, ShouldNotResemble, 2)`

Here are all the other assertions we need to learn to become effective...

{{< highlight go >}}
var (
	ShouldEqual          = assertions.ShouldEqual
	ShouldNotEqual       = assertions.ShouldNotEqual
	ShouldAlmostEqual    = assertions.ShouldAlmostEqual
	ShouldNotAlmostEqual = assertions.ShouldNotAlmostEqual
	ShouldResemble       = assertions.ShouldResemble
	ShouldNotResemble    = assertions.ShouldNotResemble
	ShouldPointTo        = assertions.ShouldPointTo
	ShouldNotPointTo     = assertions.ShouldNotPointTo
	ShouldBeNil          = assertions.ShouldBeNil
	ShouldNotBeNil       = assertions.ShouldNotBeNil
	ShouldBeTrue         = assertions.ShouldBeTrue
	ShouldBeFalse        = assertions.ShouldBeFalse
	ShouldBeZeroValue    = assertions.ShouldBeZeroValue

	ShouldBeGreaterThan          = assertions.ShouldBeGreaterThan
	ShouldBeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
	ShouldBeLessThan             = assertions.ShouldBeLessThan
	ShouldBeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
	ShouldBeBetween              = assertions.ShouldBeBetween
	ShouldNotBeBetween           = assertions.ShouldNotBeBetween
	ShouldBeBetweenOrEqual       = assertions.ShouldBeBetweenOrEqual
	ShouldNotBeBetweenOrEqual    = assertions.ShouldNotBeBetweenOrEqual

	ShouldContain       = assertions.ShouldContain
	ShouldNotContain    = assertions.ShouldNotContain
	ShouldContainKey    = assertions.ShouldContainKey
	ShouldNotContainKey = assertions.ShouldNotContainKey
	ShouldBeIn          = assertions.ShouldBeIn
	ShouldNotBeIn       = assertions.ShouldNotBeIn
	ShouldBeEmpty       = assertions.ShouldBeEmpty
	ShouldNotBeEmpty    = assertions.ShouldNotBeEmpty
	ShouldHaveLength    = assertions.ShouldHaveLength

	ShouldStartWith           = assertions.ShouldStartWith
	ShouldNotStartWith        = assertions.ShouldNotStartWith
	ShouldEndWith             = assertions.ShouldEndWith
	ShouldNotEndWith          = assertions.ShouldNotEndWith
	ShouldBeBlank             = assertions.ShouldBeBlank
	ShouldNotBeBlank          = assertions.ShouldNotBeBlank
	ShouldContainSubstring    = assertions.ShouldContainSubstring
	ShouldNotContainSubstring = assertions.ShouldNotContainSubstring

	ShouldPanic        = assertions.ShouldPanic
	ShouldNotPanic     = assertions.ShouldNotPanic
	ShouldPanicWith    = assertions.ShouldPanicWith
	ShouldNotPanicWith = assertions.ShouldNotPanicWith

	ShouldHaveSameTypeAs    = assertions.ShouldHaveSameTypeAs
	ShouldNotHaveSameTypeAs = assertions.ShouldNotHaveSameTypeAs
	ShouldImplement         = assertions.ShouldImplement
	ShouldNotImplement      = assertions.ShouldNotImplement

	ShouldHappenBefore         = assertions.ShouldHappenBefore
	ShouldHappenOnOrBefore     = assertions.ShouldHappenOnOrBefore
	ShouldHappenAfter          = assertions.ShouldHappenAfter
	ShouldHappenOnOrAfter      = assertions.ShouldHappenOnOrAfter
	ShouldHappenBetween        = assertions.ShouldHappenBetween
	ShouldHappenOnOrBetween    = assertions.ShouldHappenOnOrBetween
	ShouldNotHappenOnOrBetween = assertions.ShouldNotHappenOnOrBetween
	ShouldHappenWithin         = assertions.ShouldHappenWithin
	ShouldNotHappenWithin      = assertions.ShouldNotHappenWithin
	ShouldBeChronological      = assertions.ShouldBeChronological
{{< /highlight >}}

Dependencies don't come for free.

GoConvey only clocks in at 8300 lines, but its assertion package will introduce a further 26,000 lines (including a number of test libraries by Aaron Jacobs). 

The code extracted from testify just to support `Equals` is 323 lines and 8200 characters, the full package adds 14,000 lines and exports 80 functions.

One of the larger libraries used by the [Ginkgo](https://github.com/onsi/ginkgo) framework is [Gomega](https://onsi.github.io/gomega/) it has 1500 lines of documentation and just under 12,000 lines, Ginkgo has 20,000 lines - these counts don't include any other dependencies.

(for clarity: *"lines"* is lines in all go files using the naive `find . -name '*.go' | xargs wc -l` simply to get a sense of scale)

The one thing these libs have in common is bugs. It is annoying enough to have to debug test code, let alone 3rd party test support libraries. 

Summary
-------

Hopefully this post demonstrates that the value these assertion libs add is, at least, arguable and also details some of the complexities they add in exchange. As much as the arguments for an assertion lib are clear, perhaps this has helped tip the balance in favour of the stdlib approach.

* No doubt assertion libs can reduce typing and repetition.
* Reductions in typing with assertion libs are not profound.
* The necessary repetition without assertion libs can be minimised. 
* Some repetition locally can increase readability at no great cost.
* There is a burden of extra care in writing plain go tests - arguably a good thing.
* Type safety is often dropped - or needs explicit knowledge, or is only enforced at runtime.
* Pointers vs value comparisons can easily trip up the unwary.
* There is a DSL to learn - often considerable - an extra burden on the team and new members forever.
* The libs often introduce multiple ways to achieve the same thing, or worse: similar but subtly different ways.
* It does not resolve the infighting - a new dev will argue for their favourite lib.       
* Adding the assertion lib in the middle of a project adds an annoying inconsistency.
* Some assertion libs tend to reduce the value of tests as documentation.
* You are adding another (often large) library dependency with associated maintenance overhead, bugs, life cycle etc.  
* It is a commitment that will stay with you for a long time, and it is expensive to undo. 

I agree that not only is the discussion bike-shedding it appears to follow [Sayre's law](https://en.wikipedia.org/wiki/Sayre%27s_law) (both this and the law of triviality I think are forms of availability bias, which I have [written about in the context of unit tests](http://danmux.com/posts/test_pyramid_availability_bias/)) and we should stay focussed on the bigger challenges.

If I were to work with a team that has aligned on assertions in their testing, then I would, with some sadness, accept it and move on to things that really matter.

The only value I can see in accepting an assertion library into an established stdlib only unit tests codebase is that like us, you have probably wasted many person-days discussing this. We certainly redo the same discussion for many new team members. If only I was convinced adding one of these libs would end the discussion, then I might agree to it (for the wrong reasons).

**DRY** does not only apply to code. Next time I will point them at this, and hope it helps.

