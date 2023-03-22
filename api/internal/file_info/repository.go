package file_info

import "context"

type Repository interface {
	GetFileInfo(ctx context.Context)
}
