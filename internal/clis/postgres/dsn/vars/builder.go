package vars

// PostgresVarInfo contains variable info
// such as its dsn key representation,
// name to fetch its info and if it's
// Required or not
type PostgresVarInfo struct {
	DsnName  string
	VarName  string
	Required bool
}

// confVars contains all
// client config variables
var confVars []*PostgresVarInfo

// registerVar registers a given variable that may be Required or not
func registerVar(dsnName, varName string, required bool) {
	info := &PostgresVarInfo{dsnName, varName, required}
	confVars = append(confVars, info)
}

// requiredVar registers Required variable
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

// infoCopy returns a contents copy of a given var
func infoCopy(from *PostgresVarInfo) *PostgresVarInfo {
	return &PostgresVarInfo{
		DsnName:  from.DsnName,
		VarName:  from.VarName,
		Required: from.Required,
	}
}

// ForEachPostgresVar Iterates over all vars with a reader func. If reader
// returns some error, then ForEachPostgresVar will immediately return it
func ForEachPostgresVar(reader func(info *PostgresVarInfo) error) error {
	for _, vi := range confVars {
		cpy := infoCopy(vi)

		if err := reader(cpy); err != nil {
			return err
		}
	}

	return nil
}