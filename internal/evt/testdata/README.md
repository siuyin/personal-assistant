# NATS configuration and startup scripts

The scripts below startup the main and leaf NATS servers.

`nats-server` binary must be in the path.

```
main-start.sh

leaf-start.sh
```

The `local` connection is to the leafnode having the `leaf` domain.
The leafnode can be intermittently connected to the `hub` domain.

In order for the hub to get stream updates from leafnodes when they connect:
```txt
nats --server nats://{account name}:{account password}@{host:port} \
  stream add --js-domain hub --source {leafnode stream name}

example:
nats --server nats://acc:acc@localhost:4222 stream add --js-domain hub --source pa
```

**Important** Answer Y to the following question:
```
Import "pa" from a different JetStream domain (y/N)
```

specify the domain `leaf` in the above example.


You may now operate on the stream thus:
```txt
nats --server nats://acc:acc@localhost:4222 stream report

nats --server nats://acc:acc@localhost:4222 stream ls
nats --server nats://acc:acc@localhost:4222 stream view pa
```

And you may list server information using the system-account (user: admin, password: admin)
```
nats --server nats://admin:admin@localhost:4222 server ls
```
