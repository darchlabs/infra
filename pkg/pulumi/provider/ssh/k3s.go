package ssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func InstallK3s(c *ssh.Client, ip string) (string, error) {
	fmt.Print("Installing k3s cluster... ")

	// session for install k3s cluster
	session, err := c.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// install k3s cluster
	command := fmt.Sprintf("curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC=\"--tls-san %s\" sh -s -", ip)
	_, err = session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	// session for change permissions to k3s config file
	session, err = c.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// change permissions to k3s config file
	command = "sudo chmod 644 /etc/rancher/k3s/k3s.yaml"
	_, err = session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	// session for waiting k3s cluster
	session, err = c.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// waiting k3s cluster
	command = "while ! kubectl get nodes &> /dev/null; do sleep 1; done"
	_, err = session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	// session for get kube config
	session, err = c.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// get kube config
	command = "cat /etc/rancher/k3s/k3s.yaml"
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	fmt.Println("ok")
	fmt.Println(string(output))

	return string(output), nil
}
