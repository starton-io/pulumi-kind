---
title: Kins
meta_desc: Provides an overview of the Kind Provider for Pulumi.
layout: overview
---

The [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/) provider for Pulumi can be used to provision any of the cloud resources available in Pulumi.
The Kind provider can be use without any configuration.

## Example

{{< chooser language "typescript,go" >}}
{{% choosable language typescript %}}

```typescript
import * as kindcluster from "@starton/pulumi-kind";
const instanceKube = new kindcluster.Cluster("test", {
	kindConfig: {
		apiVersion: "kind.x-k8s.io/v1alpha4",
		kind: "Cluster",
		networking: {
			apiServerAddress: "0.0.0.0",
			apiServerPort: 6443
		},
		nodes: [
			{
				role: "control-plane",
				kubeadmConfigPatches: [
					`
kind: ClusterConfiguration
apiServer:
  extraArgs:
    "service-node-port-range": "80-32100"
`,
				],
				extraPortMappings: [
					{
						containerPort: 32081,
						hostPort: 32081,
						listenAddress: "0.0.0.0"
					},
					{
						containerPort: 32082,
						hostPort: 32082,
						listenAddress: "0.0.0.0"
					},
					{
						containerPort: 80,
						hostPort: 80,
						listenAddress: "0.0.0.0"
					},
				]

			},
			{
				role: "worker",
			},
			{
				role: "worker",
			}
		],
	},
	name: "dev",
	nodeImage: "kindest/node:v1.23.17"
}
);
```

{{% /choosable %}}
{{% choosable language go %}}

```go

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

```

{{% /choosable %}}

{{< /chooser >}}