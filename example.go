package main

import (
	"context"

	"github.com/antihax/optional"
	"github.com/xu001186/nifi-client/helper"
	"github.com/xu001186/nifi-client/nifi"
)

func main() {

	conf := nifi.NewConfiguration()
	conf.BasePath = "https://yanan001:8443/nifi-api"
	httpClient, err := helper.NewHttpClient(nil, true)
	if err != nil {
		panic(err)
	}
	conf.HTTPClient = httpClient
	client := nifi.NewAPIClient(conf)
	token, _, err := client.AccessApi.CreateAccessToken(context.TODO(), &nifi.AccessApiCreateAccessTokenOpts{
		Username: optional.NewString("5ac49eed-8bc1-4d61-942d-d8f555f42af0"),
		Password: optional.NewString("rg/Kr5ljG/0D8gNn/xr6EkGioAFlhsL1"),
	})
	if err != nil {
		panic(err)
	}
	ctx := context.WithValue(context.Background(), nifi.ContextAccessToken, token)
	process := nifi.ProcessorEntity{
		Revision: &nifi.RevisionDto{
			Version: 0,
		},
		Component: &nifi.ProcessorDto{
			ParentGroupId: "root",
			Name:          "generate_flowfile",
			Type_:         "org.apache.nifi.processors.standard.GenerateFlowFile",
			Position: &nifi.PositionDto{
				X: 0,
				Y: 0,
			},
			Config: &nifi.ProcessorConfigDto{
				ExecutionNode:                    "ALL",
				SchedulingStrategy:               "TIMER_DRIVEN",
				SchedulingPeriod:                 "0 sec",
				ConcurrentlySchedulableTaskCount: 1,
				Properties: map[string]string{
					"File Size":        "0B",
					"Batch Size":       "1",
					"Data Format":      "Text",
					"Unique FlowFiles": "false",
				},

				AutoTerminatedRelationships: []string{
					"success",
				},
			},
		},
	}

	p, _, err := client.ProcessGroupsApi.CreateProcessor(ctx, process, "root")

	if err != nil {
		swagg, ok := err.(nifi.GenericSwaggerError)
		if ok {
			panic(string(swagg.Body()))
		} else {
			panic(err)
		}

	}

	_, _, err = client.ProcessorsApi.GetProcessor(ctx, p.Id)

	if err != nil {
		panic(err)
	}

}
