package skill

type simpleError string

func (e simpleError) Error() string {
	return string(e)
}

var ErrorNoResponse = simpleError("no response")
