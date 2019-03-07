# Chapter 12. Integrating via Messaging 

## Overview 

- Integrating bounded contexts with asynchronous messaging using NServiceBus and Mass Transit
- Designing and modeling messaging systems around important events in the domain
- Understanding how messaging frameworks work
- Creating architectural diagrams
- Explanation of the conceptual differences between commands and events
- Theory and examples of industry‐standard messaging patterns
- Dealing with eventual consistency
- Monitoring errors in messaging systems
- Monitoring service level agreements (SLAs) in messaging systems

## Messaging Fundamentals
### Message Bus 
- **WHY**: A centralized component transmitting messages would suffer from single point of failure and hurt scaling (any more convincing reasons??)
- **DEFINITION**: a distributed system that has agents running on every component that sends or receives messages—avoiding the need for a centralized single point of failure
- **Benefit**: Decompling bounded contests by connecting each of them to the bus only

> **NOTE** A centralized component that receives and publishes all messages is known as a broker or message broker

### Reliable Messaging 
- **WHY**: Sending messages from one of your bounded contexts to another could be costly if there were no guarantees the message would ever get there, which will degrade user experience
- It's almost impossible to guarantee that a message will always be delivered only once
- Patterns
  - at‐least‐once
  - at‐most‐once
  - only‐once delivery
- At‐least‐once delivery involves retrying messages that failed or no acknowledgement from the receiver was received (appearing to fail). 
  - Based on a pattern known as **store‐and‐forward**

> **NOTE** *Idempotent messages* are messages that can be sent multiple times and only processed once

### Store‐and‐Forward 
- **HOW**
  - Store the message before it is sent
  - If the message reaches the recipient and is acknowledged, the local copy is deleted
  - Otherwise, it is tried again

### Commands and Events
- **COMMANDS**: messages specifying something needs to happen, and will be handled by only one receiver
- **EVENTS**
  - **DEFINITION**: messages signaling something **happened**, without care of receivers/subscribers
  - Advantage: Ability to add new subscribers without changing existing code
- Convention
  - Name commands as instructions that you want to happen
  - Name events in the past tense describing what has happened

### Eventual Consistency 
- **WHY**: In systems where each bounded contexts has its own database, messaging 
  - break big transactions into smaller ones 
  - achieve an **eventually consistent** state with all these smaller txs combined
- One of the most important aspects of eventually consistent systems is managing the **user experience**
- In eventually consistent systems, 
  - users often get an immediate confirmation that the request has been received
  - with further confirmations of following steps later
  
  which can appear to be a degradation of user experience

## Building an E‐Commerce Application with NATS

- Official Mirror of NATS is https://nats.io/

### Designing the System 
#### Domain‐Driven Design
#### Containers Diagrams 
#### Evolutionary Architecture 
### Sending Commands from a Web Application 
#### Creating a Web Application to Send Messages with NServiceBus 
#### Sending Commands 
### Handling Commands and Publishing Events 
#### Creating an NServiceBus Server to Handle Commands 
#### Configuring the Solution for Testing and Debugging 
#### Publishing Events 
#### Subscribing to Events 
### Making External HTTP Calls Reliable with Messaging Gateways 
#### Messaging Gateways Improve Fault Tolerance 
#### Implementing a Messaging Gateway 
#### Controlling Message Retries 
### Eventual Consistency in Practice 
#### Dealing with Inconsistency 
#### Rolling Forward into New States 
### Bounded Contexts Store All the Data They Need Locally 
#### Storage Is Cheap—Keep a Local Copy 
#### Common Data Duplication Concerns 
### Pulling It All Together in the UI 
#### Business Components Need Their Own APIs 
#### Be Wary of Server‐Side Orchestration 
#### UI Composition with AJAX Data 
#### UI Composition with AJAX HTML 
#### Sharing Your APIs with the Outside World 
## Maintaining a Messaging Application 
### Message Versioning 
#### Backward‐Compatible Message Versioning 
#### Handling Versioning with NServiceBus's Polymorphic Handlers 
### Monitoring and Scaling 
#### Monitoring Errors 
#### Monitoring SLAs 
#### Scaling Out 
## Integrating a Bounded Context with Mass Transit 
### Messaging Bridge 
### Mass Transit 
#### Installing and Configuring Mass Transit 
#### Declaring Messages for Use by Mass Transit 
#### Creating a Message Handler 
#### Subscribing to Events 
#### Linking the Systems with a Messaging Bridge 
#### Publishing Events 
#### Testing It Out 
#### Where to Learn More about Mass Transit 