package utils

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// Generate a self-signed X.509 certificate for a TLS server. Outputs to
// 'cert.pem' and 'key.pem' and will overwrite existing files.

import (
  "crypto/ecdsa"
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "crypto/x509/pkix"
  "encoding/pem"
  "flag"
  "fmt"
  "math/big"
  "net"
  "os"
  "strings"
  "time"
  "crypto/elliptic"

  "github.com/zeebo/errs"
)

var (
  // host       = flag.String("host", "", "Comma-separated hostnames and IPs to generate a certificate for")
  validFrom = flag.String("start-date", "", "Creation date formatted as Jan 1 15:04:05 2011")
  validFor  = flag.Duration("duration", 365*24*time.Hour, "Duration that certificate is valid for")
  isCA      = flag.Bool("ca", false, "whether this cert should be its own Certificate Authority")
  // rsaBits    = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if --ecdsa-curve is set")
  // ecdsaCurve = flag.String("ecdsa-curve", "", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
)

func publicKey(priv interface{}) interface{} {
  switch k := priv.(type) {
  case *rsa.PrivateKey:
    return &k.PublicKey
  case *ecdsa.PrivateKey:
    return &k.PublicKey
  default:
    return nil
  }
}

func pemBlockForKey(priv interface{}) *pem.Block {
  switch k := priv.(type) {
  case *rsa.PrivateKey:
    return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
  case *ecdsa.PrivateKey:
    b, err := x509.MarshalECPrivateKey(k)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
      os.Exit(2)
    }
    return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
  default:
    return nil
  }
}

func (t *TLSFileOptions) generate() (_ error) {
  if t.Hosts == "" {
    return ErrGenerate.Wrap(ErrBadHost.New("no host provided"))
  }

  if err := t.EnsureAbsPaths(); err != nil {
    return err
  }

  priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader);
  if err != nil {
    return ErrGenerate.Wrap(errs.New("failed to generate private key: %s", err.Error()))
  }

  var notBefore time.Time

  // TODO: `validFrom`
  if len(*validFrom) == 0 {
    notBefore = time.Now()
  } else {
    notBefore, err = time.Parse("Jan 2 15:04:05 2006", *validFrom)
    if err != nil {
      return ErrGenerate.Wrap(errs.New("Failed to parse creation date: %s\n", err.Error()))
    }
  }

  // TODO: `validFor`
  notAfter := notBefore.Add(*validFor)

  serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
  serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
  if err != nil {
    return ErrGenerate.Wrap(errs.New("failed to generate serial number: %s", err.Error()))
  }

  template := x509.Certificate{
    SerialNumber: serialNumber,
    Subject: pkix.Name{
      Organization: []string{"Acme Co"},
    },
    NotBefore: notBefore,
    NotAfter:  notAfter,

    KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
    ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
    BasicConstraintsValid: true,
  }

  hosts := strings.Split(t.Hosts, ",")
  for _, h := range hosts {
    if ip := net.ParseIP(h); ip != nil {
      template.IPAddresses = append(template.IPAddresses, ip)
    } else {
      template.DNSNames = append(template.DNSNames, h)
    }
  }

  // TODO: `isCA`
  if *isCA {
    template.IsCA = true
    template.KeyUsage |= x509.KeyUsageCertSign
  }

  derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
  if err != nil {
    return ErrGenerate.Wrap(err)
  }

  certOut, err := os.Create(t.CertAbsPath)
  if err != nil {
    return ErrGenerate.Wrap(errs.New("failed to open cert.pem for writing: %s", err.Error()))
  }

  pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
  certOut.Close()

  keyOut, err := os.OpenFile(t.KeyAbsPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
  if err != nil {
    return ErrGenerate.Wrap(errs.New("failed to open key.pem for writing:", err))
    return
  }

  pem.Encode(keyOut, pemBlockForKey(priv))
  keyOut.Close()

  return nil
}
