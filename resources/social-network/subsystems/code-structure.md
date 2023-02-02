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
subsystems/
  for each <subsystem>/:
    gateway/ - gRPC API used to comunicate outside the subsystem
    for each <service>/:
      api/ - gRPC API used by other services in the same subsystem or gateway  
      core/ - internal implementation 
      config/ - parameters used to manipulate service behaviour or communications related
      errors/ - error wrappers
      db/ (?) - auto generated code for database
      events/ (?) - publish/subscribe things 
      proto/ (?) - auto generated code for communications via gRPC
```
