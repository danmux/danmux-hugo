---
title: "JSON, Gzip, Snappy and Gob Across the Wire"
description: "Is it really so bad to optimise early by choosing a tighter encoding early."
date: "2014-09-21"
categories: 
    - "tech"
    - "gob"
    - "golang"
    - "JSON"
---

Coming from a background where memory and clock cycles were sparse, binary encodings have always held an appeal. Since then I’ve been told we have loads of compute power, ample cheap RAM and disk, and when the network is the bottleneck then, well, that is a good problem to have.

Its one of those ages old occasionally heated debates....

(*tl;dr* almost always use gzipped JSON)

<!--more-->

## Serialising for Wire and Disk

**Much better to use a more debuggable human readable encoding and compress it in flight during transmission or storage**, after all browsers have native support and the performance benefits of binary formats are negligible. This is the crux of the argument for those in favour of verbose text based serialisations.

I understand the argument and **this is almost always sound advice**, so why doesn’t Memcache do it, why do all NoSql and NewSql implementations offer a binary alternative if not *only* a binary inerface? Why does Thrift thrive, and MessagePack pack a punch :p ?

Is it faster, or is it just cooler to do some binary stuff? - ‘Intellectual masturbation’ as an old learned colleague delightfully coined it. Or is there really a compelling use case for these things.

Heres a [recent example of the debate](https://groups.google.com/forum/#!topic/golang-nuts/xtXh0yWOens) with all the typical (valid) points being trotted out

## In This Case...

Having just built a nicely homogenous #golang rpc microservices framework (or ‘milliservices’ as [Ian Davis called his](https://twitter.com/iand/status/510090898269831170)) I was naturally drawn to the native RPC gob encoding, but I hedged my bets and did the second worst of both worlds a gob based RPC envelope - with a []byte payload that could carry any serialisation - a la http. (the absolute worst would be http, with a gob body)

The point of this ‘neither one nor the other’ design was that much like HTTP the envelope is standard so all our components can speak the same language, output consistent log data and generally behave in a more homogenous way than you typically find in a mixed bag of microservices - what you put in the message payload is down to you and your services, of course for our internal systems gob is the natural choice for the payload as well.

## Some Trivial Calculations

I did some initial calculations and benchmarking comparing gzipped JSON to ‘snappy’ JSON to raw gob. I suppose it would be useful to include lz4, but as far as I can tell snappy is not a million miles different.

The data is a pretty tabular in format being lists of (300 ish) financial transactions, so naturally this does not favour normal JSON which repeats the ‘column’ names, however compressing virtually removes this disadvantage.

The original data contained mainly text data, which does not favour binary.

I understand that without the example data and code this is unscientific, as it is impossible to reproduce, and therefore can be read only as a hint of what the truth may be. If there were any more interest than myself making notes for me then I could do a proper job later.

|                               | Raw Gob |   Raw Json  | Gzipped Json | Snappied Json |
|-------------------------------|:-------:|:-----------:|:------------:|:-------------:|
| Size  (bytes)                 |  72111  |    252512   |     27115    |     47271     |
| Size (xGob)                   |    1    |      3.5    | 0.38 (/2.66) |  0.66 (/1.53) |
| Decode Speed  (ns/op)         | 1453910 |    1750450  |    3557621   |    3011828    |
| Decode Speed (xGob)           |    1    | 0.83 (/1.2) | 0.41 (/2.45) |  0.48 (/2.07) |

*xGob = multiplier of the Gob figure*

This data is for the payload or body only, whic relates to storage size (minus keys and indexes). 

If HTTP is used as the transport protocol then the average 300-500 bytes of HTTP header should be taken into account when considering bandwidth, though with internal systems this would be more like 100-200 bytes. 

Header size becomes a profound factor when transporting small payloads, making TCP a smarter choice, but then introducing similar debugging issues as faced with binary encodings.

Our internal gob header is typically 40-50 bytes (mainly being the 36 byte text representation of [uuid](http://tools.ietf.org/html/rfc4122), which would be better passed round as the raw 16 bytes)

For a 40 byte binary payload (an array of 20 `int16`'s for example) HTTP could easily cost 10x more bandwidth than TCP.

The summary comparison to gob for our large data (remember this is a sub optimal use case for binary) is...

|                  |              |              |
|------------------|--------------|--------------|
|Gzipped Json      | 2.5x slower  | 2.7x smaller |
|Snappy Json       |   2x slower  | 1.5x smaller |
|Uncompressed Json | 1.2x slower  |  3.5x bigger |

So even with binary unfriendly large text data its hard to strike a better balance than plain gob. If raw space and bandwidth were the primary concern then I would **go for gzipped JSON over TCP**.

I’m surprised how sluggish Snappy is. Others have [reported](https://groups.google.com/forum/#!topic/golang-nuts/7T1AKfDAOcQ) the pure Go implementation to be quite slow.

None of the decoding benchmarks streamed the data - new encoders were made every loop, I know Gob would be even faster with this optimisation, all decoders decoded into an `interface{}`, they would all have been quicker if a particular `struct` was used.

## Testing Binary Payloads

It is tempting to emphasise the more tangible representation of test data that a browser or even a prettified curl output offers, or the ease of editing a JSON file. 

It is only marginally more hassle to construct test structs (or the equivalent in your own language) in unit or integration tests, and in general a better idea, and something you’ll have to do anyway. 

In many languages (particularly Go) writing a curl like tool to interact with a binary rpc is, admittedly a tooling up overhead, however its a pretty simple, one off days work.
 
## Summary

A compressed text based encoding like JSON **is** more widely supported, quite compact and easier to test and debug, particularly if the consumer is a web browser.

**Very few scenarios justify anything other than compressed JSON** (swap JSON with XML if you are stuck in that particular hell)

If you are approaching or know you will approach network limitations, or will save £1000’s per month (substitute for your own acceptable budget) and want to keep response rates up then a binary format may well be the right choice for all payload sizes.

If your system chatters lots of small chunks of data with a high ratio of none text to text data, at very high message rates, then a binary encoding particularly over TCP (thereby avoiding the HTTP header overhead) is by far the more sensible choice.

Testing and debugging binary protocols is not as bad as people make out, and should not be a massively deciding factor.
