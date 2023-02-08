package token

import "time"

var (
	// Cookie parameters
	CookieName = "authToken"

	// Token parameter
	Duration = time.Minute * 60                   // duration of paseto token: 60 min
	IV       = "cftheB6xhdMeRgQ7Hrafi9uUWW9R5ExX" // initialization vector
)
