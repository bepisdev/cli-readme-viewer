# CLI Readme Viewer

View Github README files from the terminal.

## Usage

Simply run the `crv` binary with values for the `-owner` flag being that of the account that owns
the repo you want to read and `-name` being the repo name itself for example, to get this README
file to render in your terminal you would run the following command.

```
$ crv -owner joshburnsxyz -name cli-readme-viewer
```

The program will render out the markdown of the README file in a terminal readable format.

### Current Limitations

- Only Github repos supported
- Only Markdown supported

## Installation

1. Clone repo
2. Run `go mod tidy`
3. Run `make`
4. Put the `crv` binary on your path

