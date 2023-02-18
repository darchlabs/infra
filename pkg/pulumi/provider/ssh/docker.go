package ssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func InstallDocker(c *ssh.Client) error {
	fmt.Print("Installing docker... ")

	// session for install docker
	session, err := c.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// install docker
	command := "sudo yum install -y docker"
	_, err = session.CombinedOutput(command)
	if err != nil {
		return err
	}

	// session for start docker daemon
	session, err = c.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// start docker daemon
	command = "sudo service docker start"
	_, err = session.CombinedOutput(command)
	if err != nil {
		return err
	}

	// session for enable docker daemon
	session, err = c.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// enable docker daemon
	command = "sudo systemctl enable docker"
	_, err = session.CombinedOutput(command)
	if err != nil {
		return err
	}

	// session for waiting docker
	session, err = c.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// waiting docker
	command = "while [ \"$(sudo service docker status | grep -c running)\" -ne 1 ]; do sleep 1; done"
	_, err = session.CombinedOutput(command)
	if err != nil {
		return err
	}

	// session for add user to docker group
	session, err = c.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// add user to docker group
	command = "sudo usermod -aG docker $(whoami)"
	_, err = session.CombinedOutput(command)
	if err != nil {
		return err
	}

	fmt.Println("ok")

	return nil
}
