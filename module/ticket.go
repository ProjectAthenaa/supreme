package module

import (
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic/antibots/ticket"
)

func (tk *Task) ProcessTicket(){
	tk.ticketLocker.Lock()
	defer tk.ticketLocker.Unlock()
	hash, err := ticketClient.Deobfuscate(tk.Ctx, &ticket.DeobfuscateRequest{Proxy: *tk.FormatProxy()})
	if err != nil{
		tk.SetStatus(module.STATUS_ERROR, "couldnt get ticket hash")
		tk.Stop()
		return
	}
	tk.ticketHash = hash.Value
}

func (tk *Task) GetNtbcc(){
	tk.ticketLocker.Lock()
	defer tk.ticketLocker.Unlock()

	cookie, err := ticketClient.GenerateCookie(tk.Ctx, &ticket.GenerateCookieRequest{
		Proxy: *tk.FormatProxy(),
		Hash:  tk.ticketHash,
	})

	if err != nil{
		tk.SetStatus(module.STATUS_ERROR, "couldnt get ntbcc")
		tk.Stop()
		return
	}

	tk.FastClient.Jar.Set("ntbcc", cookie.Value)
}