package utils

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func InitSSH(addr string) *ssh.Client {
	// func InitSSH(addr string) (ssh.Session, error) {
	IDRsa := "/home/shamuel/.ssh/id_rsa"
	var user = "a802400"
	// var user = "shamuel"

	key, err := ioutil.ReadFile(IDRsa)
	if err != nil {
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		panic(err.Error())
	}

	// // Create a session. It is one session per command.
	// session, err := client.NewSession()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// return *session, err
	return client

}

func ExcecuteSSH(cmd string, conn *ssh.Client) {
	sess, err := conn.NewSession()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	sessStdOut, err := sess.StdoutPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stdout, sessStdOut)
	sessStderr, err := sess.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)
	log.Println(cmd)
	err = sess.Run(cmd)
	if err != nil {
		panic(err)
	}
}
