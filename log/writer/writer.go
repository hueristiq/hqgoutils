package writer

import "github.com/hueristiq/hqgoutils/log/levels"

type Writer interface {
	Write(data []byte, level levels.LevelInt)
}
