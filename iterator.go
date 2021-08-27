package obs

type objectPageStatus struct {
	prefix    string
	maxKeys   int
	delimiter string
	marker    string
}

func (i *objectPageStatus) ContinuationToken() string {
	return i.marker
}
