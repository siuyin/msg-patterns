# Go concurrency patters

1. fan-in: many into one
1. fan-out: one to many
1. work-queue: a queue that expands on demand to memory limits
1. work-queue-buffered_channel: bounded (pre-allocated) work queue buffer
1. stateless-server: stateless server which launches a goroutine per request
1. stateful-server: sateful server that is single threaded
