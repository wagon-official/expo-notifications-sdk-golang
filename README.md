# expo-notifications-sdk-go

Send push notifications to Expo apps using Go.

## Installation

```
go get github.com/wagon-official/expo-notifications-sdk-golang
```

## Usage

```go
package main

import (
    "fmt"
    expo "github.com/wagon-official/expo-notifications-sdk-golang"
)

func main() {
    // To check the token is valid
    pushToken, err := expo.NewExpoPushToken("ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]")
    if err != nil {
        panic(err)
    }

    // Create a new Expo SDK client
    client := expo.NewPushClient(nil)

    // Publish message
    responses, err := client.Publish(
        &expo.PushMessage{
            To: []expo.ExpoPushToken{pushToken},
            Body: "This is a test notification",
            Data: map[string]interface{}{"withSome": "data"},
            Sound: "default",
            Title: "Notification Title",
            Priority: expo.DefaultPriority,
        },
    )

    // Check errors
    if err != nil {
        panic(err)
    }

    // Validate responses
    for _, res := range responses {
        if res.ValidateResponse() != nil {
            fmt.Println(res.PushMessage.To, "failed")
        }
    }
}
```

## License

MIT
