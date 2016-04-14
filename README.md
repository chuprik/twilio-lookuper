# Twilio Lookuper

Provides retrieving information about a phone number.

## Installation

```
go get github.com/kotchuprik/twilio-lookuper
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/kotchuprik/twilio-lookuper"
)

const (
	SID   = "foo"
	Token = "bar"
)

func main() {
	client := lookuper.New(SID, Token)
	data, err := client.DoLookup("74957392222")
	if err != nil {
		panic(err)
	}
	fmt.Println(data["phone_number"])
}
```
