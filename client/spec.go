package client

// Falcon TODO: Make memberCid and clientCloud parameter as optional

type Falcon struct {
	Name         string `json:"Name,omitempty"`
	ClientId     string `json:"ClientId,omitempty" `
	ClientSecret string `json:"ClientSecret,omitempty"`
	ClientCloud  string `json:"ClientCloud,omitempty" `
	MemberCid    string `json:"MemberCid,omitempty" `
}
type Spec struct {
	Concurrency int      `json:"Concurrency,omitempty"`
	FALCON      []Falcon `json:"Falcon,omitempty"`
}

func (Spec) Validate() error {
	return nil
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 1000
	}
}
