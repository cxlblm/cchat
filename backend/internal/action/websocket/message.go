package websocket

type Base struct {
	Action string
}

type Login struct {
	Base
	Data LoginData
}

type LoginData struct {
	username string
	password string
}

type Profile struct {
	Base
	Profile ProfileData
}

type ProfileData struct {
	nick string
}
