# go-emoji-shortcode

[![Go Version](https://img.shields.io/github/go-mod/go-version/floatpane/go-emoji-shortcode)](https://golang.org)
[![Go Reference](https://pkg.go.dev/badge/github.com/floatpane/go-emoji-shortcode.svg)](https://pkg.go.dev/github.com/floatpane/go-emoji-shortcode)
[![CI](https://github.com/floatpane/go-emoji-shortcode/actions/workflows/ci.yml/badge.svg)](https://github.com/floatpane/go-emoji-shortcode/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A small, dependency-free Go library for mapping emoji shortcodes (`:smile:`, `:heart:`, ...) to Unicode emoji characters. It is designed for text editors and TUI composers that want live shortcode preview and suggestion support.

Extracted from [matcha](https://github.com/floatpane/matcha) so it can be reused by other Go projects.

## Install

```bash
go get github.com/floatpane/go-emoji-shortcode
```

Requires Go 1.26+.

## Usage

```go
package main

import (
    "fmt"

    "github.com/floatpane/go-emoji-shortcode"
)

func main() {
    e, ok := shortcode.Lookup("smile")
    if ok {
        fmt.Println(e) // 😄
    }

    matches := shortcode.Suggest("sm")
    for _, m := range matches {
        fmt.Printf("%s -> %s\n", m.Code, m.Emoji)
    }
}
```

## API

- `Lookup(shortcode string) (emoji string, ok bool)` — exact shortcode to emoji.
- `Suggest(prefix string) []Match` — fuzzy-matched suggestions for a prefix (e.g. `"sm"` matches `:smile:`, `:smirk:`, ...). Results are sorted by relevance.
- `All() []Match` — every known shortcode/emoji pair.

## Documentation

Full API reference: [pkg.go.dev/github.com/floatpane/go-emoji-shortcode](https://pkg.go.dev/github.com/floatpane/go-emoji-shortcode)

## Contributing

PRs welcome. See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT. See [LICENSE](LICENSE).
