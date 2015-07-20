---
date: "2015-07-20T17:02:07+01:00"
title: "Queues Are Not Always The Answer"
description: "Is it really so bad to optimise early by choosing a tighter encoding early."
categories:
    - "tech"
    - "queues"
    - "desing"
---

I feel developers reach for the queue all too quickly. I’m talking about stand alone message queue services like RabbitMQ, ActiveMQ etc. etc. Not an in memory data structure (which in fact can be all you need sometimes)

<!--more-->

## Sync or Queue

Often the reasons for choosing a queue include many features that any synchronous service could provide.

There are, in my opinion, only four compelling use cases for a queue over a simple remote service.

This [blog post from Iron.io](http://blog.iron.io/2012/12/top-10-uses-for-message-queue.html) is a good example of an article that whilst explaining why queues are good never mentions where a simple service would suffice. 

All the reasons they mention as to why queues are good are correct. I just want to remind (myself mainly, so I can reference it when I get challenged again as to why I have not suggested a queue) that most of the advantages they mention are not exclusive to queues. 

**Decoupling** - Provides consistent interface, but then so do flexible payload interfaces. Admittedly a queue can allow the freedom to attach an number of subscribers without the publisher needing any further knowledge.

**Redundancy** - not really redundancy, but retry in upstream failure. Retry is a great feature, which is also often present in other interfaces. Often queues are hard to cluster for redundancy. 

**Scaleability** - Queues support scaleability because of decoupling and ease of adding pub sub nodes to the queue - just the same effect as the decoupling delivered by an interface to any other out of process service.

**Elasticity & Spikability** - the major use case  for queues - impedance miss-match

**Resiliency** - This argument assumes ALL processing nodes fail, otherwise it would just be the same behaviour as dealing with spikes (lack of cycles to process the work). This approach to ‘resiliency’ does not handle the case where the user is expecting feedback from the result of the request, then the queue lets them down.

**Delivery Guarantees** - is not an advantage of queues - it needs to be stated because it feels like queues may not deliver messages - a direct interface, being synchronous, has an implicit delivery ‘guarantee’

**Ordering Guarantees** - this can be a a valid use case if multiple processes feed into a queue, and assuming the queue itself is cluster-able, and has cross cluster atomic behaviour, then a queue, could be made to act like a temporal ordering funnel, in some way collating requests in a particular order (within an agreed time window). In any distributed system at scale, the clients of the queue are unlikely to deliver messages to the queue in a temporally deterministic order thereby undermining the queue value for this feature. Event ordering across a distributed cluster is THE hard problem in distributed systems, queues may inject some order into some parts of the system.

**Buffering** - a duplication of the Elasticity feature, and again a way of dealing with impedance mismatch - if any consuming processes behind an interface can operate in parallel and at different lifetimes then the same effect is achievable without a queue.

**Understanding data flow** - an argument for visibility, also solved by a consistent set of api’s in a system that are able to be reported on for statistics.

**Asynchronous Communication** - probably the single biggest most valuable use case - if the instigating process does not care directly about the outcome, or can check the outcome later, and the system still needs some eventually consistent guarantees then the queue is a fine choice. example use cases: event based features like triggering an email, or streaming into a reporting system

I think the other major feature, not mentioned in the article, is that typically queues offer... 

**Rich Routing or Distribution Model** - Fan out, topics etc etc. 

**Security**
The other aspect of a queue that is not mentioned in that article is security. A subscriber makes a connection outbound to the queue, so an internally secure zone in your network can avoid allowing any inbound connections, and only ever connects out to the queue, which could be considered more secure. Others may argue that all this does is change the attack vector slightly to focus on compromising the queue itself, and is no more secure than available attacks on the network layers, I don’t know enough about that, but it feels more secure to me not allowing inbound connections. 

**Bad UX**
Using a queue in any part of your architecture that affects the user experience is a risky compromise, which can often mask the underlying problem.

## Summary
In summary use a queue, if you really don’t mind asynchronous or parallel work...

1. If you truly have no control over, or don’t care about the upstream service response times.
2. If you have spikes in your requests, and you can’t cost effectively keep to a reasonable response guarantee, and you know you have enough spare capacity in the off peak to catch up.
3. If you want to do some funky distribution of your messages (typically to many observers) particularly if you are not bothered about the result of any upstream subscribers work  
4. If you agree that the pull model from a secure zone to a queue adds to your network security.
