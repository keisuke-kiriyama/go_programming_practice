package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,                                                 // 証明書のシリアル番号。この目的ではランダムで生成した大きい整数で十分
		Subject:      subject,                                                      // 識別名
		NotBefore:    time.Now(),                                                   // 証明書の有効期間
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),                         // 証明書の有効期間
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // サーバー認証に使用されることを示す
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // サーバー認証に使用されることを示す
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},                           // 127.0.0.1だけ効力をもつ
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048) // RSA秘密鍵の生成。この中に公開鍵が入っている

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) // 公開鍵を使って証明書生成
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}) // 証明書データをふごうかしてcert.pemというファイルにする
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}) // 秘密鍵をふごうかして、key.pemというファイルにする
}
