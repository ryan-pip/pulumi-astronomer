---
title: Astronomer Installation & Configuration
meta_desc: Information on how to install the Astronomer provider.
layout: installation
---

## Installation

The Pulumi Astronomer provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@ryan-pip/pulumi_astronomer`](https://www.npmjs.com/package/@ryan-pip/pulumi_astronomer)
* Python: [`pulumi_astronomer`](https://pypi.org/project/pulumi_astronomer/)
* Go: [`github.com/ryan.pip/pulumi-astronomer/sdk/go/astronomer`](https://pkg.go.dev/github.com/ryan.pip/pulumi-astronomer/sdk/go/astronomer)
* .NET: [`RyanPip.Astronomer`](https://www.nuget.org/packages/RyanPip.Astronomer)


## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `astronomer` provider:

- `astronomer:apiKey` (environment: `astronomer_API_KEY`) - the API key for `astronomer`
- `astronomer:region` (environment: `astronomer_REGION`) - the region in which to deploy resources

### Provider Binary

The Astronomer provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource astronomer <version>
```

Replace the version string `<version>` with your desired version.
