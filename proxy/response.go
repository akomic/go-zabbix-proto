package proxy

import (
	"encoding/json"
	"fmt"
)

// ProxyResponse class.
type ProxyResponse struct {
	Data     []uint8 `json:"data"`
	Response string  `json:"response"`
	Info     string  `json:"info"`
	JSON     string  `json:"json"`
}

// ProxyResponse class constructor.
func NewProxyResponse(data []uint8) (r *ProxyResponse, err error) {
	if len(data) < 13 {
		err = fmt.Errorf("NewProxyResponse Input data to short")
		return
	}
	jsonData := data[13:]

	r = &ProxyResponse{Data: data, JSON: string(jsonData)}
	err = json.Unmarshal(jsonData, r)
	if err != nil {
		err = fmt.Errorf("Error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			err = fmt.Errorf("%s ; Syntax error at byte offset %d", err, e.Offset)
		}
		return
	}
	return
}

// ProxyConfigResponse class.
type ProxyConfigResponse struct {
	Globalmacro struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	} `json:"globalmacro"`
	Hosts struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	} `json:"hosts"`
	Interface struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"interface"`
	HostsTemplates struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"hosts_templates"`
	Hostmacro struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"hostmacro"`
	Items struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"items"`
	Drules struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"drules"`
	Dchecks struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"dchecks"`
	Regexps struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	} `json:"regexps"`
	Expressions struct {
		Fields []string        `json:"fields"`
		Data   [][]interface{} `json:"data"`
	} `json:"expressions"`
	Groups struct {
		Fields []string `json:"fields"`
		Data   [][]int  `json:"data"`
	} `json:"groups"`
	Config struct {
		Fields []string `json:"fields"`
		Data   [][]int  `json:"data"`
	} `json:"config"`
	Httptest struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"httptest"`
	Httptestitem struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"httptestitem"`
	Httpstep struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"httpstep"`
	Httpstepitem struct {
		Fields []string      `json:"fields"`
		Data   []interface{} `json:"data"`
	} `json:"httpstepitem"`
	Data     []uint8 `json:"data"`
	Response string  `json:"response"`
	Info     string  `json:"info"`
}

// ProxyConfigReponse class constructor.
func NewProxyConfigResponse(data []uint8) (r *ProxyConfigResponse, err error) {
	jsonData := data[13:]

	r = &ProxyConfigResponse{Data: data}
	err = json.Unmarshal(jsonData, r)
	if err != nil {
		err = fmt.Errorf("Error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			err = fmt.Errorf("%s ; Syntax error at byte offset %d", err, e.Offset)
		}
		return
	}
	return
}

func (response *ProxyConfigResponse) GetHosts() (hosts map[float64]Host) {
	macros := make(map[float64][]Hostmacro)
	for _, m := range response.Hostmacro.Data {
		macro := m.([]interface{})
		if len(macro) == 4 {
			macros[macro[1].(float64)] = append(macros[macro[1].(float64)], Hostmacro{
				Hostmacroid: macro[0].(float64),
				Macro:       macro[2].(string),
				Value:       macro[3].(string),
			})
		}
	}

	items := make(map[float64][]Item)
	for _, i := range response.Items.Data {
		item := i.([]interface{})
		if len(item) == 32 {
			items[item[4].(float64)] = append(items[item[4].(float64)], Item{
				Itemid:                item[0].(float64),
				Type:                  item[1].(float64),
				Snmp_community:        item[2].(string),
				Snmp_oid:              item[3].(string),
				Key_:                  item[5].(string),
				Delay:                 item[6].(float64),
				Status:                item[7].(float64),
				Value_type:            item[8].(float64),
				Trapper_hosts:         item[9].(string),
				Snmpv3_securityname:   item[10].(string),
				Snmpv3_securitylevel:  item[11].(float64),
				Snmpv3_authpassphrase: item[12].(string),
				Snmpv3_privpassphrase: item[13].(string),
				Lastlogsize:           item[14].(float64),
				Logtimefmt:            item[15].(string),
				Delay_flex:            item[16].(string),
				Params:                item[17].(string),
				Ipmi_sensor:           item[18].(string),
				Data_type:             item[19].(float64),
				Authtype:              item[20].(float64),
				Username:              item[21].(string),
				Password:              item[22].(string),
				Publickey:             item[23].(string),
				Privatekey:            item[24].(string),
				Mtime:                 item[25].(float64),
				Flags:                 item[26].(float64),
				Port:                  item[28].(string),
				Snmpv3_authprotocol:   item[29].(float64),
				Snmpv3_privprotocol:   item[30].(float64),
				Snmpv3_contextname:    item[31].(string),
			})
		}
	}

	hosts = make(map[float64]Host)
	for _, d := range response.Hosts.Data {
		if len(d) == 14 {
			if d[2].(float64) != 3 {
				hostid := d[0].(float64)
				hosts[hostid] = Host{
					Hostid:           hostid,
					Host:             d[1].(string),
					Status:           d[2].(float64),
					Ipmi_authtype:    d[3].(float64),
					Ipmi_privilege:   d[4].(float64),
					Ipmi_username:    d[5].(string),
					Ipmi_password:    d[6].(string),
					Name:             d[7].(string),
					Tls_connect:      d[8].(float64),
					Tls_accept:       d[9].(float64),
					Tls_issuer:       d[10].(string),
					Tls_subject:      d[11].(string),
					Tls_psk_identity: d[12].(string),
					Tls_psk:          d[13].(string),
					Macros:           macros[hostid],
					Items:            items[hostid],
				}
			}
		}
	}
	return hosts
}
