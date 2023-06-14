import * as pulumi from "@pulumi/pulumi";
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
