# 🔨 PubSub

PubSub is a library that provides a simple pub/sub mechanism with multiple backend storages.

## Features

- **Simple**: Pretty simple APIs.
- **Offline Messages**: The subscriber can receive the messages after restarting or crashing.
- 🔨 **Multiple Storage Providers**: PubSub can support multiple storage providers to store the messages.

## Installation

`go get -u github.com/zyy17/pubsub`

## Quick Start

### Start etcd

For now, pubsub **only supports etcd** as a storage provider. We will support more storage providers in the future.

You can download the [etcd](https://github.com/etcd-io/etcd/releases) and start it.

### Start a publisher

Use `pubsub.StartPublisher()` to start a publisher and publish the message:

```go
publisher, _ := pubsub.StartPublisher(pubsub.DefaultPublisherOptions())
defer publisher.Stop()

// Let's publish your first message!
// Publish message to the 'pubsub' topic.
publisher.Publish(context.Background(), publisher.NewMessage("pubsub", []byte("hello world!")))
```

### Create a subscriber

Use `pubsub.NewSubscriber` to create the subscriber:

```go
subscriber, _ := pubsub.NewSubscriber(pubsub.DefaultSubscriberOptions())
defer subscriber.Close()

// Subscribe the pubsub topic.
subscriber.Subscribe(context.Background(), []string{"pubsub"})

// Let's receive message from the 'pubsub' topic!
for {
        message, _ := subscriber.Recv()
        log.Printf("Received message: %v", message)
}
```

### Run the examples

Use the following commands to compile the `publisher` and `subscriber`(make sure you have already install [`go`](https://go.dev/dl/) and [`protoc`](https://github.com/protocolbuffers/protobuf/releases/tag/v25.1)):

```console
make examples
```

When the building is complete, run the `publisher` and `subscriber` in two different terminals:

```console
# Run the publisher to publish the messages.
./bin/publisher

# Open another terminal to receive the messages.
./bin/subscriber
```

## 🔨 TO-DO List

PubSub is still in **unstable development**. The following features are still in progress:

- [ ] unit tests
- [ ] integration tests
- [ ] limited offline message list
- [ ] message ttl
- [ ] memory backend
- [ ] redis backend
- [ ] sql backend(sqlite,mysql,pg)
