---
title: "Bazel Fawlty"
description: "Like Basil Fawlty, Bazel promises a great deal, but makes a bit of a fuss about nothing, especially in Golang projects"
date: "2018-03-06"
categories: 
    - "golang"
    - "go"
    - "tech"
    - "build"
    - "bazel"
    - "engineering"
    - "quality"
    - "continuous integration"
---

Like [Basil Fawlty](https://www.youtube.com/watch?v=mv0onXhyLlE), Bazel promises a great deal, but makes a bit of a fuss about something that could have been done so much more easily. This is at least true in any medium to large sized Golang project, it is worse if using OSX. Whilst Bazels goals are desireable they are easily and more simply achieved with the Go tools, especially since go 1.10.

<!--more-->

We spent 6 months getting bazel build to work in our codebase, and a year working with it. Now we are pulling it out, here is why.

But first some background and key concepts, if you are familiar with the concepts then skip to specific Bazel sections ...

Background
----------
The two most important features of a build system are:

* Correctness - which is actually not so easy to define. 
  * Firstly it has to do what we expect it to do. 
  * Secondly it has to behave consistently. For example if the same inputs are given to the build system the same output should be produced, also if slightly different inputs are given,the output should only contain the effects of those explicit changes. 

* Speed - The quicker the better, more on that later.

- - -

### (An Aside) Hope **is** a strategy 

Recently anyone who has read the Google SRE book, or is an Xoogler or is in that extended network can oft be heard reproducing the old business mantra: "Hope is not a strategy", not uncommonly as a bit of a facetious put down. Whilst it is a succinct and catchy phrase, it is normally not true, or at best not particularly valuable. When you get down to it hope is our baseline. In fact I'll add another 'evil' to that... assumption. At some point we have to hope and assume. for example eventually we hope the compiler authors did a good job with the next version we are about to use, or we assume that the kernel fix was good. 

At some point our foundation is 'faith' that other engineers have done a good job - that we are standing on the shoulders of giants. We are never going to know definitively, however good our tests are, or however long we 'actively wait' for others to adopt first, we can only reduce the risks. At some point there is enough of a diminishing return on our efforts that we stop at some degree of hope that the rest of the system is as stable as we expect. 

- - -

**For the rest of this article we will assume the build system behaves as we expect given the set of inputs it receives, so that all we are focussing on is the consistency aspect of correctness, and we hope that the tools are bug free.**

Reproducibility 
----------------
Passing the same explicit inputs to a build system and getting the same output is the key characteristic of a reproducible build system. This behaviour is a strong indicator of a correctness. In reality actually using a build system to rebuild the same input multiple times is rare. Normally the build artefact itself is the asset that progresses down the release pipeline, and is archived for any rollback or debugging, but not always.

Whether a second build of the same source forms part of your deployment strategy or not, if your build is not reproducible then it is also quite likely not correct.

A reproducible build system that uses a buggy compiler is not 'correct', but as mentioned, compiler bugs don't count here.

### Hermetic Builds
Hermetic simply means that the system is sealed, is is not affected by inputs that we did not explicitly add to the sealed environment, and is one of the key ingredients to reproducibility.

Truly hermetic systems are hard (impossible) to achieve. The achievable goal in reality is to reduce differences in behaviour once the result of the build is used.

When people talk about introducing 'hermetic' builds, they mean introducing more isolation than before, or put conversely: less variability between subsequent builds of the same source input.

So achieving a hermetic build boils down to controlling everything that can affect the resultant outputs, namely:

1. Source files.
2. Tooling.
3. Environment.

### Same Source

Of course the single most important input to a correct build is correct source data.

It is important to know that the source for the build is itself strictly controlled. If you make two or more builds of a particular version of your source, then that source needs to be identical in each of those builds. This is clearly one of the reasons VCS's exist.

Git is a amongst the best in this regard, but even with Git attention has to be paid to how it is used and when files that are left after swapping branches may impact the build.

The major trip up in many types of project in this regard is the dependency management of third party sources. The same is true for Go. The solution is clear: only build source that exists in your own repo, hence 'vendoring' of some sort. Pulling in source from systems not under your control increases the risk of a different behaviour.

Many tools don't do decent hash based dependency analysis, the most famous being `make`. Go pre 1.10 also had problems here, depending on time stamps too much, making it quite likely that an incorrect build could be produced under reasonably normal circumstances.

### Tooling

It is all to easy to overlook the fact that once you change the version of Go on the build machine and build the same input source as you built before the version upgrade there is a strong chance your outputs are going to be different. The same may be said for clang, gcc or javac for instance. However Go is on a faster release cycle than most.

### Environment

In terms of building Go, the environment is: 

* The operating system.
* Environment variables. 
* Command line flags to build tools.
* Files that are in scope of the build tools. 

For Go the files that can influence a build are those under `$GOROOT` and `$GOPATH`. 

If you control Source, Tools and Environment correctly, then you will have a hermetic build, and if you have that with Go then you will have a reproducible build - in fact byte by byte reproducible.

Reproducible Services
----------------------

Being confident that your build is correct is important, and if you put in place a few simple things the Go tooling will give you that. Even if you make a couple of mistakes and your build is not fully hermetic, it is quite likely that the reproducibility of your build is one of your lowest risks. 

Typically much bigger problems will be due to:

1. **Correct code** - no need to worry about a solid build system, if your code is wrong or poor quality, untested etc.
2. **Runtime configuration** - needs to be well controlled. Config change is as dangerous as code change, or maybe more so, typically because it has been managed outside the rigour of a build system. With code now as deploy-able as config, gone is the historical separation based on rate of change, all but per-environment config is migrating into the binary. How your remaining config is deployed should be as well controlled as your code. If is is not, then reproducibility in your build is amongst the least of your concerns.
3. **Runtime environment** - Network policies, images, server specification, replica counts, routing config (if external balancers, proxies are in use) etc. These are an order of magnitude more likely to affect your service than a stray compiler flag for instance.

Then lets not forget the most uncontrollable factor of all: **workload**. Your system in a test environment or local cluster is very likely to behave quite differently under different load and (particularly) concurrency patterns. There are other techniques to manage this kind of variability, load testing for instance but if you are not managing that risk, it is likely to bite much harder than your subtly incorrect build.

You can spend as much time removing the last iota of randomness in your build, but it can not address the fact that in practice the vast majority of the variability in your running services will be caused by these other things day to day.

With that in mind, lets get back to the correct build problem. 

The Promise
-----------

* It promises a hermetic build so that even incremental builds are safe. (assuming your VCS is safe in this regard), by sanboxing build assets on the file system and controlling environment variables. 
* It has good hash based dependency analysis and local cache to be able to decide what tools need to be invoked, once you have explicitly declared the file dependencies, improving the correctness of incremental builds.
* It can also decide what tests to run based on its dependency analysis, at least at package granularity, optimising test run times (if incremental)

The above are claims that are hard to refute, however it claims it can speed build times up, whereas in our Go code base the opposite is generally true, even during developer incremental builds.

Bazel should in theory be able to decide if a docker image should be rebuilt, but we had trouble making anything work successfully in this regard, only getting as far as copying a binary into the folder and recreating the Dockerfile if the templates we use changed. In any case for all production builds we version stamp during linking, so the docker image will always need to be built.

Regarding build speed I think if the remote, or distributed cache would have worked it could have improved build times further. We had difficulty getting the cache to work, and when it did work it was quicker to build an asset than download it from the cache, others mileage here will probably be better. 

The Reality
-----------

To set the scene, our codebase is mainly Go, including vendor directories, and a considerable amount of data compiled in, we have 2.5M lines of code, a full build from a clean clone on one of our Jenkins slaves takes 1 minute.

We had between 8 and 20 engineers working on that codebase. Importantly we all develop on Mac, but run in production in linux.

Bazel:

* Enforced the version of Golang in use was the one set in the code repo in the `WORKSPACE` file.
* Isolated the build from any accidental environment variables changes.
* Allowed us to only check in the sources for generated files, and not the generated files themselves.
* Sped up some local unit testing.

It should also have sandboxed the file system to avoid accidental assets affecting the build, but for us it didn't - the Sandboxing on Mac did not work, and made it difficult to keep it enabled on the build machine. So essentially incremental builds were not guaranteed. In any case the build and unit test phase is less than 10% of the whole CI cycle - hermetic incremental builds were not guaranteed and not really needed.

For the record Bazel still actually uses Go to do the building - it is not a Go compiler itself. 

What Did Bazel Cost
-------------------

Firstly dont be fooled by how 'easy' anyone claims Bazel to be. It took a huge effort to get Bazel to the point where it was no longer causing regular problems. In fact 83 commits over 8 months in our codebase specifically to get Bazel working, on top of this the lead engineer had many PR's accepted into the bazel go rules git repo.

To understand the scale of the complexity that Bazel adds - it takes > 1000 command line parameters (https://docs.bazel.build/versions/master/command-line-reference.html), it is 1.3M lines of code (from a simple cloc-ing), it currently has 1200 open issues, and had three releases this year already (two months at the time of writing) and HEAD has broken the latest go_rules again (https://github.com/bazelbuild/bazel/issues/4659)  

**Working with bazel on OSX increases the cost** I realise a lot of our issues have been down to problems specific to Mac. The poster child of Bazel - Kubernetes - explicitly disabled Bazel on Mac for some time, I'm not sure if it is enabled now, but there are still open issue on Mac regarding Bazel. If we were developing on Linux we would possibly have continued with it, but I doubt it.

It added 22,000 lines of BUILD files into the repo, though all but a handful were auto created by a tool called Gazelle, which itself added a new step in the development process. If a file is added or removed or renamed, then a BUILD file has to be updated, either manually or by running `bazel run //:gazelle`, which could take 30 seconds, or if you switched to a new hash that spanned a go rules version change it could take upto 4 mins to run...

`bazel run //:gazelle  1.44s user 1.64s system 1% cpu 3:42.60 total` 

Gazelle uses the Go tools to understand dependencies to create the BUILD files... so we are invoking Go tools to create explicit dependency declarations, to feed into Bazel, to decide when to use Go tools to build things. A gentle irony there. 

It increased the build time to 3 minutes (from 1 min before), though this includes compiling the unit tests. When you count Bazel running the tests at 1.7 mins, and Go doing it in 2.5. The total build and unit test time for Bazel is 4.7 mins vs Go's 3.5 mins.

It added a dependency on the JVM. (I probably have an overly painful past with the JVM in production), suffice to say we had to bump up the RAM on the build machines. 

It added another tool to upgrade, and it was/is changing often, the go rules changed even more often - 14 times in 1 year, sometimes with lock-step dependencies on the Bazel version, which required upgrades to all development machines and build slave images. We only bumped the Go version twice in that time. Even if the rules update did not specify a different version of Go - it would still trigger a re-download of go. A git bisect, build, debug cycle became a farce, watching the exact same go being downloaded (adding 3 mins over our network) to build a tiny delta.

It added a foreign command syntax, for all to learn. A simple build command is ok, being ...

```
bazel build  //myfunky/package
```

vs

```
go install ./myfunky/package
```

but individual tests are run like this...

```
bazel run //myfunky/package/:go_default_test -- -test.run=TestMyThing
```

previously...

```
go test ./myfunky/package -run=TestMyThing
```

...annoying.

The new syntax does not work with auto complete natively and needs some [convoluted plugin](https://github.com/bazelbuild/bazel/blob/master/scripts/zsh_completion/_bazel) to make work, but that too uses bazel behind the scenes and its bootstrapping causes a frustrating lag.

Fairly trivial things, and easy to look down on an engineer not enjoying this discomfort, but it is a valid muscle memory problem, that should not be dismissed.

It created a whole set of hoops to jump through to nicely support IDE's code completion, symbol finding and VSCode at least would still need to compile using native tools. An example open issue (https://github.com/bazelbuild/rules_go/issues/512) resulting in us committing symlinks to the bazel generated files, in fact a symlink to the symlink of the generated files...

Test coverage is not supported in Bazel, so that needed separate native builds as well as the Bazel build - sub-optimal.

We don't use cross compilation on the build machines, but it is handy to do quick builds for local testing in a container. Bazel had problems, and is still not fully supported I believe, so we reverted to standard Go here as well, further bypassing any benefits of the cache.

Whilst ultimately we started to accept the new system and extra Gazelle step, and could just about hack a BUILD file to do what we needed (though some people have checked in `stringer` and `bindata` output) it invariably caused some issue if we wanted to upgrade Go - most recently with Go 1.9.4 needing some explicit import path declaration or other.

This was the straw that broke the Bazel back. That and two major upgrades in go 1.10...

* Test caching.
* 'correct' dependency analysis and artefact caching.

And the fact that the version of Bazel and go rules changed more frequently than Go surely casts some doubt onto which system may be adding more variability.

These things made it all the more obvious Bazel was not the right tool for us.

- - - 

### (An Aside) A Note on Bazel Builds on Different Operating Systems

Let's be clear - there is no need to consider reproducibility between different operating systems running the build.

When Bazel builds Go code on linux, in an Alpine container for instance, it uses the external linker `gcc` and asks for the gold linker for ELF files. Gold is not used on OSX, Bazel also has to add a hack wrapper script: `cc_wrapper.sh`,  because...

```
# OS X relpath is not really working. This is a wrapper script around gcc
# to simulate relpath behavior. 
```

This only relates to the linker, which 'probably' behaves pretty consistently - but the point is there are differences.

I don't think anyone is suggesting, even with cross compilation (which Bazel does not support yet), that there are any guarantees when different versions of build OS and GCC are in play.

Stronger guarantees are achievable by building in a well controlled container. But in practice this level of consistency between dev and production - literally does not matter 99% of the time.

Whatever anyone says - running on your local machine (yes even in containers) is **nothing like** running in production. 

- - - 

Back to the Future.
-------------------

No doubt if we ever reach Google scale, or have many build time dependencies between many different languages, then an explicit dependency declaration and distributed artefact cache, will start to pay off - especially if the development machines are the same OS as the build machines.

However in our medium size pure Go environment is it simply not worth the investment.

It took one weekend to remove all trace of Bazel **and** have a correct build.  

1. The build machine downloads the version of Go specified in the repo.
2. Builds are always a clean clone into a newly launched container. (using Kubernetes to schedule)
3. All source is checked in - including generated files.
4. Make .PHONY targets are used as familiar entry points in front of a few shell scripts

We are already Vendoring (and you still need to do that with Bazel)

Even if you are not set up to schedule a fresh container on every build you can put a certain amount of trust in your CI job to help with your controlled build environment (with caveats). More than that though, Go is well behaved and it is clear about what env vars may affect the build. In fact its defaults are probably all you need..

`go env`

```
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/dan/myproj"
GORACE=""
GOROOT="/usr/local/Cellar/go/1.8.3/libexec"
GOTOOLDIR="/usr/local/Cellar/go/1.8.3/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/cn/8f4yl60x4y19m4m7lc2d8r380000gn/T/go-build643053316=/tmp/go-build -gno-record-gcc-switches -fno-common"
CXX="clang++"
CGO_ENABLED="1"
PKG_CONFIG="pkg-config"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
```

So you can check these vars in if you like and export them on each new build.

We have swapped the `make bazel-gazelle` step for a `go generate` step, and instead of our build server triggering a script to confirm BUILD.bazel files are up to date, we confirm that the generated files are.

All in all removing Bazel has meant an increase of 40 lines of bash.

We will introduce a pre-commit Git hook that checks the local version of Go, and if the generated files are out of date, just to encourage the incremental local development environments to be in close sync (even though this is the rarest source of problems)

Conclusion
----------

Bazel has brave ambitions, and may suit bigger projects, and will one day be stable and valuable. In our project it was just too young and too much trouble, and could not offer anything that can not be achieved more simply with standard tools.
