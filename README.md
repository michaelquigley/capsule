# capsule

_(this is a work-in-progress)_

`capsule` is a software framework that defines a publishing format, designed to contain structured prose writing and multi-media artifacts. A publication created using the `capsule` framework is referred to as "a capsule".

Let's start with a picture:

![anchor-feature](docs/images/anchor-feature.png)

A capsule contains a multiplicity of hierarchical "anchors". Think of an anchor as performing in both the role of a "folder" and also as a "document". An anchor is a folder, in that it can contain other anchors. An anchor is a document in that it can be "displayed" or "rendered" and shown as an independent unit of "content".

In the diagram above, `/`, `book/`, `chapter 1/`, etc. are all anchors.

The other items contained in the `chapter 1/` anchor (`picture.jpg`, `index.md`) are "features". Features are sub-components of the anchor that get combined into a presentation of the anchor. The anchor acts as an "namespace", containing all of the items related to that specific unit of content.

This is intended to be abstract, with well-defined idioms. In a typical example where we're building a "blog" or a "journal", there will be a large number of peer anchors, one for each "entry" or "post". Each of those anchors will likely contain prose text in some structured format (markdown, gemtext), along with configuration, metadata, images, multimedia, binaries, or other artifacts. Rules built into the `capsule` framework will then define how the features are presented, based on default idioms, augmented with metadata expressed within the anchor to customize the behavior.

The picture also shows two anchors which contain features named `structure.json`. These are "structural directives". Structural directives describe in high-level the relationships between the anchors within the capsule. The structural directives are used by applications presenting capsule content to drive their presentation. You can think of the structural directives as an abstract representation of the navigational elements necessary to present the capsule.

Structural directives describe things like, "this capsule contains a timeline of anchors". Or, "this capsule contains a graph of hypertext". Structural directives can apply to subsets of the capsule, hierarchically. This allows for different portions of the capsule namespace to follow different presentation rules.

The ultimate goal of the `capsule` framework is to create a stable, simple, semantically-rich content format that can support multiple types of content development. Authors creating capsules should encounter minimal friction from the format or the tools. The capsule format should provide relatively flexible expressive power to allow the creation of varied content structures. This format should offer clean integration with a multitude of content-creation tools, in that it will be manifested as folders in a filesystem, containing files.

I intend to use `capsule` to provide myself with a publishing pipeline for my regular creative journal output, along with technical writing and technical research output for a number of projects. I would love to collaborate with other creatives and developers who find these ideas compelling and might want to use `capsule` in their own work.

## Capsule Presentation

A capsule is intended to be a single source of truth for any number of presentations of the capsule's content. Initially, the `capsule` framework will ship with a set of "transpilers", which can create presentations targeting presentations for multiple formats. For example, there will definitely be a transpiler that can generate static HTML. There will likely be a transpiler that can produce a PDF. The transpiler framework will support extensibility by design, and will allow users of `capsule` to tailor the presentation to work for their unique requirements. Ultimately the content format and the idiomatic rules are the main concept, and users will be free to innovate tooling on top of those ideas.

It's also on the roadmap to include a [Gemini][gemini] transpiler.

### Capsule Browser

Longer term, it might be useful to create a native `capsule`-based browser stack, and corresponding network layer. This approach would be similar to what [Gemini][gemini] has done. In the `capsule` universe, the structural directives end up dovetailing with the capabilities in these browsers, such that the users of capsules have more ultimate control over the layout and presentation of the content. The structural directives capture the semantics of the presentation, and the browser can provide extensive additional capabilities on top of those semantics.

A native capsule browser would allow capsules to live directly on the internet, and be presented without going through transpilation steps.

## Cryptographic Authenticity

`capsule` will include support for cryptographically authenticating and signing content. I'm going to hand-wave this away for the time being until I've had some time to work through but... but the intention is to include native support in the format.

## Future-Proofing

What if we get to a point where we've invested a lot of time and effort into building capsules, and we decide we want to go somewhere else?

We can just create a transpiler to port our capsules to whatever native formats we would like.

[gemini]: https://gemini.circumlunar.space/