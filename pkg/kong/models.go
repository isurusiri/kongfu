package kong

// Service represents a Kong service object
type Service struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Protocol     string `json:"protocol"`
	Path         string `json:"path,omitempty"`
	Retries      int    `json:"retries,omitempty"`
	ConnTimeout  int    `json:"connect_timeout,omitempty"`
	WriteTimeout int    `json:"write_timeout,omitempty"`
	ReadTimeout  int    `json:"read_timeout,omitempty"`
}

// Route represents a Kong route object
type Route struct {
	ID            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Protocols     []string `json:"protocols"`
	Methods       []string `json:"methods,omitempty"`
	Hosts         []string `json:"hosts,omitempty"`
	Paths         []string `json:"paths,omitempty"`
	HTTPSRedirect bool     `json:"https_redirect_status_code,omitempty"`
	RegexPriority int      `json:"regex_priority,omitempty"`
	StripPath     bool     `json:"strip_path,omitempty"`
	PreserveHost  bool     `json:"preserve_host,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	ServiceID     string   `json:"service,omitempty"`
}

// Plugin represents a Kong plugin configuration
type Plugin struct {
	ID         string                 `json:"id,omitempty"`
	Name       string                 `json:"name"`
	ServiceID  string                 `json:"service,omitempty"`
	RouteID    string                 `json:"route,omitempty"`
	ConsumerID string                 `json:"consumer,omitempty"`
	Config     map[string]interface{} `json:"config"`
	Enabled    bool                   `json:"enabled"`
	Tags       []string               `json:"tags,omitempty"`
}

// Configuration represents a complete Kong configuration
type Configuration struct {
	Services []Service `json:"services"`
	Routes   []Route   `json:"routes"`
	Plugins  []Plugin  `json:"plugins"`
}

// GetLiveConfiguration retrieves the current configuration from a Kong Admin API
func GetLiveConfiguration(adminURL, auth string) (*Configuration, error) {
	// ToDo:
	// 1. Make HTTP requests to Kong Admin API endpoints
	// 2. Parse JSON responses into the models defined above
	// 3. Handle authentication (Basic Auth or token)
	// 4. Handle pagination and error responses

	return &Configuration{
		Services: []Service{},
		Routes:   []Route{},
		Plugins:  []Plugin{},
	}, nil
}
