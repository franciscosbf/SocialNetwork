### cmd/

```text
for each <subsystem>/:
  for each  <service>/:
    <service>.go - service launcher
```

### pkg/

```text
for each <subsystem>/:
  for each  <service>/:
    |init calls used for service startup| 
```

### internal/

```text
clis/ - for databases, kafka, redis, rabbitmq, ...
common/ - contains generic implementations, interfaces...
subsystems/
  for each <subsystem>/:
    gateway/ - gRPC API used to comunicate outside the subsystem
    for each <service>/:
      api/ - gRPC API used by other services in the same subsystem or gateway  
      core/ - internal implementation 
      errors/ - error codes
      storage/ - persistency things
        db/ (?) - auto generated code for database
        |other related components like caching and bridge between db cli and external access|
      events/ (?) - publish/subscribe things 
      proto/ (?) - auto generated code for communications via gRPC
```
