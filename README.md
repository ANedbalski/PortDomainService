#Port Domain Service

This project is made as a part of test task

## Architectural solution
As incoming file can contain gigabytes of data, the REST service is accepting the file and save it to the storage. After that it is firing task to parse it.

Another process of application is parsing daemon. It is running and awaiting a task to parse the data.
In real world it will be implemented using Kafka or RabbitMQ or other. In the test task it is made with chan.
In this case the parsing process is launching with serve process as a goroutine

to run the REST server:
>ports s -c ports.yml

To run the parsing process
>ports p -c ports.yml

##