# textbackend [![Build Status](https://travis-ci.org/writescript/textbackend.svg?branch=master)](https://travis-ci.org/writescript/textbackend) [![Coverage Status](https://coveralls.io/repos/writescript/textbackend/badge.svg?branch=master&service=github)](https://coveralls.io/github/writescript/textbackend?branch=master)

simple textbackend to create a row/level organized content array.  
it is implemented in golang and docs can be found at [![GoDoc](https://godoc.org/github.com/writescript/textbackend?status.svg)](https://godoc.org/github.com/writescript/textbackend)

_pseudo data model..._

    content[
      { ... }
      { level: 0, text: "some words at one line of the document" }
      { level: 1, text: "the level can be used to tab one inside" }
      { level: 1, text: "..." }
      { ... }
    ]

## License
This code is published under an MIT license. See [LICENSE](LICENSE) file for more information.
