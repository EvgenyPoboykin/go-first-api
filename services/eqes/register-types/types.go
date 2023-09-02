package registerTypes

type ItemSignature struct {
	SerialNumber string `json:"serialNumber"`
	Thumbprint   string `json:"thumbprint"`
}

type PostArgs struct {
	Signature string `json:"signature"`
}
