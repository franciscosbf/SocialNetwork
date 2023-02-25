### cmd/

```text
subsystems/
  for each <subsystem>/:
    for each  <service>/:
      <service>.go - service launcher
```

### pkg/

```text
subsystems/
  for each <subsystem>/:
    for each  <service>/:
      |init calls used for service startup| 
```

### internal/

```text
clis/ - for databases, kafka, redis, rabbitmq, ...
  for each <cli>/:
    config/
      |builder.go - to define variables config|
    |client.go - contains client initializer|
config/ - config parser
envvars/ - environment variables things 
  providers/ - variable providers 
subsystems/
  for each <subsystem>/:
    gateway/ - gRPC API used to comunicate outside the subsystem
    for each <service>/:
      api/ - gRPC API used by other services in the same subsystem or gateway  
      core/ - internal implementation 
      errors/ - error codes
      storage/ - persistency things
        repo/ - repository (db things)
          db/ (?) - auto generated code for database
        cache/ - store
      events/ (?) - publish/subscribe things 
      proto/ (?) - auto generated code for communications via gRPC
```
