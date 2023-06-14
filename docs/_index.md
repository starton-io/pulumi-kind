---
title: Starton
meta_desc: Provides an overview of the Kind Provider for Pulumi.
layout: overview
---

The [Starton](https://starton.com/) provider for Pulumi can be used to provision any of the cloud resources available in Pulumi.
The Kind provider can be use without any configuration.

## Example

{{< chooser language "typescript,go" >}}
{{% choosable language typescript %}}

```typescript
import * as kindcluster from "@starton/kind";
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
	"fmt"
	scaleway "github.com/lbrlabs/pulumi-scaleway/sdk/go/scaleway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		publicIp, err := scaleway.NewInstanceIp(ctx, "example", &scaleway.InstanceIpArgs{})
		if err != nil {
			return fmt.Errorf("error creating public IP: %v", err)
		}

		server, err := scaleway.NewInstanceServer(ctx, "example", &scaleway.InstanceServerArgs{
			Image: pulumi.String("ubuntu_focal"),
			IpId:  publicIp.ID(),
			Type:  pulumi.String("DEV1-S"),
			Tags: pulumi.StringArray{
				pulumi.String("go"),
			},
		})
		if err != nil {
			return fmt.Errorf("error creating instance server: %v", err)
		}

		ctx.Export("server", server.Name)

		return nil
	})
}
```

{{% /choosable %}}

{{< /chooser >}}