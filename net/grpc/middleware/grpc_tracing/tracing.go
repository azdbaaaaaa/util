package grpc_tracing

import "google.golang.org/grpc/metadata"

type TextMapWriter struct {
	metadata.MD
}

func (t TextMapWriter) Set(key, val string) {
	t.MD[key] = append(t.MD[key], val)
}

type TextMapReader struct {
	metadata.MD
}

func (t TextMapReader) ForeachKey(handler func(key, val string) error) error {
	for key, val := range t.MD {
		for _, v := range val {
			if err := handler(key, v); err != nil {
				return err
			}
		}
	}
	return nil
}
