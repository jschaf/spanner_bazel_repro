package main

import (
	"context"

	"cloud.google.com/go/spanner"
)

func main() {
	ctx := context.Background()
	_, err := spanner.NewClient(ctx, "projects/your-project-id/instances/your-instance-id")
	if err != nil {
		panic(err)
	}
}
