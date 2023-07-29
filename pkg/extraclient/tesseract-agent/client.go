package tesseract_agent

import (
	"context"
	"fmt"
	"github.com/upmio/config-wrapper/app/config"
	"github.com/upmio/config-wrapper/app/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func SyncConfig(agentHostType, host, port, namespace, configmapName string) (string, error) {
	// grpc.Dial负责和gRPC服务建立链接

	addr := ""
	switch agentHostType {
	case "domain":
		fullDomainName := fmt.Sprintf("%s.%s.svc.cluster.local", host, namespace)
		addr = net.JoinHostPort(fullDomainName, port)
	case "ip":
		addr = net.JoinHostPort(host, port)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := config.NewSyncConfigServiceClient(conn)

	req := config.SyncConfigRequest{
		Namespace:     namespace,
		ConfigmapName: configmapName,
	}

	resp, err := client.SyncConfig(context.Background(), &req)
	if err != nil {
		return resp.GetMessage(), err
	}

	return "", nil
}

func ServiceLifecycleManagement(agentHostType, host, namespace, port, actionType string) (string, error) {
	// grpc.Dial负责和gRPC服务建立链接

	addr := ""
	switch agentHostType {
	case "domain":
		fullDomainName := fmt.Sprintf("%s.%s.svc.cluster.local", host, namespace)
		addr = net.JoinHostPort(fullDomainName, port)
	case "ip":
		addr = net.JoinHostPort(host, port)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	req := service.ServiceRequest{}

	client := service.NewServiceLifecycleClient(conn)

	switch actionType {
	case "start":
		resp, err := client.StartService(context.Background(), &req)
		if err != nil {
			return resp.GetMessage(), err
		}
	case "stop":
		resp, err := client.StopService(context.Background(), &req)
		if err != nil {
			return resp.GetMessage(), err
		}
	default:
		return "", fmt.Errorf("[%s] server not support", actionType)
	}

	return fmt.Sprintf("[%s] server ok", actionType), nil
}