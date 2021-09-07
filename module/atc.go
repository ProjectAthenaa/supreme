package module

import (
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"strings"
)

func (tk *Task) ATC(){
	tk.SetStatus(module.STATUS_ADDING_TO_CART)
	req, err := tk.NewRequest("POST", fmt.Sprintf("https://www.supremenewyork.com/shop/%s/atc", tk.sizeId), nil)
	if err != nil{
		tk.SetStatus(module.STATUS_ERROR, "could not create atc req")
		tk.Stop()
		return
	}
	res, err := tk.Do(req)
	if err != nil{
		tk.SetStatus(module.STATUS_ERROR, "could not atc")
		tk.Stop()
		return
	}

	if strings.Index(string(res.Body), `"success":true`) > 0{
		tk.SetStatus(module.STATUS_ADDED_TO_CART)
	}
}
