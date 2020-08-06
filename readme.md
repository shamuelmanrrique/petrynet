# Distributed systems Reliability, Causation and Group Delivery

The finality of this project is to achieve the objectives of practice number one in subject Net and Distributed System at Zaragoza University.

## The project structure is:

```
reports     -->   This folder contains project specification and requirements.
  src         -->   This folder contains all code about the project.
        centralsim          -->  
        cmd                 -->  
        distconssim         -->  
        utils               -->
  .gitignore  -->   File indicate files or folder to ignore
  readme.md   -->   Describe all require information you let to know about the project
```

## The objectives of the project are learn and understand:

* 
* 

# Installation

This project requires:

```
go (>= 1.13)
```

Other library used:

* [vclock]()
* [go-multicast]()

# Source code

You can check the latest sources with the command:

> git clone 

**It's very important set correct path to run project or clone repository in folder "/home/userName/go/src/"**

# Copiar ssh a remote

> ssh-copy-id -i ~/.ssh/id_rsa smmanrrique@localhost

# Execute main using one o this mode [TCP,UDP, CHANDY]

For execute main go program yo must use follow flag:

* name  --> Insert name like machine# (# is a number 1-3)
* mode  --> Mode to execute [tcp, udp, chandy] | default tcp
* log   --> With true Send output to log file otherwise print on terminal | default false

You need to open one terminal by every machine and execute go script in this order.

## machina3

> go run main.go -name "machine3" -mode "tcp" -log true

## machina2

> go run main.go -name "machine2"  -mode "tcp" -log true

## machina1

> go run main.go -name "machine1" -mode "tcp" -log true


# Execute Test

> ssh-copy-id -i ~/.ssh/id_rsa smmanrrique@localhost
















# petrynet

# Excecute remote test
/usr/local/go/bin/go test /home/a802400/go/src/sd_petry_nets/src/distconssim -run TestSubNet0
/usr/local/go/bin/go test /home/a802400/go/src/sd_petry_nets/src/distconssim -run TestSubNet1
/usr/local/go/bin/go test /home/a802400/go/src/sd_petry_nets/src/distconssim -run TestSubNet2



go run main.go -i "127.0.1.1:5000" -n "TestSubNetL0" 
go run main.go -i "127.0.1.1:5001" -n "TestSubNetL1"  
go run main.go -i "127.0.1.1:5002" -n "TestSubNetL2"  