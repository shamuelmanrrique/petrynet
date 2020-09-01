# Petry Net

The finality of this project is learn about Petry Networks in subject Net and Distributed System at Zaragoza University.

## The project structure is:

```
reports     -->   This folder contains project specification and requirements.
  src         -->   This folder contains all code about the project.
        centralsim          -->  
        cmd                 -->  
        config              -->  
        distconssim         -->  
        logs                -->  
        test                -->  
        utils               -->
  .gitignore  -->   File indicate files or folder to ignore
  readme.md   -->   Describe all require information you let to know about the project
```

## The objectives of the project are:

* Design and implement a distributed Petri nets simulator.
* Implement conservative synchronization through the LEFs mechanism to
distributed petri nets.

# Installation

This project requires:

```
go (>= 1.13)
```

# Source code

You can check the latest sources with the command:

> git clone https://github.com/smmanrrique/sd_petry_nets.git

**It's very important set correct path to run project or clone repository in folder "/home/userName/go/src/"**

# Copiar ssh a remote

> ssh-copy-id -i ~/.ssh/id_rsa remoteName@localhost

# Execute main from Test

> ~/go/src/sd_petry_nets/src/test && go test -v -run TestSSH  

# Execute Test with three subNets

> ~/go/src/sd_petry_nets/src/test && go test -v -run TestDist  

# Execute Test with five subNets

> ~/go/src/sd_petry_nets/src/test && go test -v -run Test5Dist  

# Execute Test with five subNets and different times

> ~/go/src/sd_petry_nets/src/test && go test -v -run TestTime5Dist  