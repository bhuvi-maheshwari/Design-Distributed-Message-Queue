# In-Memory Distributed Message Queue

## Problem Statement

Design an in-memory distributed queue system similar to Kafka.

## Requirements

- The queue should be in-memory and should not require access to the file system.
- The system must support multiple topics.
- Producers can publish messages to multiple topics.
- Consumers can subscribe to multiple topics.
- The system must be multi-threaded to support parallel production and consumption of messages.
- Consumers should print messages in the format `<consumer_id> received <message>` upon receiving them.

## Input/Output Format

- **Create Topics**: `topic1`, `topic2`
- **Create Producers**: `producer1`, `producer2`
- **Create Consumers**: `consumer1`, `consumer2`, `consumer3`, `consumer4`, `consumer5`
- **Subscription Setup**:
  - All 5 consumers subscribe to `topic1`
  - Consumers `consumer1`, `consumer3`, and `consumer4` also subscribe to `topic2`
- **Publishing Messages**:
  - `producer1` publishes `Message 1` and `Message 2` to `topic1`
  - `producer2` publishes `Message 3` to `topic1`
  - `producer1` publishes `Message 4` to `topic2`
  - `producer2` publishes `Message 5` to `topic2`

## Expectations

- The code should be functional and demonstrable.
- The code should be modular and readable.
- Separation of concerns should be addressed.
- The code should be organized into multiple files (if not using C/C++).
- The code should be easily testable, with a main method for testing.
- (Optional) Unit tests are encouraged.
- No GUI is required.

## Optional Requirements

- **Consumer Groups**: Implement consumer groups where a message can be consumed by only one consumer per group.
