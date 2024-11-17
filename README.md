Golangci supports the following linters: https://golangci-lint.run/usage/linters

Go Linter requires golangci installed as a binary – brew install golangci/tap/golangci-lint

You may write your own configuration files: https://golangci-lint.run/usage/configuration/


In the previous example, there were unused variables (x, p, resp) that caused compiler errors. Many linters, including golangci-lint, may stop execution if they encounter compiler errors. This is because some linters require a valid AST (Abstract Syntax Tree) to perform their analysis, and compiler errors prevent the AST from being fully constructed.

Recommendations: https://golangci-lint.run/welcome/install/#install-from-sources

Usage of gollangci in GoLand: https://golangci-lint.run/welcome/integrations/#goland