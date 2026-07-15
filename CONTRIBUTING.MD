# Contributing

Please refer to Cloudsmith's standard guide on [Open-Source Contributing](https://docs.cloudsmith.com/contributing).

## Toolchain

Host tools are pinned with [mise](https://mise.jdx.dev/) via `.mise.toml`
(`go` and `jq`). Run `mise install` once to provision them. Binding generation
also requires [Docker](https://www.docker.com/) — it runs the pinned
`openapi-generator-cli` image and is a separate prerequisite that mise does not
manage. GNU sed is used for the post-generation fixups (`gsed` on macOS via
Homebrew, plain `sed` on Linux).

Convenience tasks:

* `mise run generate` — regenerate the Go bindings
* `mise run update-bindings` — bump the version, regenerate, and open a PR

## Updating bindings

### Manual approach

* Update `PKG_VERSION` in `bin/generate`
* Run `bin/generate` (or `mise run generate`) to generate bindings
* Create a PR specifing API and binding version
* Create git tag and release in github

### Automated approach (local)

* Run `./scripts/update-bindings.sh` (or `mise run update-bindings`) to automatically update the bindings.
* This will then provide you with the URL for the PR to release the updated bindings.
* Preferred usage: `./scripts/update-bindings.sh`
* For full options and usage examples, run: `./scripts/update-bindings.sh --help`
* After PR is merged, follow the displayed instructions to create the git tag and GitHub release.

### Automated approach (CI)

* The **Update API bindings** GitHub Actions workflow regenerates the bindings
  and opens a PR — on demand (`workflow_dispatch`, optionally with a specific
  version) or daily on a schedule. It opens a PR only when the regenerated
  bindings actually change, and closes any superseded automated PRs first.

## Contributor License Agreement

By making any contributions to Cloudsmith Ltd projects you agree to be bound by the terms of the Cloudsmith Ltd [Contributor License Agreement](https://docs.cloudsmith.com/contributor-license-agreement).

## Coding Conventions

Follow existing structure and style of the build and release scripts. Languages bindings are generated automatically.

## Need Help?

See the section for raising a question in the [Contributing Guide](https://docs.cloudsmith.com/contributing).
