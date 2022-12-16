package main

import (
	"context"

	apiv1 "github.com/charconstpointer/ihateannotations/proto/gen/go/api/v1"
)

func main() {
	c := apiv1.NewApiServiceClient(nil)
	res, err := c.Foo(context.TODO(), nil, nil)
	if
}
