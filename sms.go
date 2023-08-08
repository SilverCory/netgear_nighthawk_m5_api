package netgear_nighthawk_m5_api

import (
	"net/url"
)

// SMSDeleteAll a function to delete all sms.
func (n *NetgearNighthawkM5Api) SMSDeleteAll() error {
	values := url.Values{
		"sms.deleteAll": []string{"1"},
	}

	return n.FormRequest("config", values)
}

// SMSReadID a function to read a sms by id.
func (n *NetgearNighthawkM5Api) SMSReadID(id string) error {
	if id == "" {
		return nil
	}

	values := url.Values{
		"sms.readId": []string{id},
	}

	return n.FormRequest("config", values)
}

// SMSSend a function to send a sms.
func (n *NetgearNighthawkM5Api) SMSSend(number, message string) error {
	values := url.Values{
		"sms.sendMsg.clientId": []string{"webUi"},
		"sms.sendMsg.receiver": []string{number},
		"sms.sendMsg.text":     []string{message},
		"action":               []string{"send"},
	}

	return n.FormRequest("smsSendMsg", values)
}
