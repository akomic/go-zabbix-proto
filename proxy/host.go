package proxy

type Host struct {
	Hostid           float64
	Host             string
	Status           float64
	Ipmi_authtype    float64
	Ipmi_privilege   float64
	Ipmi_username    string
	Ipmi_password    string
	Name             string
	Tls_connect      float64
	Tls_accept       float64
	Tls_issuer       string
	Tls_subject      string
	Tls_psk_identity string
	Tls_psk          string
	Macros           []Hostmacro
	Items            []Item
}

type Hostmacro struct {
	Hostmacroid float64
	Hostid      float64
	Macro       string
	Value       string
}

type Item struct {
	Itemid                float64
	Type                  float64
	Snmp_community        string
	Snmp_oid              string
	Key_                  string
	Delay                 float64
	Status                float64
	Value_type            float64
	Trapper_hosts         string
	Snmpv3_securityname   string
	Snmpv3_securitylevel  float64
	Snmpv3_authpassphrase string
	Snmpv3_privpassphrase string
	Lastlogsize           float64
	Logtimefmt            string
	Delay_flex            string
	Params                string
	Ipmi_sensor           string
	Data_type             float64
	Authtype              float64
	Username              string
	Password              string
	Publickey             string
	Privatekey            string
	Mtime                 float64
	Flags                 float64
	Port                  string
	Snmpv3_authprotocol   float64
	Snmpv3_privprotocol   float64
	Snmpv3_contextname    string
}
