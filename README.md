## API gateway

A simple implementation of api gateway pattern

### How to use

Example:
```
./gatewey 
    -listen=127.0.0.1:8181
    
    // service 1
    -endpoint="/s1"
    -dest=http://127.0.0.1:9091
    
    // service 2
    -endpoint="/s2
    -dest=http://127.0.0.1:9092
    
    // another services...
```

Where:
* listen - it's our gateway address
* endpoint - it's a endpoint for service called _s1_
* dest - it's a destination / real address of given service _s1_

So everything which is sending to the our gateway on _/s1_ route
will forward to the destination address specified in the _dest_ param.
You can use combination of _endpoint_ and _dest_ param many times.


