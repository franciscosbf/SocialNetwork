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

package secure

import (
	"crypto/tls"
	"crypto/x509"
)

// GenClientTls creates a new client TLS config given the server name, client
// certificate and key, and at least one CA server certificate. Returns an error
// if it wasn't possible to create a client's public/private key pair. Expects
// all certs encoded in PEM format
func GenClientTls(serverName, clientCert, clientKey, caCert string) (*tls.Config, error) {
	// Creates client authentication with its
	// certificate and key pair, encoded using PEM
	cliAuth, err := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
	if err != nil {
		return nil, err
	}

	// Stores CA certificates in a new pool
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM([]byte(caCert))

	return &tls.Config{
		ServerName:   serverName,
		RootCAs:      caPool,
		Certificates: []tls.Certificate{cliAuth},
	}, nil
}
