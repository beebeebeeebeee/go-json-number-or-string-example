package app

type JsonString string

func (j *JsonString) UnmarshalJSON(data []byte) error {
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	*j = JsonString(data)
	return nil
}

func (j *JsonString) String() string {
	return string(*j)
}
