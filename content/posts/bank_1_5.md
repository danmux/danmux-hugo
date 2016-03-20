---
title: "Bank 1.5"
description: "Consumer banking still struggles with its online identity."
date: "2011-09-23"
categories: 
    - "tech"
    - "testing"
    - "development"
    - "engineering"
    - "quality"
---

[The test pyramid](https://www.mountaingoatsoftware.com/blog/the-forgotten-layer-of-the-test-automation-pyramid) has its place - it gets across a simple idea, but it has been taken too literally, and applied innapropriately.

<!--more-->

The Pyramid and Dificult Compromises
-------------

Moi understanding of Cohns original article was that people too often downplay the service or integration tests, and I think discussions of the pyramid miss this emphasis, and have focused too much on the Pyramid itself.

Unit vs X Tests (where X = any name for anything other than Unit)  must be the most tiresome debate ever in the history of software development. Over the years I have often found myself encouraging and sometimes justifying my compromises. My arguments have never been devastatingly good, because like any engineering there are often many good enough ways to achieve the desired outcome.

'Good enough' is at the heart of all engineering. Good enough to manage the risks with just enough confidence. From past experience I have, in general, found that developers from a computer science background sometimes struggle with 'engineering compromise' more than those from an engineering background.

Judging good enough and risk is hard and takes confidence and experience to get the balance right in the many different projects and products we work within. A set of tried and tested prescriptions is always more straight forward. 

The Base of the Pyramid Problem
-------------------------------

It is way too prescriptive:

* It oversimplifies the situation, so it is used as a crutch to avoid critical thinking (notice I avoided the phrase containing the word 'cargo'). How often would 70:20:10 be perfect for your situation. 
* The top of the pyramid refers explicitly to GUI based tests, often irrelevant, often the people who mention it forget this. (However, even a testing trapezoid would still be wrong)
* True user testing is completely ignored - experience - emotion - engagement (Dare I say quality as [Persig](https://en.wikipedia.org/wiki/Zen_and_the_Art_of_Motorcycle_Maintenance) defines it "Quality is the knife-edge of experience, found only in the present, known or at least potentially accessible to all of us")
* It is based on the false assumption that integration tests are slow and flaky, they can be as fast and robust as unit tests.
* It is 4 years old - the world has turned. Modern GUI's (think Mobile and [SPA](https://en.wikipedia.org/wiki/Single-page_application)) are almost always cleanly divided by a well defined API. (if not then it's not a testing problem, but a design problem)
* It downplays the relevance of boundary interactions, particularly during concurrency.
* Unit tests are subject to availability bias - and take on an inflated importance because of it.
* Unit tests themselves are a kind of availability bias - because they are easy to write and run they are used in the place of a more difficult, more complex, more valuable strategy, one that only becomes clear when you critically assess the real risks affecting quality. 

An Illustration of Broken Thinking
----------------------------------
The article most cited to me is the well read [Test Pyramid](http://martinfowler.com/bliki/TestPyramid.html) by Martin Fowler which references the original Mike Cohn article, [The Forgotten Layer of the Test Automation Pyramid](https://www.mountaingoatsoftware.com/blog/the-forgotten-layer-of-the-test-automation-pyramid) 

Martins post links to another well known supporting article from the Google Testing Blog with the deliberately captivating title...

[Just Say No to More End-to-End Tests](http://googletesting.blogspot.ch/2015/04/just-say-no-to-more-end-to-end-tests.html)

Which captures the narrow experiences and environment of one setup with some fundamental problems. Some parts of that Google testing blog post I simply did not fully understand until a few re-reads, and even then I was left questioning the validity of what I read. I felt it sounded outdated, and was based on some strange assumptions, it appears I am not alone.

A counterpoint is posted on Martin’s post, which is also worth reading: [Making End-to-End Tests Work](https://www.symphonious.net/2015/04/30/making-end-to-end-tests-work/) which succinctly points out some limitation of the test system mentioned in the Google testing blog post above...

> "If your idea of fantastic test infrastructure starts with the words “every night” and ends with an email being sent you’re doomed"

The Google post is further analytically [dismembered here](http://bryanpendleton.blogspot.com.au/2015/04/on-testing-strategies-and-end-to-end.html?m=1), and this article captures my own confusion well...

> "Whatever went wrong with this project, though, one thing is very clear to me:
>
> The testing strategy is not the problem here."

Detail of Availability Bias
---------------------------------
(my own pseudo science - feel free to skip to next section (but there is mention of interesting work in any case))

Recently I have been reading a fascinating book that has long been on my to-read list: [Kahneman's Thinking Fast and Slow](http://www.amazon.co.uk/Thinking-Fast-Slow-Daniel-Kahneman/dp/0141033576). Fortuitously Kahneman has gifted me another tool in my arsenal against the over valuing of unit tests. Availability bias.

The over importance of unit testing may just be down to availability bias. 

When I analyse the past encounters over the issue of unit testing vs other types of testing, I suspect this bias (or availability heuristic as it is also known) may be the main factor that tips the balance in favour of unit tests and has resulted in them appearing to be so important.

Whilst I can't do the book justice here, I'll try and capture the salient points. Availability bias is an example of many evolutionary shortcuts in the human mind that can substitute a simple emotional decision for an otherwise difficult complex one, and has been demonstrated in many, brilliant, simple psychological experiments.

In particular, according to Kahneman, Norbert Schwarz showed the paradox that we are less confident in our decision when asked to come up with more reasons why it is a good decision. Our brains make an overconfident emotional decision when we can immediately produce a few supporting reasons, and a less confident decision, when forced to think harder. This is one of those paradoxes which are obvious when pointed out.

The book goes onto describe some related effects, regarding risk, which triggered my connection with testing, because I think much more in terms of risk than of test type, or test metric. Research by Paul Slovic, Sarah Lichtenstein and Baruch Fischhoff Showed that scientists opinions on the benefits of a particular technology could be increased by downplaying the risks, and similarly that the perception of the risks of a technology would be decreased, just by describing the benefits.

It is clear how this emotional bias can be applied to software testing: 

1. A lack of unit tests most often described as (and can be) a risk - therefore the benefits of unit tests are exaggerated in our minds.
2. We can quickly recall a few examples of people who repeat the advantages of unit tests, and quickly recall a few personal experiences that support their advantage so the risks of focussing on them are downplayed.

This is all based in normal animal survival laziness - it takes energy for our deep thinking brain to engage.

A more direct consequence of our natural laziness also has a different type of biasing effect on the increased value attributed to unit testing. 

It is hard for anyone deeply embedded in a development team to truly know how valuable the team is, or, where necessary, how valuable individual members of the team are, or even how good the product is. The situation has developed where the very people responsible for the success of a development team are some of the least likely to be able to make an objective decision. Under these conditions it is easy to see how 'measurement' is needed, it is this natural laziness that allows easy measurements to become so influential.

Unit tests are easy, and the quick visibility they afford in the form of the second most harmfull metric in development - coverage, creates something measurable. This is an easy metric, and when coupled with the biases of 1 and 2 above it is easy to see how unit testing and test coverage take on a an overly exaggerated value. Perhaps most often outside of the team who writes the tests.

I also think it is likely that we are still suffering the rebound from 10 years ago when test automation and unit testing were much less an integral part of the development cycle. The easiest goto tool in the interim has been the unit test.

The Smart Way
-------------

Whilst this post actually started out as a document of my linking availability bias with the over emphasis on unit testing and was not meant to be about the value of various tests in practice, it is probably clear that I encourage a more individually considered approach, than a set of prescribed rules.

I want to make it clear that I totally understand that there are loads of cases, particularly on the computer science focussed side of development where unit tests are essentially the only sensible test strategy. But they are only one part of the package, and often a small part.

In this article: [Why the Testing Pyramid is Misleading](http://www.joecolantonio.com/2015/12/09/why-the-testing-pyramid-is-misleading-think-scales/) the author discusses Todd Gardners (TrackJs) views which could be concieved to be indirectly addressing availability bias of unit tests, by making the point that we tend to not think in terms of risk. This article resonates strongly with my own approach to quality (again where quality is typically a bigger thing than that which we normally test for)

That article references the video offering very sensible advice from Todd Gardner, Software Engineer and Entrepreneur at TrackJS, namely think critically...

[Case Studies in Terrible Testing](https://vimeo.com/144684986)
(slide deck [here](http://www.slideshare.net/todd3091/case-studies-in-terrible-testing))

Especially salient is the advice from 25 minutes on, though this picture of the relative importance of different testing to mitigate the scale of the risks to the success of one particular project should illustrate the main thrust of the presentation...

![testing scales](https://lh3.googleusercontent.com/-uyf9z1SiSgw/Vnft1EYuFaI/AAAAAAAAMlQ/j_PQhbHL-jI/w1167-d-h870-p-rw "Testing Scales")
_Figure 1. A custom set of sliding scales of differing test strategies_
 
When we truly think about how we can best reduce the risk to success, and use alternative appropriate tools to manage that risk we start to focus more on delivering real value. 

Monitoring, CD, and a real understanding of the users perceived value or quality through canary code, A/B testing and the like will have a much bigger impact than 70% coverage.

> "Fast to fix is almost as good as never broken (and sometimes better)" 

_(from Todd’s slides above)_

The ‘scales of risk’ metaphor is much more intelligent and appropriate than the pyramid. The pyramid is one combination of the risk scales which may well align with a correct assignment of risk in a minority of real world cases (as would the ice cream cone). Though as far as I remember, not on any projects I have worked on.


