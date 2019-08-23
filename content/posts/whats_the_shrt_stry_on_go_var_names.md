---
title: "What's the Shrt Stry on Go Var Names"
description: "Go likes it short, too short?"
date: "2018-08-23"
categories: 
    - "golang"
    - "go"
    - "tech"
    - "development"
    - "engineering"
    - "quality"
---

Is I have mentioned in a previous post we are all products of our own histories, and this will definitely impact your appreciation or frustration of the subtleties of the Go community guidelines on naming.

<!--more-->

I have, as is not uncommon, been discussing Go naming conventions with a new team with a few newish Gophers in that team. As I have had this conversation considerably more than once, I thought I would write those thoughts down.

One colleague referenced an article which represented the conflict well.   

So I decided to document those thoughts as a response to that article. It is 4 years old but the issues are the same:

http://michaelwhatcott.com/go-code-that-stutters/ 

I make my points below in sections named as per the article.

Field and Variable Names
------------------------

The author appears fine with the relatively standard short forms, though has a small issue with `err` reluctantly accepting it, but considers the member `b.n` a 'fail'.

I agree that `n` could probably have a longer name in it's scope as a struct member, it's a marginal call. 

This one is little more interesting though. 

Other names are not quite right `count`, `position` and `length` are all a bit wrong and even misleading. Maybe `index` or `offset`, but offset of what? `currentCopyIndex` or `bufferCopyPosition` which are a bit too long and ugly. `index` may cause folk to think of a larger data structure - until you scan it's type as an int. 

It needs a few words to accurately describe what `b.n` is. Rather than be misleading it is better to stick to a relatively none descriptive name (that has some history in this role) and allow the reader the few seconds to check what it does. 

Looking in the context of the whole file allows you to to pretty quickly get a good handle on the role of `b.n` - more so than that single snippet in isolation.

Also for me and many other people `n` (int) is very often an index into an array.  

So given the slight trickiness ... when you look at this...

```
    func (b *Writer) WriteString(s string) (int, error) {
        nn := 0
        for len(s) > b.Available() && b.err == nil {
            n := copy(b.buf[b.n:], s)
            b.n += n
            nn += n
            s = s[n:]
            b.Flush()
        }
        if b.err != nil {
            return nn, b.err
        }
        n := copy(b.buf[b.n:], s)
        b.n += n
        nn += n
        return nn, nil
    }
```  

Versus...

```
    func (b *Writer) WriteString(s string) (int, error) {
        count := 0
        for len(s) > b.Available() && b.err == nil {
            n := copy(b.buf[b.offset:], s)
            b.offset += n
            count += n
            s = s[n:]
            b.Flush()
        }
        if b.err != nil {
            return count, b.err
        }
        n := copy(b.buf[b.offset:], s)
        b.offset += n
        count += n
        return count, nil
    }
```

Then I think that the understandability of both pieces of code are similar. 

Even if for you `n` carries no history, you are very likely to need more context and a bit of time to see what `offset` / `index` is offsetting or indexing into to understand things fully, much the same as with `b.n`

The code here is pretty simple to understand in any case.

The second snippet is just a bit more 'noisy'. The first snippet is slightly more aesthetic, it looks more simple, has less potential to mislead and is as tricky (or easy) to follow as in the first case.

For this specific issue the code is simple enough that arguing about `b.n` vs `b.offset` or `even b.bufferCopyPos` adds almost nothing to the quality of the code.

Receiver Names
--------------

The author claims that: 

>writer.--- or just self.---, both of which communicate more accurately than b, which could be any old variable (local, package-wide, or the receiver itself).

The google code [review comments](https://github.com/golang/go/wiki/CodeReviewComments#receiver-names) say it more clearly than I, but I totally agree with them `self` and `this` carry baggage that is entirely misleading in Go - there is no object or object hierarchy - no `self.parent`. These must be avoided for avoidance of confusion.

The point made in the review comments about receiver names is that they are used a lot in really well defined contexts. Using longer receiver names add a lot of clutter.

Even with the name `writer` you have to glance back to see where it is defined, it becomes habitual as a Gopher to scan the receiver definition. You will need to do this in both cases `b.` or `writer.`

The author would prefer the following code: 
   
```
    func (writer *Writer) WriteString(s string) (int, error) {
        count := 0
        for len(s) > writer.Available() && writer.err == nil {
            n := copy(writer.buf[writer.offset:], s)
            writer.offset += n
            count += n
            s = s[n:]
            writer.Flush()
        }
        if writer.err != nil {
            return count, writer.err
        }
        n := copy(writer.buf[writer.offset:], s)
        writer.offset += n
        count += n
        return count, nil
    }
```   

Compare to the original snippet as above...

```
    func (b *Writer) WriteString(s string) (int, error) {
        nn := 0
        for len(s) > b.Available() && b.err == nil {
            n := copy(b.buf[b.n:], s)
            b.n += n
            nn += n
            s = s[n:]
            b.Flush()
        }
        if b.err != nil {
            return nn, b.err
        }
        n := copy(b.buf[b.n:], s)
        b.n += n
        nn += n
        return nn, nil
    }
```  

...and intuitively feel which is cleaner / clearer. Your preference may be different from mine.

Personally I find both bits of code pretty acceptable, both should pass most reviews. One is slightly more messy for little gain, and it is a fair opinion to say there is a loss of readability, it is equally fair for someone to feel the original is less readable. But by how much really? Enough to cause a debate and a refactor? In this case I don't think so.

Given that the majority of the Go devs that arrive at your company enjoy Go because of it's "Go-ness" then it will cause less PR noise in the future if you tend towards the "familiarity admits brevity" concept.

Function Names
--------------

The author's comments on function names make it clear that they are not too familiar with the particular names in question, or at least are arguing that they will not be familliar for a lot of people, but for me they carry a lot of history and meaning.

`format.DecomposeFormat` - means nothing to me, 

`Decompose` can mean a lot of things.

`fmt.Sscanf` - I know immediately what this does.

I would also argue that even for a new dev not familiar with `C` the following...

```
fmt.Sscanf
fmt.Scanf
```

Instead of the suggested... 

```
format.ScanFormat 
format.DecomposeFormat
```

Has better consistency. 

Both will need a new developer to investigate what they are, but once you understand the clear pattern with the prefix 'S' I predict your learning curve will be less with the former.

For those like me that clicked with Go, specifically those from a 'C' background, the first lot already have a lot more meaning. 

The author goes onto suggest... 

> so why not import the format package like it was built-in?
 
The simple answer is because that is even more confusing! 

Even the very simple example they use shows that confusion.

```
   PrintLine("Which line is easier to grok?")
   fmt.Println("Which line is easier to grok?")
```

The second one is explicit and clear - it is from the `fmt` package. The first one looks like a package local func (oops I meant function!) but it isn't - it is even hard for some IDE's to work out what happened :)

Obviously for me the actual function names are equally understandable (history again I suppose), but the first one I immediately wake up to as being 'odd' I have not seen it before - so this one I have to think about more.

I understand that another person may find `PrintLine` clearer than `Println` - but please don't source in the packages like this, in general (there are a few good cases for sourcing in packages).   
 
(Mixed-up) Example
------------------
The final section references: 

[Andrew Gerrand's - What's in a name?](https://talks.golang.org/2014/names.slide#7)

And the author swaps the [good and bad](https://talks.golang.org/2014/names.slide#7) examples, all I can say on this is as per the example above: both are clear enough, one is cleaner, one is more consistent with the broader community expectation, and that can be a good thing.

Conclusion
----------
It is good that the author recognises that this debate in these examples are relatively low value, and in general just a personal preference. There are other cases where this brevity of naming is indeed an important debate.

I think there is a bit of confusion regarding the guidance from the Go team - the guidance is how they do Go, and I agree how they would like the community to do Go as well. Consistency is powerful.

My Conclusion
-------------

I think it is fine for anyone to use the language as they prefer, but when you are in a minority within a community and you actively work against the common patterns, and you are not working in isolation then you are inviting ongoing conversations about why your code follows different guidelines.

Even if you don't completely agree with the whatever guidelines but they are not *too* harmful in your eyes, why not just go along with the flow? The path will be easier, will end up more enjoyable, and your code no worse quality. It will be considered more readable in a lot of peoples eyes, even if not yours. 
     
I think it is a great mistake to fight against the language we choose, or to dilute it's conventions with personal preferences based on individual history that explicitly fights against the style the language promotes. I think in general there should be a familiarity and consistency, not a few changes here and there, because of individual preference.

One reason I clicked with Go was precisely because it supported strong feelings I had after programming for the last 27 years in, C, C++, C#, Java, python, PHP (and some work in javascript) in particular it gave me 'permission' to use short variable names again, but this time with some good guidance about when **not** to use them.

I understand the core Go team guidance, and can argue for it, so it is not just a defer to authority but... it has had a lot of thought by many many highly experienced developers working in huge code bases, with large teams and turnover all with a laser focus on understandability for the next person.

The `w` vs `writer` or similar name debate in the case of the scope being < 20 lines, for instance is a marginal call either way. It is a relatively low value decision, and I completely understand the desire for `writer` but it adds very very little value, for a little more clutter, perhaps some marginal context but at the cost of contradicting the generally accepted conventions of the language.

I would say to any team facing the same friction, "lets not debate this too much longer and move forward erring toward the shorter names." 

If anyone coded `client` where I would prefer to see a `c` I will normally not be bringing it up in review other than perhaps `cleaner as c`. But then I would be OK, if not exactly happy, for the author to ignore and move on.