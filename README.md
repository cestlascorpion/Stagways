# Stagways

push system base on REDIS Stream and other features(?)

>- multi-producer
>- multi-consumer_group, each consumer_group has only one consumer
>- best-effort delivery and aim to reduce latency/increase throughput
>- dynamically create a large number of streams frequently

```txt

req: time-stamp   -> delay calculate, flow control
     stream-id    -> router info
     stream-type  -> flow control policy
     task-id      -> track info
     msg-data     -> real payload


                                                                      -> [subscriber-1(proxy-1)] -> [handler func]
 [handler func] -> [publisher(im service)] -> [redis stream(channel)] -> [subscriber-2(proxy-2)] -> [handler func]
                                                                      -> [subscriber-3(proxy-3)] -> [handler func]


steam-id <-> subscriber List, same as `Channel on ProxyList`
```

reference: <https://redis.io/docs/data-types/streams-tutorial>
