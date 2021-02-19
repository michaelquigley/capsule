# capsule

`capsule` is a software framework that defines a publishing format, designed to contain structured prose writing and multi-media artifacts. A publication created using the `capsule` framework is referred to as "a capsule".

A capsule contains a multiplicity of "content objects". Content objects can contain a multiplicity of "content items". A capsule also contains "structural directives", which describe the relationship of the content objects within the capsule to each other.

The ultimate goal of the `capsule` framework is to create a stable, simple, semantically-rich content format that can support multiple types of content development. Authors creating capsules should encounter minimal friction from the format or the tools. The capsule format should provide relatively flexible expressive power to allow the creation of varied content structures.

I intend to use `capsule` to provide myself with a publishing pipeline for my regular creative journal output, along with technical writing and technical research output for a number of projects. I would love to collaborate with other creatives and developers who find these ideas compelling and might want to use `capsule` in their own work.

## Structural Examples

A book published as a capsule could contain a content object for each chapter. Structural directives would then describe the sequential ordering of those content objects.

A journal published as a capsule could contain multiple structural directives describing different "views" of the content object timeline contained in the capsule. The capsule might also contain a primary structural directive that describes the linear ordering of the content objects as a timeline.

## Presentation, Initially

Initially the `capsule` framework will provide transpiler tooling to create capsule presentations in different formats:

* A simple static HTML framework will allow the creation of websites presenting the content objects and structural directives using idiomatic, extensible presentations

* `capsule` also intends to provide a [Gemini][gemini] transpiler

## Capsule Browser

Longer term, it might be useful to create a `capsule`-based browser stack, and corresponding network layer. This approach would be similar to what [Gemini][gemini] has done. In the `capsule` universe, the structural directives end up dovetailing with the capabilities in these browsers, such that the users of capsules have more ultimate control over the layout and presentation of the content. The structural directives capture the semantics of the presentation, and the browser can provide extensive additional capabilities on top of those semantics.