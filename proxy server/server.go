/*A proxy server is an intermediary server that sits between client devices (such as computers or smartphones) and other servers (typically web servers or other internet resources). It acts as a gateway for requests from clients seeking resources from these servers. Here's how a proxy server works and its various purposes:

1. **Request Forwarding**: When a client device, such as a web browser, makes a request to access a specific resource (like a website), it sends the request to the proxy server instead of directly to the target server.

2. **Intermediary**: The proxy server then evaluates the request, and based on its configuration and rules, it decides whether to forward the request to the target server or handle it in some other way.

3. **Anonymity and Privacy**: Proxy servers can be configured to provide anonymity and privacy for users. When a user accesses a website through a proxy server, the target server sees the IP address of the proxy server rather than the user's actual IP address, helping to mask the user's identity.

4. **Content Filtering**: Organizations often use proxy servers to filter and control the content that employees can access on the internet. This can help enforce security policies, block malicious websites, or restrict access to certain types of content.

5. **Caching**: Proxy servers can store copies of frequently accessed web pages and resources. When a client requests a resource that the proxy server has cached, it can serve the cached copy instead of fetching it from the target server. This can reduce bandwidth usage and improve response times.

6. **Load Balancing**: Proxy servers can distribute incoming requests to multiple backend servers, helping to balance the load and ensure that no single server becomes overwhelmed with traffic. This is often referred to as a "reverse proxy."

7. **Security**: Proxy servers can enhance security by acting as a barrier between the internet and an organization's internal network. They can inspect and filter incoming traffic for malicious content, helping to protect against cyber threats.

8. **Access Control**: Organizations can use proxy servers to control who can access specific resources on the internet. They can enforce access policies based on user credentials or other factors.

9. **Bypassing Geographical Restrictions**: Some users use proxy servers to bypass geographical restrictions imposed by websites or services. By connecting to a proxy server in a different location, they can appear as if they are accessing the internet from that location, potentially gaining access to region-restricted content.

10. **Logging and Monitoring**: Proxy servers can log all incoming and outgoing traffic, allowing administrators to monitor and analyze network activity for security and performance reasons.

Proxy servers come in various types and can serve different purposes, depending on how they are configured and used. They play a crucial role in network management, security, and privacy in both personal and organizational contexts.


The go get command can fetch dependencies from a module proxy server without touching the original server that hosts the code 

The Go Command (go get) try to fetch modules from the proxy server
When the code is not available on the proxy server, Go can download it directly from the hosting server.

The proxy configuration is set into the environment variable GOPROXY.

By typing the command

go env GOPROXY
you can check its current value. At the time of writing, the default value is :

https://proxy.golang.org,direct


What is the name of the environment variable that controls the Go Module Proxy usage?

GOPROXY.
How to disable the usage of module proxies in a go get command?

$ GOPROXY="off" go get gitlab.com/loir402/foo