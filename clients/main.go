package clients

type Client struct {
	Timestamp string `json:"timestamp,omitempty"`
	Location  string `json:"location,omitempty"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	Unit      string `json:"unit,omitempty"`
}

type ClientBundle struct {
	Timestamp string      `json:"timestamp,omitempty"`
	Fields    []FieldItem `json:"fields,omitempty"`
	// ID        string      `json:"id,omitempty"`
	// ClientID  string      `json:"clientId,omitempty"`
}

type ClientBundles []ClientBundle

type ClientResponseMeta struct {
	Timestamp string `json:"timestamp,omitempty"`
	Status    string `json:"status,omitempty"`
}

type ClientResponse struct {
	// ID       string               `json:"id,omitempty"`
	ClientID string               `json:"clientId,omitempty"`
	Items    []ClientResponseMeta `json:"items,omitempty"`
	Data     map[string]string    `json:"data,omitempty"`
}

type FieldItem struct {
	// Timestamp string `json:"timestamp,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
	Map   string `json:"map,omitempty"`
	// Location  string `json:"location,omitempty"`
}
