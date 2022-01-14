package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/3vilM33pl3/hexclient/internal/pkg/hexcloud"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
	"io/ioutil"
	"runtime"
	"time"
)

type Client struct {
	grpcClient hexcloud.HexagonServiceClient
	secure bool
}

func NewClient(serverAddr string, secure bool) (c *Client, err error) {
	c = &Client{}
	c.secure = secure
	err = c.Connect(serverAddr)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}
	return
}

func (c *Client) Status() (message string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status, err := c.grpcClient.GetStatus(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return status.Msg, nil
}

func (c *Client) Connect(serverAddr string) (err error) {
	var opts []grpc.DialOption
	var config *tls.Config
	opts = append(opts, grpc.WithAuthority(serverAddr))

	if c.secure {

		if runtime.GOOS == "windows" {
			rootCert, _ := ioutil.ReadFile("ca.cert")
			cp := x509.NewCertPool()
			if !cp.AppendCertsFromPEM([]byte(rootCert)) {
				return errors.New("credentials: failed to append certificates")
			}
			config = &tls.Config{
				InsecureSkipVerify: false,
				RootCAs:            cp,
			}
		} else {
			// if darwin, freebsd or linux
			systemRoots, err := x509.SystemCertPool()
			if err != nil {
				return errors.New("Error getting system certificate pool #{err}")
			}
			config = &tls.Config{
				RootCAs: systemRoots,
			}
		}

		cred := credentials.NewTLS(config)

		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return fmt.Errorf("Unable to connect: %s", err)
	}
	c.grpcClient = hexcloud.NewHexagonServiceClient(conn)

	return nil
}

var rootCert string =
"-----BEGIN CERTIFICATE-----\n"+
"MIIDujCCAqKgAwIBAgILBAAAAAABD4Ym5g0wDQYJKoZIhvcNAQEFBQAwTDEgMB4G\n"+
"A1UECxMXR2xvYmFsU2lnbiBSb290IENBIC0gUjIxEzARBgNVBAoTCkdsb2JhbFNp\n"+
"Z24xEzARBgNVBAMTCkdsb2JhbFNpZ24wHhcNMDYxMjE1MDgwMDAwWhcNMjExMjE1\n"+
"MDgwMDAwWjBMMSAwHgYDVQQLExdHbG9iYWxTaWduIFJvb3QgQ0EgLSBSMjETMBEG\n"+
"A1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbjCCASIwDQYJKoZI\n"+
"hvcNAQEBBQADggEPADCCAQoCggEBAKbPJA6+Lm8omUVCxKs+IVSbC9N/hHD6ErPL\n"+
"v4dfxn+G07IwXNb9rfF73OX4YJYJkhD10FPe+3t+c4isUoh7SqbKSaZeqKeMWhG8\n"+
"eoLrvozps6yWJQeXSpkqBy+0Hne/ig+1AnwblrjFuTosvNYSuetZfeLQBoZfXklq\n"+
"tTleiDTsvHgMCJiEbKjNS7SgfQx5TfC4LcshytVsW33hoCmEofnTlEnLJGKRILzd\n"+
"C9XZzPnqJworc5HGnRusyMvo4KD0L5CLTfuwNhv2GXqF4G3yYROIXJ/gkwpRl4pa\n"+
"zq+r1feqCapgvdzZX99yqWATXgAByUr6P6TqBwMhAo6CygPCm48CAwEAAaOBnDCB\n"+
"mTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUm+IH\n"+
"V2ccHsBqBt5ZtJot39wZhi4wNgYDVR0fBC8wLTAroCmgJ4YlaHR0cDovL2NybC5n\n"+
"bG9iYWxzaWduLm5ldC9yb290LXIyLmNybDAfBgNVHSMEGDAWgBSb4gdXZxwewGoG\n"+
"3lm0mi3f3BmGLjANBgkqhkiG9w0BAQUFAAOCAQEAmYFThxxol4aR7OBKuEQLq4Gs\n"+
"J0/WwbgcQ3izDJr86iw8bmEbTUsp9Z8FHSbBuOmDAGJFtqkIk7mpM0sYmsL4h4hO\n"+
"291xNBrBVNpGP+DTKqttVCL1OmLNIG+6KYnX3ZHu01yiPqFbQfXf5WRDLenVOavS\n"+
"ot+3i9DAgBkcRcAtjOj4LaR0VknFBbVPFd5uRHg5h6h+u/N5GJG79G+dwfCMNYxd\n"+
"AfvDbbnvRG15RjF+Cv6pgsH/76tuIMRQyV+dTZsXjAzlAcmgQWpzU/qlULRuJQ/7\n"+
"TBj0/VLZjmmx6BEP3ojY+x1J96relc8geMJgEtslQIxq/H5COEBkEveegeGTLg==\n"+
"-----END CERTIFICATE-----\n"