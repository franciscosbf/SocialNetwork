/*
Copyright 2023 Francisco Simões Braço-Forte

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
