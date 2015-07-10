# textbackend [![Build Status](https://travis-ci.org/writescript/textbackend.svg?branch=master)](https://travis-ci.org/writescript/textbackend)

testbackend to create row/level organized text content.  
the data model of one textbackend looks like this:

    content[
      { ... }
      { level: 0, text: "some words at one line of the document" }
      { level: 1, text: "the level can be used to tab one inside" }
      { level: 1, text: "..." }
      { ... }
    ]

More information and API-Docs can be found at [![GoDoc](https://godoc.org/github.com/writescript/textbackend?status.svg)](https://godoc.org/github.com/writescript/textbackend)

## License
This code is published under an MIT license. See [LICENSE](LICENSE) file for more information.
