package mock

type Retriver1 struct {
	Contents string
}

func (r *Retriver1) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r Retriver1) Get(url string) string {
	return r.Contents
}


