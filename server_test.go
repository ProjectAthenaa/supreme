package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/protos/monitor"
	monitor_controller "github.com/ProjectAthenaa/sonic-core/protos/monitorController"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
	module2 "github.com/ProjectAthenaa/supreme/module"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
	"time"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func init() {
	//go debug.StartShapeServer()
	lis = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	module.RegisterModuleServer(server, module2.Server{})
	go func() {
		server.Serve(lis)
	}()
}

func TestModule(t *testing.T) {
	subToken, controlToken, monitorChannel := uuid.NewString(), uuid.NewString(), uuid.NewString()

	ip := "localhost"
	port := "8866"

	core.Base.GetRedis("cache").Publish(context.Background(), fmt.Sprintf("proxies:%s", product.SiteSupreme), fmt.Sprintf(`%s:%s`, ip, port))

	tk := &module.Data{
		TaskID: uuid.NewString(),
		Profile: &module.Profile{
			Email: "poprer656sad@gmail.com",
			Shipping: &module.Shipping{
				FirstName:   "Omar",
				LastName:    "Hu",
				PhoneNumber: "6463222013",
				ShippingAddress: &module.Address{
					AddressLine:  "7004 JFK BLVD E",
					AddressLine2: nil,
					Country:      "US",
					State:        "NEW JERSEY",
					City:         "WEST NEW YORK",
					ZIP:          "07093",
					StateCode:    "NJ",
				},
				BillingAddress: &module.Address{
					AddressLine:  "7004 JFK BLVD E",
					AddressLine2: nil,
					Country:      "US",
					State:        "NEW JERSEY",
					City:         "WEST NEW YORK",
					ZIP:          "07093",
					StateCode:    "NJ",
				},
				BillingIsShipping: true,
			},
			Billing: &module.Billing{
				Number:          "4207670236068972",
				ExpirationMonth: "05",
				ExpirationYear:  "25",
				CVV:             "997",
			},
		},
		Proxy: &module.Proxy{
			//Username: &username,
			//Password: &password,
			IP:   ip,
			Port: port,
		},
		TaskData: &module.TaskData{
			RandomSize:  false,
			RandomColor: false,
			Color:       []string{"White"},
			Size:        []string{"Large"},
		},
		Channels: &module.Channels{
			UpdatesChannel:  subToken,
			CommandsChannel: controlToken,
			MonitorChannel:  monitorChannel,
		},
	}

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	client := module.NewModuleClient(conn)
	_, err = client.Task(context.Background(), tk)
	if err != nil {
		t.Fatal(err)
	}

	//todo: correct inputs to monitor
	conn, err = grpc.Dial("localhost:4000", grpc.WithInsecure())
	monitorClient := monitor.NewMonitorClient(conn)
	monitorClient.Start(context.Background(), &monitor_controller.Task{
		Site:         string(product.SiteSupreme),
		Lookup:       &monitor_controller.Task_Other{Other: true},
		RedisChannel: monitorChannel,
		Metadata:     tk.Metadata,
	})

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))

	t.Log("connecting to redis")
	pubsub := core.Base.GetRedis("cache").Subscribe(ctx, fmt.Sprintf("tasks:updates:%s", subToken))
	t.Log("connected to redis")

	for msg := range pubsub.Channel() {
		var data module.Status
		_ = json.Unmarshal([]byte(msg.Payload), &data)
		fmt.Println(data.Status, data.Information["message"])
		if data.Status == module.STATUS_STOPPED {
			return
		}
	}
}
