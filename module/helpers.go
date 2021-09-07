package module

import (
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic/frame"
)

func (tk *Task) FormatPhone() string{
	return tk.Data.Profile.Shipping.PhoneNumber[:3] +"+"+ tk.Data.Profile.Shipping.PhoneNumber[3:6] +"+"+ tk.Data.Profile.Shipping.PhoneNumber[6:]
}

func (tk *Task) AwaitMonitor() {
	pubsub, err := frame.SubscribeToChannel(tk.Data.Channels.MonitorChannel)
	if err != nil{
		tk.Stop()
		return
	}
	defer pubsub.Close()

	tk.SetStatus(module.STATUS_MONITORING)
	for data := range pubsub.Chan(tk.Ctx){
		if tk.size == data["size"].(string) && tk.color == data["color"].(string){
			tk.sizeId = data["sizeId"].(string)
			break
		}
	}
	tk.SetStatus(module.STATUS_PRODUCT_FOUND)
}