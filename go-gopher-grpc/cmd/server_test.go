package cmd_test

import (
	"context"
	"testing"

	cmd "github.com/scraly/learning-go-by-examples/go-gopher-grpc/cmd"
	pb "github.com/scraly/learning-go-by-examples/go-gopher-grpc/pkg/gopher"

	. "github.com/onsi/gomega"
)

func TestGetGopher(t *testing.T) {
	s := cmd.Server{}

	testCases := []struct {
		name        string
		req         *pb.GopherRequest
		message     string
		expectedErr bool
	}{
		{
			name:        "req ok",
			req:         &pb.GopherRequest{Name: "yoda-gopher"},
			message:     "https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png\n",
			expectedErr: false,
		},
		{
			name:        "req with empty name",
			req:         &pb.GopherRequest{},
			expectedErr: true,
		},
		{
			name:        "nil request",
			req:         nil,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctx := context.Background()

			// call
			response, err := s.GetGopher(ctx, testCase.req)

			t.Log("Got : ", response)

			// assert results expectations
			if testCase.expectedErr {
				g.Expect(response).ToNot(BeNil(), "Result should be nil")
				g.Expect(err).ToNot(BeNil(), "Result should be nil")
			} else {
				g.Expect(response.Message).To(Equal(testCase.message))
			}
		})
	}
}

// func HelloTest(t *testing.T) {
//     s := server{}

//     // set up test cases
//     tests := []struct{
//         name string
//         want string
//     } {
//         {
//             name: "world",
//             want: "Hello world",
//         },
//         {
//             name: "123",
//             want: "Hello 123",
//         },
//     }

//     for _, tt := range tests {
//         req := &pb.HelloRequest{Name: tt.name}
//         resp, err := s.SayHello(context.Background(), req)
//         if err != nil {
//             t.Errorf("HelloTest(%v) got unexpected error")
//         }
//         if resp.Message != tt.want {
//             t.Errorf("HelloText(%v)=%v, wanted %v", tt.name, resp.Message, tt.want)
//         }
//     }
// }
