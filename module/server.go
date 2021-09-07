package module

import (
	"context"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/antibots/ticket"
)

var (
	ticketClient ticket.TicketClient
)

type Server struct {
	module.UnimplementedModuleServer
}

func init() {
	var err error
	ticketClient, err = sonic.NewTicketClient("localhost:3000")
	if err != nil {
		panic(err)
	}
	return
}

func (s Server) Task(_ context.Context, data *module.Data) (*module.StartResponse, error) {
	//v, _ := json.Marshal(data)
	//fmt.Println(string(v))
	task := NewTask(data)
	if err := task.Start(data); err != nil {
		return nil, err
	}

	return &module.StartResponse{Started: true}, nil
}
