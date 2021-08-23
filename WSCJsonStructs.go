package wetsponge

type RECVGameEvent struct {
	Body struct {
		EventName    string `json:"eventName"`
		Measurements struct {
			Count     int `json:"Count"`
			RecordCnt int `json:"RecordCnt"`
			SeqMax    int `json:"SeqMax"`
			SeqMin    int `json:"SeqMin"`
		} `json:"measurements"`
		Properties struct {
			AccountType                    int    `json:"AccountType"`
			ActiveSessionID                string `json:"ActiveSessionID"`
			AppSessionID                   string `json:"AppSessionID"`
			Biome                          int    `json:"Biome"`
			Build                          string `json:"Build"`
			BuildPlat                      int    `json:"BuildPlat"`
			Cheevos                        bool   `json:"Cheevos"`
			ClientID                       string `json:"ClientId"`
			CurrentInput                   int    `json:"CurrentInput"`
			DeviceSessionID                string `json:"DeviceSessionId"`
			Dim                            int    `json:"Dim"`
			GlobalMultiplayerCorrelationID string `json:"GlobalMultiplayerCorrelationId"`
			Message                        string `json:"Message"`
			MessageType                    string `json:"MessageType"`
			Mode                           int    `json:"Mode"`
			MultiplayerCorrelationID       string `json:"MultiplayerCorrelationId"`
			NetworkType                    int    `json:"NetworkType"`
			Plat                           string `json:"Plat"`
			PlayerGameMode                 int    `json:"PlayerGameMode"`
			SchemaCommitHash               string `json:"SchemaCommitHash"`
			Sender                         string `json:"Sender"`
			ServerID                       string `json:"ServerId"`
			Treatments                     string `json:"Treatments"`
			UserID                         string `json:"UserId"`
			WorldFeature                   int    `json:"WorldFeature"`
			WorldSessionID                 string `json:"WorldSessionId"`
			EditionType                    string `json:"editionType"`
			IsTrial                        int    `json:"isTrial"`
			Locale                         string `json:"locale"`
			VrMode                         bool   `json:"vrMode"`
		} `json:"properties"`
	} `json:"body"`
	Header struct {
		MessagePurpose string `json:"messagePurpose"`
		RequestID      string `json:"requestId"`
		Version        int    `json:"version"`
	} `json:"header"`
}
type RECVCmdRequest struct {
	Body struct {
		StatusCode    int    `json:"statusCode"`
		StatusMessage string `json:"statusMessage"`
	} `json:"body"`
	Header struct {
		MessagePurpose string `json:"messagePurpose"`
		RequestID      string `json:"requestId"`
	} `json:"header"`
}
type RECVlist struct {
	Body struct {
		CurrentPlayerCount int    `json:"currentPlayerCount"`
		MaxPlayerCount     int    `json:"maxPlayerCount"`
		Players            string `json:"players"`
		StatusCode         int    `json:"statusCode"`
		StatusMessage      string `json:"statusMessage"`
	} `json:"body"`
	Header struct {
		MessagePurpose string `json:"messagePurpose"`
		RequestID      string `json:"requestId"`
		Version        int    `json:"version"`
	} `json:"header"`
}
type RECVsay struct {
	Body struct {
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	} `json:"body"`
	Header struct {
		MessagePurpose string `json:"messagePurpose"`
		RequestID      string `json:"requestId"`
		Version        int    `json:"version"`
	} `json:"header"`
}
type RECVtp struct {
	Body struct {
		Destination struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"destination"`
		StatusCode    int      `json:"statusCode"`
		StatusMessage string   `json:"statusMessage"`
		Victim        []string `json:"victim"`
	} `json:"body"`
	Header struct {
		MessagePurpose string `json:"messagePurpose"`
		RequestID      string `json:"requestId"`
		Version        int    `json:"version"`
	} `json:"header"`
}
