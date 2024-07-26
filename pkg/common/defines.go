package common

type MediaType int32

const (
	MediaTypeTrash MediaType = iota
	MediaTypeMovie
	MediaTypeTv
)

var (
	defaultOpts = map[string]string{
		"include_adult": "true",
	}
)

func GetGetDefaultTmdbSearchOpts() map[string]string {
	opts := make(map[string]string, len(defaultOpts))
	for k, v := range defaultOpts {
		opts[k] = v
	}
	return opts
}
