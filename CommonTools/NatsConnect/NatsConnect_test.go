package NatsConnect

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/nats-io/go-nats"

	"github.com/ximidar/Flotilla/BuildResources/Test/CommonTestTools"

	"github.com/nats-io/gnatsd/server"
	"github.com/nats-io/gnatsd/test"
)

func TestRegularConnection(t *testing.T) {

	// Start a default server
	srv := test.RunDefaultServer()
	defer srv.Shutdown()
	addr := srv.Addr()

	// try to connect
	conn, err := DefaultConn(addr.String(), "regular")
	CommonTestTools.CheckErr(t, "Couldn't connect to regular server", err)

	fmt.Printf("Connected? %v", conn.IsConnected())
}

func TestTLSConnection(t *testing.T) {
	fmt.Println("starting tls")
	// Make temp certs
	MakeCert()
	tlsConf := makeTLSConf()

	opts := &server.Options{
		TLS:        true,
		TLSConfig:  tlsConf,
		TLSTimeout: 2,
		Trace:      true,
		Debug:      true,
	}

	srv := test.RunServer(opts)
	defer srv.Shutdown()

	addr := srv.Addr()

	// Test connection
	conn, err := DefaultConn(addr.String(), "tls")
	CommonTestTools.CheckErr(t, "Couldn't connect to regular server", err)

	fmt.Printf("Connected? %v", conn.IsConnected())

}

func makeTLSConf() *tls.Config {

	certFile := "/tmp/cert/flotilla.pem"
	keyFile := "/tmp/cert/flotilla.key"
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{
		ServerName:   nats.DefaultURL,
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}
	return config
}

func MakeCert() string {
	certFolder := "/tmp/cert/"
	confFile := path.Clean(certFolder + "/TLS.cnf")
	// remake the certs folder
	if _, err := os.Stat(certFolder); os.IsNotExist(err) {
		os.MkdirAll(certFolder, os.ModePerm)
	} else {
		os.RemoveAll(certFolder)
		os.MkdirAll(certFolder, os.ModePerm)
	}

	// make the conf file
	f, err := os.Create(confFile)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(certFile)
	if err != nil {
		f.Close()
		panic(err)
	}
	f.Close()

	// Make a monitor for a command we want to run
	command := "openssl"
	args := []string{"req",
		"-new",
		"-x509",
		"-newkey",
		"rsa:2048",
		"-config",
		confFile,
		"-keyout",
		"flotilla.key",
		"-out",
		"flotilla.pem",
		"-outform",
		"PEM"}

	cmd := exec.Command(command, args...)
	cmd.Dir = certFolder

	err = cmd.Run()
	if err != nil {
		out, _ := cmd.Output()
		fmt.Println("Could not run openssl command", string(out), err)
		panic(err)
	}
	return certFolder
}

var certFile = "# Reference: https://www.switch.ch/pki/manage/request/csr-openssl/" + "\n" +
	"# Reference#2: https://help.ubuntu.com/community/OpenSSL" + "\n" +
	"# OpenSSL configuration file for creating a CSR for a server certificate" + "\n" +
	"# Adapt at least the FQDN and ORGNAME lines, and then run " + "\n" +
	"# openssl req -new -x509 -newkey rsa:2048 -config config.cnf -keyout flotilla.key -out flotilla.pem -outform PEM" + "\n" +
	"# on the command line." + "\n" +
	"" + "\n" +
	"# the fully qualified server (or service) name" + "\n" +
	"FQDN = www.flotilla.com" + "\n" +
	"" + "\n" +
	"# the name of your organization" + "\n" +
	"# (see also https://www.switch.ch/pki/participants/)" + "\n" +
	"ORGNAME = Flotilla" + "\n" +
	"" + "\n" +
	"# subjectAltName entries: to add DNS aliases to the CSR, delete" + "\n" +
	"# the '#' character in the ALTNAMES line, and change the subsequent" + "\n" +
	"# 'DNS:' entries accordingly. Please note: all DNS names must" + "\n" +
	"# resolve to the same IP address as the FQDN." + "\n" +
	"ALTNAMES = DNS:$FQDN" + "\n" +
	"" + "\n" +
	"# --- no modifications required below ---" + "\n" +
	"[ req ]" + "\n" +
	"default_bits = 2048" + "\n" +
	"default_md = sha256" + "\n" +
	"prompt = no" + "\n" +
	"encrypt_key = no" + "\n" +
	"distinguished_name = dn" + "\n" +
	"req_extensions = req_ext" + "\n" +
	"" + "\n" +
	"[ dn ]" + "\n" +
	"C = US" + "\n" +
	"O = $ORGNAME" + "\n" +
	"CN = $FQDN" + "\n" +
	"" + "\n" +
	"[ req_ext ]" + "\n" +
	"subjectAltName = $ALTNAMES" + "\n"
