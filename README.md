# Stagways

push system base on REDIS Stream.

```txt

req: time-stamp   -> delay calculate, flow control
     stream-id    -> router info
     stream-type  -> flow control policy
     task-id      -> track info
     msg-data     -> real payload


                                        -> [subscriber-1] -> [handler]
 [req] -> [publisher] -> [redis stream] -> [subscriber-2] -> [handler]
                                        -> [subscriber-3] -> [handler]


steam-id <-> subscriber List, same as `Channel on ProxyList`
```

reference: *https://redis.io/docs/data-types/streams-tutorial*
