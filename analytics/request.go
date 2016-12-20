package analytics

import ga "google.golang.org/api/analytics/v3"
import "fmt"

func Request() {
	manage := ga.NewManagementService(Service)
	c, err := manage.Accounts.List().Do()
	if nil != err {
		fmt.Printf("%v", err)
	}
	for _, item := range c.Items {
		fmt.Printf("ID: %s Name: %v\n", item.Id, item.Name)
	}
}
