package main

import (
	"encoding/json"
	"fmt"

	"github.com/carmel/go-pwa/pkg/webpush"
)

const (
	subscription    = ``
	vapidPublicKey  = ""
	vapidPrivateKey = ""
)

func main() {
	// Decode subscription
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	// Send Notification
	resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
		Subscriber:      "example@example.com", // Do not include "mailto:"
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	if err != nil {
		// TODO: Handle error
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
