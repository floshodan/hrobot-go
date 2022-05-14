package main

import (
	"context"
	"fmt"
	"os"

	"github.com/floshodan/hrobot-go/hrobot"
)

func main() {

	//export HETZNER_TOKEN=USERNAME:PASSWORD
	client := hrobot.NewClient(hrobot.WithToken(os.Getenv("HETZNER_TOKEN")))

	key := "98:a1:3a:00:90:f4:c8:d3:8b:62:61:dc:80:ab:84:03"
	ssh_key, _, _ := client.SSHKey.GetByFingerprint(context.Background(), key)
	fmt.Println(ssh_key.Data)
	//resp, _ := client.SSHKey.Delete(context.Background(), key)
	//fmt.Println(resp.StatusCode)

	//resp, _ := client.SSHKey.Delete(context.Background(), key)
	//fmt.Println(resp.StatusCode)

	ssh, _, _ := client.SSHKey.List(context.Background())

	for _, rec := range ssh {
		fmt.Println(rec.Name)
		fmt.Println(rec.Fingerprint)
		fmt.Println(rec.Data)
	}

	//keydata := "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAII8xXtJ978GpJIxm3SKOGeuHFb3jiiI7E1O+euYlXEfE code@floshodan.io"
	//keydata := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDKGlvtvsTB0TN8yUlOWPa1Th/OudTFEwfs/Kiuvt1HB69UBdkcOTYN7sRX/SybKWH9g0DKvY0SMJxpVX1ejne1juF05U7nqDhBgryVj7u8398F7NYL/oFnaaVG86q2B8RXaJt+BlGZqnOVZ/u7Y/3FEewuCtOG2a401SSPxVANWM9Np9VQl1VnUdfYaevVYQJDKlYgABg8hilmbtx617Md/t+X0RnLHesWgqfCiSTb6EzaXMemQEZYztBWQpijH3aq82CK8Q4cP4dyXwR8bnpRLW33uDfRa6nOIP+w/5nCyIrpZEhyYnh6Z+iFVlMvCTXDXnYUAw3A0BGO8GneeHM6rYTIf1ecqoa+wcgeU/Huyvs4Ou3jZC/ziKQe11FV3EfwNtgbJTf1RkyIreDYfKyyQ8yugx3geGDH/rcM20KLqyEHZaSjY9Bq6kvAFW+qq3+tkgpYsn00lsSaFv8bBeM6MC7YZeqqb0jBfQEJTG4fPTSUAN1wZa43SrBXRRyuQt5MK+TOvty54QQIlroD1agS4lrB7IFJOf8U4CxM58x5egZWhCImARrvi1YHyAeVZMd+MGdaaMhwDoIhMOUJIT5oLbuvFcadU0WYWp5+6GjcYNf9hXEatsqRDaAn3xxAKGaxna/PLhKzait7K0HMlyv9tSIX+CUeiSPx4HLZOBCoww== flo@DESKTOP-4M630K4"

	/*data := &hrobot.CreateKeyOpts{
		Name: "linux key",
		Data: keydata,
	}
	*/

	/*		order := &hrobot.OrderMarketOpts{
				Product_id:     "1545696",
				Authorized_key: "98:a1:3a:00:90:f4:c8:d3:8b:62:61:dc:80:ab:84:03",
				Test:           true,
			}
	*/
	//new_key, resp, _ := client.SSHKey.Create(context.Background(), data)

	//fmt.Println(new_key.Name, resp.Status)

}
