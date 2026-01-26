module markitos-it-app-website

go 1.24.0

require (
	github.com/yuin/goldmark v1.7.16
	google.golang.org/grpc v1.78.0
	markitos-it-svc-documents v0.0.0
)

replace markitos-it-svc-documents => ../markitos-it-svc-documents

require (
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
