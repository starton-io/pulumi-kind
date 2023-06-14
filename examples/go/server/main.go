package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/starton-io/pulumi-kind/sdk/go/kind"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		kind, err := kind.NewCluster(ctx, "test", &kind.ClusterArgs{
			KindConfig: kind.ClusterKindConfigArgs{
				ApiVersion: pulumi.String("kind.x-k8s.io/v1alpha4"),
				Kind:       pulumi.String("Cluster"),
				Nodes: &kind.ClusterKindConfigNodeArray{
					kind.ClusterKindConfigNodeArgs{
						Role: pulumi.String("control-plane"),
						KubeadmConfigPatches: pulumi.StringArray{
							pulumi.String(`kind: ClusterConfiguration
apiServer:
  extraArgs:
    "service-node-port-range": "80-32100"`),
						},
					},
					kind.ClusterKindConfigNodeArgs{
						Role: pulumi.String("worker"),
					},
					kind.ClusterKindConfigNodeArgs{
						Role: pulumi.String("worker"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("status", kind.Completed)
		return nil
	})
}
