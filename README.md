# Twilio Lookuper

Provides retrieving information about a phone number via [official API](https://www.twilio.com/docs/api/lookups) of Twilio.

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
	// data contains map[phone_number:+74957392222 national_format:8 (495) 739-22-22 carrier:<nil> url:https://lookups.twilio.com/v1/PhoneNumbers/+74957392222 caller_name:<nil> country_code:RU]
	data, err := client.DoLookup("74957392222")
	if err != nil {
		panic(err)
	}
	fmt.Println(data["phone_number"])
}
```
