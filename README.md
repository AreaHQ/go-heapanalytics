# go-heapanalytics

Provides a simple adapter for using the heap analytics API. See https://heapanalytics.com/docs/server-side for information on the API covered.

## Client options

There are a number of client options that can be set. 

The pattern used here is based on [Dave Cheneys blog post](http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

It's nice and clean as nil or blank values do not need to be passed in where default values are needed.

Some predefined options have been setup as listed below.

### URL 

([source](https://github.com/AreaHQ/go-heapanalytics/blob/master/client.go#L27))

Allows you to customise the client base url used by the adapter.

### HttpClient 

[source](https://github.com/AreaHQ/go-heapanalytics/blob/master/client.go#L35), [example](#track-with-httpclient-option)

Allows you to customise the http client used by the adapter.

## Examples

Below is example code for using the adapter.

### Track using default parameters:

```go
package main

import (
	"fmt"

	"github.com/AreaHQ/go-heapanalytics"
)

func main() {
	// test id used in documentation
	appID := "11"

	c := heapanalytics.NewClient(appID)

	identity := "my@identifier.net"
	event := "test_event"
	properties := map[string]interface{}{"test_prop": "test_val"}

	err := c.Track(identity, event, properties)
	if err != nil {
		fmt.Printf("Error sending track: %s", err.Error())
	} else {
		fmt.Println("Event sent.")
	}
}
```

### Track with HttpClient Option:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/AreaHQ/go-heapanalytics"
)

func main() {
	// test id used in documentation
	appID := "11"

	// this may be a custom client that you're reusing
	httpClient := http.DefaultClient

	c := heapanalytics.NewClient(appID, heapanalytics.HttpClient(httpClient))

	identity := "my@identifier.net"
	event := "test_event"
	properties := map[string]interface{}{"test_prop": "test_val"}

	err := c.Track(identity, event, properties)
	if err != nil {
		fmt.Printf("Error sending track: %s", err.Error())
	} else {
		fmt.Println("Event sent.")
	}
}
```

### Add user properties using default parameters

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/AreaHQ/go-heapanalytics"
)

func main() {
	// test id used in documentation
	appID := "11"

	// this may be a custom client that you're reusing
	httpClient := http.DefaultClient

	c := heapanalytics.NewClient(appID, heapanalytics.HttpClient(httpClient))

	identity := "my@identifier.net"
	properties := map[string]interface{}{"test_prop": "test_val"}

	err := c.AddUserProperties(identity, properties)
	if err != nil {
		fmt.Printf("Error sending track: %s", err.Error())
	} else {
		fmt.Println("User properties sent.")
	}
}

```
