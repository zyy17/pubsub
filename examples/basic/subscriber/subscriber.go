// Copyright 2023 zyy17 <zyylsxm@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"log"

	"github.com/zyy17/pubsub"
)

func main() {
	opts := pubsub.DefaultSubscriberOptions()
	subscriber, err := pubsub.NewSubscriber(opts)
	if err != nil {
		log.Fatalf("Failed to create subscriber: %v", err)
	}
	defer subscriber.Close()

	log.Printf("Connected to publisher %s", opts.PubSubEndpoint)

	ctx := context.Background()
	if err := subscriber.Subscribe(ctx, []string{"pubsub"}); err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	for {
		message, err := subscriber.Recv()
		if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}
		log.Printf("Received message: %v", message)
	}
}
