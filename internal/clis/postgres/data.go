package postgres

// Connection contains required connection data 
// and optional pool configuration. Fields with 
// zero values and empty strings are ignored 
type Connection struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dbname   string `json:"dbname"`
	SslMode  string `json:"sslMode"`

	MaxConnections    int    `json:"pool_max_conns"`
	MinConnections    int    `json:"pool_min_conns"`
	MaxConnLifetime   string `json:"pool_max_conn_lifetime"`
	MaxConnIdleTime   string `json:"pool_max_conn_idle_time"`
	HeathCheckPeriod  string `json:"pool_health_check_period"`
	MaxConnLifeJitter string `json:"pool_max_conn_lifetime_jitter"`
}
