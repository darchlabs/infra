package ssh

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func InstallGolang(c *ssh.Client) error {
	fmt.Print("Installing golang... ")

	// session for install golang
	session, err := c.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// install golang
	commands := []string{
		"wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz",
		"tar -xvf go1.16.4.linux-amd64.tar.gz",
		"sudo mv go /usr/local",
	}
	_, err = session.CombinedOutput(strings.Join(commands[:], "; "))
	if err != nil {
		return err
	}

	// session for set ENV values for golang
	session, err = c.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// set ENV values for golang
	commands = []string{
		"echo \"export GOPATH=$HOME/go\" >> ~/.bashrc",
		"echo \"export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin\" >> ~/.bashrc",
		"source ~/.bashrc",
	}
	_, err = session.CombinedOutput(strings.Join(commands[:], "; "))
	if err != nil {
		return err
	}

	fmt.Println("ok")

	return nil
}
