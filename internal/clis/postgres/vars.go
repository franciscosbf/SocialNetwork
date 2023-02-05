package postgres

// confVars contains pairs of
// dsn key pair and var name
var confVars = []*struct {
	dsnName string
	varName string
}{
	// Connection elements

	{"user", "POSTGRES_USER"},
	{"password", "POSTGRES_PASSWORD"},
	{"host", "POSTGRES_HOST"},
	{"port", "POSTGRES_PORT"},
	{"dbname", "POSTGRES_DBNAME"},
	{"sslmode", "POSTGRES_SSL_MODE"},

	// Connections pool configuration

	{"pool_max_conns", "POSTGRES_POOL_MAX_CONNS"},
	{"pool_min_conns", "POSTGRES_POOL_MIN_CONNS"},
	{"pool_max_conn_lifetime", "POSTGRES_POOL_MAX_CONN_LIFETIME"},
	{"pool_max_conn_idle_time", "POSTGRES_POOL_MAX_CONN_IDLE_TIME"},
	{"pool_health_check_period", "POSTGRES_POOL_HEALTH_CHECK_PERIOD"},
	{"pool_max_conn_lifetime_jitter", "POSTGRES_POOL_MAX_CONN_LIFETIME_JITTER"},
}
