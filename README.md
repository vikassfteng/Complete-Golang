# Complete-Golang
## KAFKA
Apache Kafka is an open-source distributed event streaming platform. It’s used to build real-time data pipelines and streaming applications.
* Producers send data (messages) into topics.
* Consumers read data from those topics.
* The messages are stored persistently, and consumers can read them at their own pace.


--- Each topic is split into partitions for parallelism and scalability.
--- More partitions = more consumers can process data in parallel.


-----------Kafka Brokers---------

A broker is a Kafka server that stores data and handles requests.
Kafka runs in a cluster of brokers — they distribute and replicate partitions.
Example: If you have 3 brokers and a topic with 6 partitions, the partitions will be spread across the brokers.


-----------Kafka Replication---------- 

Kafka ensures high availability by replicating partitions across brokers.
One broker is the leader of a partition; others are followers.
If a broker goes down, another takes over — so no data is lost.


-----------What does a broker do?------------

Receives data (messages) from producers.
Stores data on disk (in partitions).
Serves data to consumers when they request it.
Coordinates with other brokers to manage topics and partitions.



-----------What Happens if a Kafka Server (Broker) Goes Down?------------

If a Broker That Hosts Partition Replicas Goes Down
No immediate data loss (because of replication).
Kafka automatically switches to the in-sync replica (ISR) on another broker.
Producers and consumers are redirected to the new leader for the partition.
✅ Result: Kafka keeps working — no downtime for producing or consuming.


-------------What Happens When All Replicas of a Partition Are Down?---------
Partition Becomes Unavailable
The partition has no leader, and there are no in-sync replicas (ISR).
Kafka marks the partition as offline.
Producers cannot write, and consumers cannot read from that partition.
💥 Basically: That piece of data becomes inaccessible until at least one replica comes back online.

-------------How to Increase Kafka Throughput------------

1. Use Batching
Send messages in batches instead of one-by-one.
Kafka producers have batch.size (bytes) — increase it for better throughput.
"Instead of sending 1,000 individual messages/sec, batch them into groups of 100 for fewer requests."

2. Enable Compression
Use compression.type=gzip or lz4 to reduce message size.
Smaller messages = faster network transfers = higher throughput.

3. Increase Partitions
More partitions → more consumers can read in parallel.
This spreads load across brokers and CPUs.






---------------How Does Event-Driven Architecture Work?---------------

Event Generation:
An event occurs in the system, such as a user submitting a form, a sensor detecting motion, or a payment transaction.

Event Propagation:
The event is placed on an event bus (e.g., Kafka, RabbitMQ), making it available for consumption.

Event Consumption:
A consumer (service, system, etc.) listens for relevant events and processes them. For example, a notification service might listen for a "new user registration" event and send a welcome email.

Event Processing:
Event processors react to events, possibly triggering additional actions or generating new events.