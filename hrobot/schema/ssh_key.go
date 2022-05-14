package schema

type SSHKeys []struct {
	Key struct {
		Name        string `json:"name"`
		Fingerprint string `json:"fingerprint"`
		Type        string `json:"type"`
		Size        int    `json:"size"`
		Data        string `json:"data"`
	} `json:"key"`
}

type Key struct {
	Key struct {
		Name        string `json:"name"`
		Fingerprint string `json:"fingerprint"`
		Type        string `json:"type"`
		Size        int    `json:"size"`
		Data        string `json:"data"`
	} `json:"key"`
}

// defines the schema for a single key
type SSHKeyGetResponse struct {
	SSHKey SSHKeys `json:"ssh_key"`
}

// SSHKeyListResponse defines the schema of the response
// when listing SSH keys.
//type SSHKeyListResponse struct {
//SSHKeys []SSHKey
//}
