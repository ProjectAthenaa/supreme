package module

import (
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic/frame"
	"github.com/ProjectAthenaa/supreme/config"
	"strings"
)

func (tk *Task) FormatPhone() string {
	return tk.Data.Profile.Shipping.PhoneNumber[:3] + "+" + tk.Data.Profile.Shipping.PhoneNumber[3:6] + "+" + tk.Data.Profile.Shipping.PhoneNumber[6:]
}

func (tk *Task) AwaitMonitor() {
	pubsub, err := frame.SubscribeToChannel(tk.Data.Channels.MonitorChannel)
	if err != nil {
		tk.Stop()
		return
	}
	defer pubsub.Close()

	tk.SetStatus(module.STATUS_MONITORING)
	for data := range pubsub.Chan(tk.Ctx) {
		if containsI(tk.Data.Metadata[*config.Module.Fields[3].FieldKey], data["size"].(string)) && containsI(tk.Data.Metadata[*config.Module.Fields[2].FieldKey], data["color"].(string)){
			tk.sizeId = data["sizeId"].(string)
			tk.PID = data["pid"].(string)
			break
		}
	}
	tk.SetStatus(module.STATUS_PRODUCT_FOUND)
}


func containsI(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}
