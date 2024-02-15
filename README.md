# Layer-7-Load-balancer

We are creating a layer 7 application load balancer

Here are the following features:

1. Routing Strategy

   - Round Robin (MUST)
   - Least Connections (TODO)
   - IP Hash (TODO)

2. Dynamically add backend servers (MUST)

   - do this with either api endpoints
   - or with pub sub queue/ message broker

3. Perform Health checks on backend servers and remove unhealthy servers (MUST)

4. Rate Limiting Algorithm (MUST-ISH)

   - a specific user can only send 10 requests per minute

5. Logging and Monitoring (MUST-ISH)

   - integration with open telemetry and datadog (TODO)

6. High Availability (TODO)

   - create a kill switch in the load balancer
   - on failure make sure that load balancer comes back alive quickly
   - Figure out How

7. Perform SSL Termination

8. Perform Caching

STEP BY STEP TODO:

- make sure to implement graceful shutdown
- make sure we can dynamically add servers to our load balancers
- add health checks and remove unhealthy servers
- use worker pool(LATER)
- reserach on rate limiting algo
- implement rate limiting algo
- integrate with data dog
- integrate with open telemetry
