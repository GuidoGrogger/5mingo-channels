# 5minGo-channels

Part of our team 5minGo series. Example for using channels in concurrent processing. 

## Description
Creates many shoe sales agents that run in parallel and sell one shoe per second in average per agent. 
Stops when total sales goal is reached. 
With a sales goal of 300 and 10 sales agents the process will run approximately 30 seconds

Experiment with different amount of parallel agents and see how many are possible on your system. 
You will be impressed.

## Additional facts
- Channels are blocking by default
- Sending a message to a channel blocks until the message is read
- Reading a message from a channel blocks until a message is available
- Deadlocks can happen and will be detected
- Channels should be always closed by the sender side not the receiver 
