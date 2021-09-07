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
		if (tk.Data.TaskData.RandomSize || existInSlice(data["size"].(string), tk.Data.TaskData.Size)) && (tk.Data.TaskData.RandomColor|| existInSlice(data["color"].(string), tk.Data.TaskData.Color)){
			tk.sizeId = data["sizeId"].(string)
			tk.PID = data["pid"].(string)
			break
		}
	}
	tk.SetStatus(module.STATUS_PRODUCT_FOUND)
}

func existInSlice(keyword string, stringSlice []string)bool{
	for _, substr := range stringSlice{
		if substr == keyword{
			return true
		}
	}
	return false
}