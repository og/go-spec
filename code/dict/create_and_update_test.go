package dict

import "testing"

func ExampleCreateAndUpdate() {
	Create(QueryCreate{
		QueryFrom: QueryFrom{
			Status: QueryFrom{}.Dict().Status.Normal,
			Type: QueryFrom{}.Dict().Type.Exigency,
			Title: "a",
			Content:"some",
		},
		ID: "1",
	})
	Update(QueryUpdate{
		QueryFrom: QueryFrom{
			Status: QueryFrom{}.Dict().Status.Normal,
			Type: QueryFrom{}.Dict().Type.Exigency,
			Title: "a",
			Content:"some",
		},
	})
}

func TestCreateAndUpdate(t *testing.T) {
	ExampleCreateAndUpdate()
}
