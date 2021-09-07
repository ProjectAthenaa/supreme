package module

func (tk *Task) FormatPhone() string{
	return tk.Data.Profile.Shipping.PhoneNumber[:3] +"+"+ tk.Data.Profile.Shipping.PhoneNumber[3:6] +"+"+ tk.Data.Profile.Shipping.PhoneNumber[6:]
}
