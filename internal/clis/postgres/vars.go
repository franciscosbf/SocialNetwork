package postgres

// varInfo contains variable info
// such as its dsn key representation,
// name to fetch its info and if it's
// required or not
type varInfo struct {
	dsnName  string
	varName  string
	required bool
}

// confVars contains all
// client config variables
var confVars []*varInfo

// registerVar registers a given variable that may be required or not
func registerVar(dsnName, varName string, required bool) {
	info := &varInfo{dsnName, varName, required}
	confVars = append(confVars, info)
}

// requiredVar registers required variable
func requiredVar(dsnName, varName string) {
	registerVar(dsnName, varName, true)
}

// requiredVar registers optional variable
func optionalVar(dsnName, varName string) {
	registerVar(dsnName, varName, false)
}

// init registers all variables
func init() {
	// Connection elements
	requiredVar("user", "POSTGRES_USER")
	requiredVar("password", "POSTGRES_PASSWORD")
	requiredVar("host", "POSTGRES_HOST")
	requiredVar("port", "POSTGRES_PORT")
	requiredVar("dbname", "POSTGRES_DBNAME")
	optionalVar("sslmode", "POSTGRES_SSL_MODE")

	// Connections pool configuration
	optionalVar("pool_max_conns", "POSTGRES_POOL_MAX_CONNS")
	optionalVar("pool_min_conns", "POSTGRES_POOL_MIN_CONNS")
	optionalVar("pool_max_conn_lifetime", "POSTGRES_POOL_MAX_CONN_LIFETIME")
	optionalVar("pool_max_conn_idle_time", "POSTGRES_POOL_MAX_CONN_IDLE_TIME")
	optionalVar("pool_health_check_period", "POSTGRES_POOL_HEALTH_CHECK_PERIOD")
	optionalVar("pool_max_conn_lifetime_jitter", "POSTGRES_POOL_MAX_CONN_LIFETIME_JITTER")
}
