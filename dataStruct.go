package netboxsearch

type NetBoxJSON struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	DNSName     string `json:"dns_name"`
	NatOutside  int    `json:"nat_outside"`
	Description string `json:"description"`
	Prefix      string `json:"prefix"`
	Vrf         string `json:"vrf"`
	Site        string `json:"site"`
	Interface   struct {
		ID     int    `json:"id"`
		URL    string `json:"url"`
		Device struct {
			ID          int    `json:"id"`
			URL         string `json:"url"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
		} `json:"device"`
	} `json:"interface"`
	DeviceType struct {
		ID           int `json:"id"`
		Manufacturer struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"manufacturer"`
		Model       string `json:"model"`
		Slug        string `json:"slug"`
		DisplayName string `json:"display_name"`
	} `json:"device_type"`
	DeviceRole struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"device_role"`
	Site struct {
		ID int `json:"id"`
	} `json:"site"`
	PrimaryIP4 struct {
		ID      int    `json:"id"`
		URL     string `json:"url"`
		Address string `json:"address"`
	} `json:"primary_ip4"`
	LocalContextData struct {
		x map[string]interface{} `json:"-"`
	} `json:"config_context"`
}
