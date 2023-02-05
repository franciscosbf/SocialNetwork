package postgres

import (
	"fmt"
	"strings"
)

// DsnConn represents dsn connection values
type DsnConn struct {
	values []string
}

// addString adds a string value to dsn
func (d *DsnConn) addString(key string, value string) {
	if value == "" {
		return
	}

	pair := fmt.Sprintf("%v=%v", key, value)
	d.values = append(d.values, pair)
}

// addInt adds an int value to dsn
func (d *DsnConn) addInt(key string, value int) {
	if value == 0 {
		return
	}

	pair := fmt.Sprintf("%v=%v", key, value)
	d.values = append(d.values, pair)
}

// unify builds the dsn string
func (d *DsnConn) unify() string {
	return strings.Join(d.values, " ")
}

// BuildDsn returns a valid Postgres connection dsn
func BuildDsn(connData *Connection) string {
	raw := &DsnConn{}

	// Connection elements
	raw.addString("user", connData.User)
	raw.addString("password", connData.Password)
	raw.addString("host", connData.Host)
	raw.addInt("port", connData.Port)
	raw.addString("dbname", connData.Dbname)
	raw.addString("sslmode", connData.SslMode)

	// Optional config parameters
	raw.addInt("pool_max_conns", connData.MaxConnections)
	raw.addInt("pool_min_conns", connData.MinConnections)
	raw.addString("pool_max_conn_lifetime", connData.MaxConnLifetime)
	raw.addString("pool_max_conn_idle_time", connData.MaxConnIdleTime)
	raw.addString("pool_health_check_period", connData.HeathCheckPeriod)
	raw.addString("pool_max_conn_lifetime_jitter", connData.MaxConnLifeJitter)

	return raw.unify()
}
