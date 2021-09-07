package module

import (
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic/base"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"sync"
)

var _ face.ICallback = (*Task)(nil)

type Task struct {
	*base.BTask
	liveJson string
	cParam string
	sizeId string
	csrf string

	ticketHash string
	ticketLocker *sync.Mutex

	color string
	size string
}

func NewTask(data *module.Data) *Task {
	task := &Task{BTask: &base.BTask{Data: data}}
	task.Callback = task
	task.Init()
	return task
}

func (tk *Task) OnInit() {
	return
}

func (tk *Task) OnPreStart() error {
	return nil
}

func (tk *Task) OnStarting() {
	tk.FastClient.CreateCookieJar()
	tk.ticketLocker = &sync.Mutex{}

	tk.Flow()
}
func (tk *Task) OnPause() error {
	return nil
}
func (tk *Task) OnStopping() {
	tk.FastClient.Destroy()
	//panic("")
	return
}

func (tk *Task) Flow() {
	funcarr := []func(){
		tk.ATC,
		tk.GetNtbcc,
		tk.GetCSRF,
		//hcaptcha
		tk.GetCheckout,
	}

	for _, f := range funcarr {
		select {
		case <-tk.Ctx.Done():
			return
		default:
			f()
		}
	}
}
