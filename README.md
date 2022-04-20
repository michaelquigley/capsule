# capsule

`capsule` is a software framework that defines an opinionated, extensible content management model.

`capsule` provides solutions for _content authors_ to compose, curate, and publish collections of structured content as coherent publications. `capsule` has a specific focus on the indie web, and independent, non-platform publications.

`capsule` provides a programming framework, which allows _software developers_ to create and extend solutions in an organized, uniform manner with beneficial properties. `capsule` is implemented in clean, modern, idiomatic golang.

`capsule` is intended to be format agnostic. Use the same content collection to publish into a multiplicity of formats (static web, PDF, gemini, etc.).

Use `capsule` when you want structured content, but don't want a heavyweight service-oriented content repository. Or use `capsule` when a conventional static site generator isn't quite enough (or the right kind of) structure.

## Capsules

In the `capsule` universe, we refer to a single "content repository" as a _capsule_. At this time, the concept of a capsule is more or less aligned with a single "publication", but this decision is not necessarily important to the framework.

The default content model included with `capsule` is a simple implementation built around a conventional filesystem tree. Filesystem primitives such as hierarchical directories, files, and file extensions are used along with `capsule`-specific conventions allowing content authors to build semantic meaning into their content.

The default filesystem-based capsule model is extremely lightweight, and will work on any reasonable computer system. Use a Raspberry Pi to manage your content workflow.

## Interfaces

`capsule` provides interfaces:

* around the filesystem structures defining a capsule, such that content authors have a stable representation to use in developing large, long-lived content repositories

* around the programming model used to operate on capsules, such that developers have a coherent set of abstractions that can be used to develop tooling and systems that operate on disparate capsules in a uniform way; develop solutions for content authors and distribute them

* around the static web compiler, it's toolkit and its extension points; work on styling and presentation separately from content authoring

* that allow for clean specialization through extension; developers *will not* have to create forks of `capsule` just to maintain their own customizations; use capsule like an extensible framework

## Meta-structures

The `capsule` framework provides first-class support for _meta-structures_. Meta-structures are programmatically generated content derived from the content objects contained within a capsule.

Generate tables of contents, indexes, search databases, or timelines, injecting those into the model cleanly so that they can be consumed by downstream publishing components, which bind them into a presentation.

Use meta-structures to post-process and augment content objects (semantic linking, summarization, etc.) before sending content down the publishing pipeline.

Easily extend `capsule` with any kind of meta-structure you can imagine.

## The Filesystem

Using the filesystem as a facility for managing content objects allows content authors to work with pretty much any kind of tooling for both creating content, and also for managing versions of capsules.

A tool like `git` (with LFS support enabled) could be used to manage versions of objects in a capsule. Any solution that can manage trees of files will be suitable for managing the versioned state of a capsule.

Like most other parts of `capsule`, if you want to use a different type of repository implementation you can replace the reference filesystem implementation through clean interfaces with well-defined semantics.

### Filesystem Primitives

`capsule` leverages the filesystem to represent a capsule as a hierarchical tree of _nodes_. Nodes contain a multiplicity of _features_.

![anchor-feature](docs/images/anchor-feature.png)

In the above image, `/`, `book/`, and `chapter 1/` all represent nodes. `(structure.json)`, `(picture.jpg)`, etc. all represent features.

A specific `capsule` implementation can decide if nodes should represent "pages" ("articles", "entries", etc.) or if nodes should represent larger structures like "chapters" or "books". This becomes and organizational choice in how you build your capsules.

In the filesystem implementation (and the other reference implementations), nodes are represented by directories and features are represented by files inside that directory. This gives us a good starting point to build capsules right out of the box that cover a wide range of scenarios.

## Static Web Compiler

`capsule` ships with a static web compiler reference implementation. The compiler is designed to be easily extensible to support multiple types of web structures ("blogs", "journals", "books", "wikis", "zettelkasten", etc.). The compiler itself provides interfaces separating its concerns and allowing for clean specialization through extension. A number of reference implementations and examples are included to make it easy to understand how to use `capsule` as the foundation for your own static web content strategy.

By default, the static compiler uses a node within a capsule to represent a "page", and features to represent the structured prose documents, images, multimedia, and other assets that are unique to that page.

## Changes to the Interfaces

`capsule` will provide migration paths for both content authors and for developers. The intention is that content authors will have reasonable tooling from the `capsule` framework to allow them to migrate their capsules from older idioms to newer idioms. Software developers will have framework support for managing code that operates on multiple `capsule` versions.

It is an important property of `capsule` that capsule structures remain stable, such that those capsules will be supported many years into the future. `capsule` is intended to be a long-term solution for content creation, publishing, and archival. Yet, easily remove these capabilities if you want to keep your implementation lean and focused.

`capsule` is built upon the foundation of self-hosting, right-sized solutions, and the modern decentralized web. It aims to do this without sacrificing the right kinds of ergonomics for modern developers, and with a minimum of bloat.

## Roadmap

I am developing `capsule` to be the foundation for my long-term content strategy across several different web sites and self-published releases. As such, I'm naturally tailoring the development to suit my own use cases. Getting to a `v0.1`, which includes the first generation model containing the relevant meta-structures, and a reasonably-baked static web compiler are my main goals. I will be champagne-drinking (dog-fooding) these bits in my own workflows across a number of projects involving multiple people.

In `v0.2` I would like to develop a compiler implementation that can re-target capsules into single-document formats, like PDF. I would like to be able to iteratively develop both web sites, and book-like renderings from the same content.

It is also a target of interest to develop a compiler implementation that can produce a [Gemini][gemini] output.

I am actively interested in collaborators who are looking for a stable, modern, lightweight content management workflow.

[gemini]: https://gemini.circumlunar.space/