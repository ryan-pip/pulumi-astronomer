// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Cluster data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const example = astronomer.getCluster({
 *     id: "clozc036j01to01jrlgvueo8t",
 * });
 * ```
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("astronomer:index/getCluster:getCluster", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * Cluster identifier
     */
    id: string;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    /**
     * Cluster cloud provider
     */
    readonly cloudProvider: string;
    /**
     * Cluster creation timestamp
     */
    readonly createdAt: string;
    /**
     * Cluster database instance type
     */
    readonly dbInstanceType: string;
    /**
     * Cluster identifier
     */
    readonly id: string;
    /**
     * Whether the cluster is limited
     */
    readonly isLimited: boolean;
    /**
     * Cluster metadata
     */
    readonly metadata: outputs.GetClusterMetadata;
    /**
     * Cluster name
     */
    readonly name: string;
    /**
     * Cluster node pools
     */
    readonly nodePools: outputs.GetClusterNodePool[];
    /**
     * Cluster pod subnet range
     */
    readonly podSubnetRange: string;
    /**
     * Cluster provider account
     */
    readonly providerAccount: string;
    /**
     * Cluster region
     */
    readonly region: string;
    /**
     * Cluster service peering range
     */
    readonly servicePeeringRange: string;
    /**
     * Cluster service subnet range
     */
    readonly serviceSubnetRange: string;
    /**
     * Cluster status
     */
    readonly status: string;
    /**
     * Cluster tags
     */
    readonly tags: outputs.GetClusterTag[];
    /**
     * Cluster tenant ID
     */
    readonly tenantId: string;
    /**
     * Cluster type
     */
    readonly type: string;
    /**
     * Cluster last updated timestamp
     */
    readonly updatedAt: string;
    /**
     * Cluster VPC subnet range
     */
    readonly vpcSubnetRange: string;
    /**
     * Cluster workspace IDs
     */
    readonly workspaceIds: string[];
}
/**
 * Cluster data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const example = astronomer.getCluster({
 *     id: "clozc036j01to01jrlgvueo8t",
 * });
 * ```
 */
export function getClusterOutput(args: GetClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterResult> {
    return pulumi.output(args).apply((a: any) => getCluster(a, opts))
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterOutputArgs {
    /**
     * Cluster identifier
     */
    id: pulumi.Input<string>;
}
