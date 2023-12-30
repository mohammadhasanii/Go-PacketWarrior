
# Gopher-PacketWarrior

This Go program allows blocking incoming packets on a router by adding iptables rules. It can be useful in situations where backend traffic is too high and needs to be limited at the network level.

The program connects to the router, logs in using username and password provided by the user, and adds iptables rules to drop packets for specified ports like 80 and 443. This will prevent the router from accepting any new incoming HTTP/HTTPS requests.

The rules are saved permanently and the program logs out after configuring the firewall.

Some use cases where this program would be helpful:

- When a backend system is overloaded due to high traffic and on the verge of crashing. Blocking incoming requests at the router layer can provide relief to the backend.

- When you need to temporarily put a system in maintenance mode but want to avoid downtime. The firewall rules can block new requests so that existing requests can complete without interruption.

- For trial systems that need to be limited to a certain QPS level. The firewall limits can throttle traffic to prevent overload.

So in summary, this Go program allows you to dynamically block incoming packets when you need to control traffic to downstream systems. By adding iptables rules from Go code, the router can be restricted programmatically.

Let me know if you would like me to explain or expand on any part of this description.
# Description
We are not supposed to answer all http requests all the time
Sometimes we need to block the input of our request from the router for a while
This project will help you

![Logo](https://i.pinimg.com/originals/af/d0/b7/afd0b71a1509e022a6b3359085c50789.png)


## Run Project

Tip : This source can only be run on Windows, it does not run on Linux kernels


You need golang version above 1.8 to run

```golang
 go run app.go
```

