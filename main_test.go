package main

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_response(t *testing.T) {
	type args struct {
		body       string
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want events.APIGatewayProxyResponse
	}{
		{
			name: "hello world json",
			args: args{
				body:       "{'message': 'Hello World!'}",
				statusCode: http.StatusOK,
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "{'message': 'Hello World!'}",
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := response(tt.args.body, tt.args.statusCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helloHandler(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{
			name: "basic",
			args: args{},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "{\"message\":\"Hello, World!\"}",
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := helloHandler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("helloHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helloHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
