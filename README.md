# learn-go-kafka

some command line golang kafka clients to test and demo basic usage of different golang kafka client libs 


## confluent  confluent-kafka-go 
github.com/confluentinc/confluent-kafka-go 

## segmentio /kafka-go 
github.com/segmentio/kafka-go 


## get started

open 3 terminal sessions

## start a local kafka instance using docker 
from root dir of this repo :

`> docker-compose up`

this will start a minimal docker based kafka cluster comprised of zookeeper and a single kafka broker.

# to build and run segmentio client libs 
## start a producer
in separate terminal session
```
> cd kafka-go 
> make producer   
```

## start a consumer
in separate terminal session

```
> cd kafka-go 
> make consumer   
```


# to build and run confluent client libs
confluent libs use cgo wrapper to implement the kafka sdk 
You may need to install a working gcc ( yum install devtools (or similar)  )  in your environment to allow the go build of the client lib to suceed. 

## start a producer
in separate terminal session

``` 
> cd confluent 
> make producer config=./config_local.properties
```

## start a consumer
in separate terminal session


``` 
> cd  confluent 
> make consumer config=./config_local.properties
```
