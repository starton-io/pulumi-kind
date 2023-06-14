---
title: Pulumi-kind Installation & Configuration
meta_desc: Information on how to install the Kind provider.
layout: installation
---

## Installation

The Pulumi Kind provider made by [starton](https://www.starton.com/) is available as a package in several Pulumi languages:

* JavaScript/TypeScript: [`@starton/pulumi-kind`](https://www.npmjs.com/package/@starton/pulumi-kind)
* Go: [`github.com/starton-io/pulumi-kind/sdk/go/`](https://pkg.go.dev/github.com/starton-io/pulumi-kind/sdk)

### Provider Binary

The kind provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource kind <version> --server github://api.github.com/starton-io
```

Replace the version string with your desired version.
