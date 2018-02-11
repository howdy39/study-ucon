package app

import (
	"github.com/favclip/ucon"
	"google.golang.org/appengine"
)

// UseAppengineContext GAE向けにBubbleにcontextを設定する
func UseAppengineContext(b *ucon.Bubble) error {
	if b.Context == nil {
		b.Context = appengine.NewContext(b.R)
	} else {
		b.Context = appengine.WithContext(b.Context, b.R)
	}
	return b.Next()
}
