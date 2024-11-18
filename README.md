### Linters in GoLand

Currently, GoLand does not natively support `golangci-lint`; it can only be used with the help of the Go Linter plugin. The Go Linter plugin configures `golangci-lint` in GoLand, enabling access to linters that are not otherwise available within the IDE.

`golangci-lint` supports a wide range of linters, which you can find listed here: [Supported Linters](https://golangci-lint.run/usage/linters).

The Go Linter plugin requires `golangci-lint` to be installed as a binary. You can install it using Homebrew:

```bash
  brew install golangci/tap/golangci-lint
```

You can also create your own configuration files: [Configuration Documentation](https://golangci-lint.run/usage/configuration/).

For installation recommendations, see: [Install from Sources](https://golangci-lint.run/welcome/install/#install-from-sources).

To learn more about using `golangci-lint` in GoLand, refer to the [Integration Guide](https://golangci-lint.run/welcome/integrations/#goland).
