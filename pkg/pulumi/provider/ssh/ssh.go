package ssh

import (
	// "fmt"
	// "time"

	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

// run script to install k8s cluster inside of instance
func InstallDependencies(ip string, user string, privateKey string) (string, error) {
	// parse ssh key to signer
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return "", err
	}

	// define config to use in ssh connection
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// connect to the server
	addr := fmt.Sprintf("%s:22", ip)

	// while for try connect to server
	var kubeConfig string
	attempts := 0
	for {
		// connect to server
		client, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			attempts++
			log.Printf("attempt: %d, error: %+v \n", attempts, err.Error())
			time.Sleep(5 * time.Second)

			continue
		}
		defer client.Close()

		// install golang
		err = InstallGolang(client)
		if err != nil {
			return "", err
		}

		// install docker
		err = InstallDocker(client)
		if err != nil {
			return "", err
		}

		// install k3s cluster
		kubeConfig, err = InstallK3s(client, ip)
		if err != nil {
			return "", err
		}

		break
	}

	return kubeConfig, nil
}
