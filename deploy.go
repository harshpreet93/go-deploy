package deploy

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
	"os"
	"strconv"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

type AppDeploymentConfig struct {
	AccessToken string
	Region      string
	Standard    bool
	NumVCPU     int
	MemoryGigs  int
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

//create slug from specs eg. s-1vcpu-1gb
func createSlugString(standard bool, numVCPU int, memoryGigs int) string {
	result := "c-"
	if standard {
		result = "s-"
	}
	result += (strconv.Itoa(numVCPU) + "vcpu-")
	result += (strconv.Itoa(memoryGigs) + "gb")
	return result
}
func getSSHPubKeyMD5Signature() (*ssh.PublicKey, error) {
	rawKey, err := ioutil.ReadFile("~.ssh/id_rsa.pub")
	if err != nil {
		fmt.Errorf("error reading ~.ssh/id_rsa.pub", err)
		return nil, err
	}
	pubKey, err := ssh.ParsePublicKey(rawKey)
	if err != nil {
		fmt.Errorf("error parsing valid key", err)
		return nil, err
	}
	return &pubKey, nil
}
func spinUpNewDroplet(name string, region string, standard bool, numVCPU int, memoryGigs int, token string) {

	tokenSource := &TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)
	slug := createSlugString(standard, numVCPU, memoryGigs)
	pubKey, err := getSSHPubKeyMD5Signature()
	
	sshKey := godo.DropletCreateSSHKey{Fingerprint: }
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region, //"nyc3",
		Size:   slug,
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-18-04-x64",
		},
		SSHKeys: []godo.DropletCreateSSHKey{sshKey},
	}

	ctx := context.TODO()

	newDroplet, response, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)
		os.Exit(1)
	}

	fmt.Println("New Droplet created: ", newDroplet)
	fmt.Println("Response: ", response)

}
