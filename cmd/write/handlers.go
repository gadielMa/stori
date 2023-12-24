package write

import "context"

type AccountService interface {
	ProcessSummary(ctx context.Context) error
}
