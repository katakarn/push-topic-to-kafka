package version

type Response struct {
	Version   string `json:"version"`
	BuildTime string `json:"buildTime"`
	Build     string `json:"build"`
}
