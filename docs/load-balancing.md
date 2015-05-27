# Load balancing

The best solution that I have managed to come across at this point - justcontainers/nginx-loadbalancer & justcontainers/loadbalancer-feeder - is a pod-aware reverse proxy load balancer with SSL support.

There are a few issues which I would like to address:
 - SSL cert management relies on having the cert files on the load balancer (which I have addressed).
 - Generation of hosts is not clear
 - Generation of server endpoints relies on port 8080 being exposed on container
 - Reliance on metadata name for labelling is not desirable
 - No explanation of how a host is detected and linked
 - Web service on endpoint containers not covered