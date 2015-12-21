The Test Pyrimid and Availability Bias
======================================

This [blog post from Iron.io](http://blog.iron.io/2012/12/top-10-uses-for-message-queue.html) 

[The test pyramid](https://www.mountaingoatsoftware.com/blog/the-forgotten-layer-of-the-test-automation-pyramid) has its place - it gets across a simple idea, but it has been take to be too literally, and applied innapropriately.

My understanding of Cohns original article was that people too often downplay the service or integration tests, and I think the pyramid itself has taken on too much of its own character.

Unit vs X Tests (where X = any name for anything other than Unit)  must be the most tiresome debate ever in the history of software developments. Over the years I have often found myself encouraging and sometimes justifying my compromises. My arguments have never been devestatingly good, because like any engineering there are often many good enough ways to achieve the desired outcome.

Good enough in my opinion IS engineering. From past experience I have, in general, found that developers from a computer science background sometimes struggle with 'engineering compromise' more than those from an engineering background.


The Meat of the Pyramid Problem
-------------------------------

* It over simplifies the situation, so it is used as a crutch to avoid critical thinking (notice i avoided the phrase containing the word 'cargo'). How often would 70/20/10 be perfect for your situation. 
* The top of the pyramid refers explicitly GUI based tests, often irrelevant, often the people who mention it forget this. (Even a testing trapezoid would still be wrong.
* True user testing is completely ignored - experience - emotion - engagement (Dare I say quality as Persig defines it "Quality is the knife-edge of experience, found only in the present, known or at least potentially accessible to all of us)
* It is based on the false assumption that integration tests are slow and flaky, they can be as fast and robust as unit tests.
* It is 4 years old - the world has turned. Modern GUI's (think Mobile and SPA) are almost always cleanly divided by a well defined API. (if not then its not a testing problem, but a design problem)
* It down-plays the relevance of boundary interractions, particularly during concurrency.
* Unit tests are subject to availability bias - and take on an inflated importance because of it.
* Unit tests themselves are a kind of availability bias - because they are easy to write and run they are used in the place of a more dificult, more complex, more valuable strategy, one that only becomes clear when you critically assess the real risks affecting quality. 

An Illustration of Broken Thinking
----------------------------------
The most cited to me is the well read [article](http://martinfowler.com/bliki/TestPyramid.html) by Martin Fowler that references the original Mike Cohn article, [The Forgotten Layer of the Test Automation Pyramid](https://www.mountaingoatsoftware.com/blog/the-forgotten-layer-of-the-test-automation-pyramid) 

Martins post links to another well known supporting article from the Google Testing Blog with the deliberately captivating title...

[Just Say No to More End-to-End Tests](http://googletesting.blogspot.ch/2015/04/just-say-no-to-more-end-to-end-tests.html)

Which captures the narrow experiences and environment of one setup, with some fundamental problems. Some parts of that Google testing blog post I simply did not fully understand until a few re-reads, and even then I was left questioning the validity of what I read, I felt it sounded outdated, and was based on some strange assumptions, it appears I am not alone.

A counter point is posted [here](https://www.symphonious.net/2015/04/30/making-end-to-end-tests-work/) which succinctly points out some limitation of the system mentioned in the blog post above...

"If your idea of fantastic test infrastructure starts with the words “every night” and ends with an email being sent you’re doomed"

The Google post is further analytically dismembered here, and this article captures my own confusion well...

"Whatever went wrong with this project, though, one thing is very clear to me:

The testing strategy is not the problem here."

http://bryanpendleton.blogspot.com.au/2015/04/on-testing-strategies-and-end-to-end.html?m=1


Detail of Availability Bias - 
----------------------------------------------------
(my own pseudo science - feel free to skip to next section (but there is mention of interesting work in any case))

Recently I have been reading a book that has long been on my to-read list: Kahneman's Thinking Fast and Slow. Fortuitously Kahneman has gifted me another tool in my arsenal against the over valuing of unit tests. Availability bias.

The over importannce of unit testing may just be down to availability bias. 

When I analyse the past encounters over the issue of unit testing vs other types of testing, I suspect this bias (or availability heuristic as it is also known) may be the main factor that tips the balance in favour of unit tests and has resulted in them appearing to be so important.

Whilst I can't do the book justice here, I'll try and capture the salient points. Availability bias is an example of many evolutionary shortcuts in the human mind that can substitute a simple emotional decision for an otherwise difficult complex one, and has been demonstrated in many, brilliant, simple psycological experiments.

In particular, according to Kahneman, Norbert Schwarz showed the paradox that we are less confident in our decision when asked to come up with more reasons why it is a good decision. Our brains make an overconfident emotional decision when we can immediately produce a few supporting reasons, and a less confident decision, when forced to think harder. This is one of those paradoxes which are obvious when pointed out.

Th book goes onto describe some related effects, regarding risk, which triggered my connection with testing, because I think much more in terms of risk than of test type, or test metric. Research by Paul Slovic, Sarah Lichtenstein and Baruch Fischhoff. Showed that scientists opinions on the benefits of a particular technology could be increased by down playing the risks, and similarly that the perception of the risks of a technology would be decreased, just by describing the benefits.

It is clear how this emotional bias can be applied to software tesing: 

1. A lack of unit tests is considered to be (and can be) a risk - therefore the benefits of unit tests are exaggerated in our minds.
2. We can quickly recall a few examples of people who repeat the advantages of unit tests, and quickly recalled personal experiences that support their advantage so their value are exaggerated.

This is all based in normal animal survival lasiness - it takes energy for our deep thinking brain to engage.

A more direct consequence of our natural lasiness also has a different type of biasing effect on the increased value attributed to unit testing. 

It is hard for anyone deeply embedded in a development team to trully know how valuable the team is, or, where necessary, how valuable individual members of the team are, or even how good the product is. The situation has developed where the very people responsible for the success of a development team are some of the least likely to be able to make an objective decision. Under these conditions it is easy to see how 'measurement' is needed, it is this natural lasiness that allows easy measurements to become so influential.

Unit tests are easy, and the quick visibility they afford in the form of the second most harmfull metric in development - coverage, creates a measurable metric. This is an easy metric, and when coupled with the biases of 1 and 2 unit testing and test coverage take on a an overly exaggerated value. Perhaps most of all outside of the team who writes the tests.

I think it is likely that we are still suffering the rebound from 10 years ago when test automation and unit testing were much less an integral part of the development cycle - where the easiest goto tool was unit tests.


The Smart Way
-------------

An article that discusses indirectly the availability bias of unit tests, by making the point that we tend to not think in terms of risk, resonates strongly with my own approach to quality (again where Quality is a much bigger thing than that which we normally test for)

http://www.joecolantonio.com/2015/12/09/why-the-testing-pyramid-is-misleading-think-scales/

Which links to a video offering very sensible advice from Todd Gardner, Software Engineer and Entrepreneur at TrackJS, namely think critically...

Case study in terrible testing. https://vimeo.com/144684986 - 

Especially salient is the advice from 25 minutes on. "Fast to fix is almost as good as never broken" 

Monitoring, CD, and a real understanding of the users percieved value - or quality can be the real risk mitigators.

In my opinion the scales of risk metaphore is much more intelligent, and appropriate.

https://lh3.googleusercontent.com/WKD1prA9idbziR7HJrTwqzJZe5KHCu4_YXNXBTIvG5WUIP1DoiiMEZqso8AOPEJyof4r=w6337-h3961-rw-no








Whilst this post actually started out as a document of my linking Availability Bias with the over emphasis on unit testing and was not meant to be about the value of various tests in practice, it is probably clear that I encourage a more individually considered pragmatic approach, than a set of prescribed rules.

There are many cases, particualrly on the computer science side where unit tests are essentially the only sensible test strategy. Then there are probably more cases in the average commercial environment where integration tests comparing the domain of user interractions with the range of expected responses can have considerably more valuable.

So just to confirm my general position I would choose a set of integration tests that are likely to traverse a multitude of services validating the vaguaries of infrastructure, 


