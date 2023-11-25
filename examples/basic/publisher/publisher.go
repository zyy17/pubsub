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
	"fmt"
	"log"
	"time"

	"github.com/zyy17/pubsub"
)

func main() {
	opts := pubsub.DefaultPublisherOptions()
	publisher, err := pubsub.StartPublisher(opts)
	if err != nil {
		log.Fatalf("Failed to create publisher: %v", err)
	}
	defer publisher.Stop()

	log.Printf("Starting publisher on %s", opts.ServicePort)

	// Wait for the publisher to start.
	time.Sleep(10 * time.Second)

	ctx := context.Background()

	// Publish 10 messages to the topic "foo" for every 3 seconds.
	for i := 0; i < 10; i++ {
		msg := publisher.NewMessage("pubsub", []byte(fmt.Sprintf("[%d] Hello world!", i)))
		if err := publisher.Publish(ctx, msg); err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}
		time.Sleep(3 * time.Second)
	}
}
