package service

import (
	"context"
	"fmt"
	"log"

	graph "github.com/microsoftgraph/msgraph-beta-sdk-go"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/users"
)

type MSTask struct {
	todoClient *users.ItemTodoRequestBuilder
	client     *graph.GraphServiceClient
}

func NewMSTaskService(grapClient *graph.GraphServiceClient) *MSTask {
	client := grapClient.Me().Todo()
	// r, err := grapClient.Me().ToGetRequestInformation(context.TODO(),nil)
	// log.Println(r.GetUri())
	u, err := grapClient.Me().Get(context.TODO(), &users.UserItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.UserItemRequestBuilderGetQueryParameters{
			Select: []string{"id", "mail"},
		},
	})
	if err != nil {
		log.Println(err.Error())
	}
	name := *u.GetMail()
	log.Println(name)

	c, err := grapClient.Me().Calendars().Get(context.TODO(), nil)
	if err == nil {
		v := c.GetValue()
		for _, val := range v {
			fmt.Println(*val.GetName())
		}
	}
	log.Println(err.Error())

	return &MSTask{
		client:     grapClient,
		todoClient: client,
	}
}

func (m *MSTask) ListTasks(ctx context.Context) {
	req, _ := m.client.Me().Todo().Lists().ToGetRequestInformation(ctx, nil)
	headers := req.Headers
	url, _ := req.GetUri()
	log.Printf("Request URL: %s", url)
	log.Printf("Headers: %v", headers)
	tasks, err := m.client.Me().Todo().Lists().ByTodoTaskListId("AQMkADAwATM0MDAAMS0yYTYxLTZlNmYtMDACLTAwCgAuAAADGh7TII6DP0CJqowzCzgfcAEAJ48nyYsJC0WS0pUCmeiM1wAEpjdOGwAAAA==").Get(ctx, nil)
	log.Println(tasks)
	if err != nil {
		log.Fatalln("ERROR GetList: ", err.Error())
		return
	}

	// lists := tasks.GetValue()``
	// for _, list := range lists {
	// 	name := *list.GetDisplayName()
	// 	fmt.Println(name)
	// }
}
func (m *MSTask) CreateTask(ctx context.Context) {}
func (m *MSTask) GetTask(ctx context.Context)    {}
