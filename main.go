// Command criteria-adapter-starter-go is a minimal hello-world Criteria adapter.
// Edit greeterAdapter below to build your own; call adapterhost.Serve for local
// launch or adapterhost.ServeRemote for remote (phone-home) mode.
package main

import (
	"context"

	"github.com/brokenbots/criteria-go-adapter-sdk/adapterhost"
	v2 "github.com/brokenbots/criteria-adapter-proto/criteria/v2"
)

type greeterAdapter struct {
	adapterhost.UnimplementedPermissions
}

func (greeterAdapter) Info(context.Context, *v2.InfoRequest) (*v2.InfoResponse, error) {
	return &v2.InfoResponse{
		Name:         "my-adapter",
		Version:      "0.1.0",
		Description:  "A starter Criteria adapter",
		SourceUrl:    "https://github.com/your-org/your-adapter",
		Capabilities: []string{"execute"},
	}, nil
}

func (greeterAdapter) OpenSession(context.Context, *v2.OpenSessionRequest) (*v2.OpenSessionResponse, error) {
	return &v2.OpenSessionResponse{}, nil
}

func (greeterAdapter) Execute(_ context.Context, req *v2.ExecuteRequest, sender adapterhost.ExecuteEventSender) error {
	name := req.GetInput()["name"]
	if name == "" {
		name = "world"
	}
	ev, err := v2.NewExecuteResultEvent("greet", map[string]any{"greeting": "Hello, " + name + "!"})
	if err != nil {
		return err
	}
	return sender.Send(ev)
}

func (greeterAdapter) Log(context.Context, *v2.LogRequest, adapterhost.LogEventSender) error {
	return nil
}

func (greeterAdapter) CloseSession(context.Context, *v2.CloseSessionRequest) (*v2.CloseSessionResponse, error) {
	return &v2.CloseSessionResponse{}, nil
}

func main() {
	// Local launch (the host starts this process). For remote (phone-home)
	// mode, replace this with adapterhost.ServeRemote(&greeterAdapter{}, opts)
	// — see examples/remote/ and the SDK README.
	adapterhost.Serve(&greeterAdapter{})
}
