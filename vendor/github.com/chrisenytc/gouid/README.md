# GoUID

> A go package for generate random uids by length

[![Build Status](https://secure.travis-ci.org/chrisenytc/gouid.png?branch=master)](https://travis-ci.org/chrisenytc/gouid) [![Build Status](https://drone.io/github.com/chrisenytc/gouid/status.png)](https://drone.io/github.com/chrisenytc/gouid/latest)

## Getting Started

1º Install GoUid

```bash
$ go get github.com/chrisenytc/gouid
```

## How to Use

Example:

```go
package main

import (
	"fmt"
	"github.com/chrisenytc/gouid"
)

func main() {
	n := gouid.UId{8}
	//Set seed to get random uids
	n.SetSeed()
	test1 := n.NewUId()
	test2 := n.NewUId()
	//Print results
	fmt.Println(test1, test2)
}
```

## Contributing

See the [CONTRIBUTING Guidelines](https://github.com/chrisenytc/gouid/blob/master/CONTRIBUTING.md)

## Support
If you have any problem or suggestion please open an issue [here](https://github.com/chrisenytc/gouid/issues).

## License 

The MIT License

Copyright (c) 2014, Christopher EnyTC

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the "Software"), to deal in the Software without
restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following
conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

