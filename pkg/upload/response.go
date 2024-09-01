package upload

type UploadResponse struct {
	PublicID   string `json:"public_id"`
	PrivateID  string `json:"private_id"`
	PublicUrl  string `json:"public_url"`
	PrivateUrl string `json:"private_url"`
}
