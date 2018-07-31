package deploy

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/gobuffalo/packr"
	"golang.org/x/oauth2"
	"os"
	"strconv"
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
func getSSHPubKey() {

}
func spinUpNewDroplet(name string, region string, standard bool, numVCPU int, memoryGigs int, token string) {

	tokenSource := &TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)
	slug := createSlugString(standard, numVCPU, memoryGigs)
	fmt.Println("slug "+slug)
	userDataBox := packr.NewBox("./user-data.yaml")
	sshKey := godo.DropletCreateSSHKey{Fingerprint:"51:42:0e:45:50:0c:57:39:8f:bd:86:13:9c:29:3e:84"}
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region, //"nyc3",
		Size:   slug,
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-18-04-x64",
		},
		UserData: userDataBox.String("user-data.yaml"),
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
