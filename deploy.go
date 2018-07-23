package deploy

import (
	"github.com/digitalocean/godo"
	"fmt"
	"context"
	"golang.org/x/oauth2"
)

  =

func getBinLocation() {

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


func createDOClient(token string) *godo.Client {
	tokenSource := &TokenSource{
		AccessToken: token,
	}
	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)
	context.Background()
	return client
}

func spinUpNewDroplet(name string, region string, standard bool, numVCPU int, memoryGigs int) {

	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: "nyc3",
		Size:   "s-1vcpu-1gb",
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-14-04-x64",
		},
	}

	ctx := context.TODO()

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)
		return err
	}
	fmt.Println("New Droplet created: ",newDroplet. )

}
