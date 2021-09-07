package module

import (
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	csrfRe = regexp.MustCompile(`"csrf-token" content="([^"]+)"`)
)

func (tk *Task) GetCSRF(){
	req, err := tk.NewRequest("GET", "https://www.supremenewyork.com/checkout", nil)
	if err != nil{
		tk.SetStatus(module.STATUS_ERROR, "couldnt create checkout pag req")
		tk.Stop()
		return
	}
	res, err := tk.Do(req)
	if err != nil{
		tk.SetStatus(module.STATUS_ERROR, "couldnt get checkout page")
		tk.Stop()
		return
	}

	tk.csrf = url.QueryEscape(string(csrfRe.FindSubmatch(res.Body)[1]))
}

func (tk *Task) GetCheckout(){
	var addrline2 string
	if tk.Data.Profile.Shipping.ShippingAddress.AddressLine2 != nil{
		addrline2 = *tk.Data.Profile.Shipping.ShippingAddress.AddressLine2
	}

	req, err := tk.NewRequest("GET", fmt.Sprintf(`https://www.supremenewyork.com/checkout.js?authenticity_token=Bm4CabD%%2FIRjYToix4TSi%%2BFhsTzE8%%2FNXkOJVHuzHcZIlt7iYBefl9qhjNpJ2wqqtSpNsATzbu2uQS6clcYG1WUg%%3D%%3D&`+
	`current_time=%s&`+
	`order%%5Bbilling_name%%5D=%s&`+
	`order%%5Bemail%%5D=%s&`+
	`order%%5Btel%%5D=%s&`+
	`order%%5Bbilling_address%%5D=%s&`+
	`order%%5Bbilling_address_2%%5D=%s&`+
	`order%%5Bbilling_zip%%5D=%s&`+
	`order%%5Bbilling_city%%5D=%s&`+
	`order%%5Bbilling_state%%5D=%s&`+
	`order%%5Bbilling_country%%5D=%s&`+
	`same_as_billing_address=1&`+
	`store_credit_id=&`+
	`order%%5Bterms%%5D=1&`+
	`h-captcha-response=&`+
	`cnt=%d`,
	time.Now().Unix(),
	tk.Data.Profile.Shipping.FirstName+"+"+tk.Data.Profile.Shipping.LastName,
	url.QueryEscape(tk.Data.Profile.Email),
	tk.FormatPhone(),
	strings.ReplaceAll(tk.Data.Profile.Shipping.ShippingAddress.AddressLine, " ", "+"),
	strings.ReplaceAll(addrline2, " ", "+"),
	tk.Data.Profile.Shipping.ShippingAddress.ZIP,
	strings.ReplaceAll(tk.Data.Profile.Shipping.ShippingAddress.City, " ", "+"),
	tk.Data.Profile.Shipping.ShippingAddress.StateCode,
	tk.Data.Profile.Shipping.ShippingAddress.Country,
	1 + rand.Intn(2),
	), nil)
	if err != nil{
		tk.SetStatus(module.STATUS_ERROR, "couldnt create checkout submit")
		tk.Stop()
		return
	}
	_, err = tk.Do(req)
	if err != nil{
		tk.SetStatus(module.STATUS_ERROR, "couldnt get checkout submit")
		tk.Stop()
		return
	}
}