// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Clusters data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleClusters = astronomer.getClusters({});
 * const exampleClustersFilterByNames = astronomer.getClusters({
 *     names: ["my cluster"],
 * });
 * const exampleClustersFilterByCloudProvider = astronomer.getClusters({
 *     cloudProvider: "AWS",
 * });
 * export const clusters = exampleClusters;
 * ```
 */
export function getClusters(args?: GetClustersArgs, opts?: pulumi.InvokeOptions): Promise<GetClustersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("astronomer:index/getClusters:getClusters", {
        "cloudProvider": args.cloudProvider,
        "names": args.names,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersArgs {
    cloudProvider?: string;
    names?: string[];
}

/**
 * A collection of values returned by getClusters.
 */
export interface GetClustersResult {
    readonly cloudProvider?: string;
    readonly clusters: outputs.GetClustersCluster[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly names?: string[];
}
/**
 * Clusters data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleClusters = astronomer.getClusters({});
 * const exampleClustersFilterByNames = astronomer.getClusters({
 *     names: ["my cluster"],
 * });
 * const exampleClustersFilterByCloudProvider = astronomer.getClusters({
 *     cloudProvider: "AWS",
 * });
 * export const clusters = exampleClusters;
 * ```
 */
export function getClustersOutput(args?: GetClustersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClustersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("astronomer:index/getClusters:getClusters", {
        "cloudProvider": args.cloudProvider,
        "names": args.names,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusters.
 */
export interface GetClustersOutputArgs {
    cloudProvider?: pulumi.Input<string>;
    names?: pulumi.Input<pulumi.Input<string>[]>;
}
