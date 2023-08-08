Netgear Nighthawk M5 MR5200
==========================

## Introduction
An API to interact with the netgear nighthawk M5 MR5200 LTE router.


## Compatibility
This library has only been tested with the Netgear Nighthawk M5 MR5200 LTE router running `NTGX55_12.04.17.00`.

## Usage
`go get github.com/silvercory/netgear_nighthawk_m5_api`

```go
package main

import "github.com/silvercory/netgear_nighthawk_m5_api"

func main() {
	var client = netgear_nighthawk_m5_api.NewNetgearNighthawkM5Api("192.168.1.1")

	// Login
	if err := client.Login("Hunter1"); err != nil {
		panic(err)
	}

	// Get router information
	modelInfo, err := client.ModelInfo()
	if err != nil {
		panic(err)
	}

	// Get the first unread message.
	var messages = modelInfo.Sms.Msgs
	if len(messages) == 0 {
		panic("no messages")
	}

	var message = messages[0]

	// Mark the message as read.
	if err := client.SMSReadID(message.Id); err != nil {
		panic(err)
	}

	// Send a message back to the sender.
	if err := client.SMSSend(message.Sender, "Hello, World!"); err != nil {
		panic(err)
	}

	// Logout
	if err := client.Logout(); err != nil {
		panic(err)
	}
}
```