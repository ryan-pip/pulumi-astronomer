// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ClusterOptions data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleClusterOptions = astronomer.getClusterOptions({
 *     type: "HYBRID",
 * });
 * const exampleClusterOptionsFilterByProvider = astronomer.getClusterOptions({
 *     type: "HYBRID",
 *     cloudProvider: "AWS",
 * });
 * export const clusterOptions = exampleClusterOptions;
 * ```
 */
export function getClusterOptions(args: GetClusterOptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterOptionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("astronomer:index/getClusterOptions:getClusterOptions", {
        "cloudProvider": args.cloudProvider,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterOptions.
 */
export interface GetClusterOptionsArgs {
    cloudProvider?: string;
    type: string;
}

/**
 * A collection of values returned by getClusterOptions.
 */
export interface GetClusterOptionsResult {
    readonly cloudProvider?: string;
    readonly clusterOptions: outputs.GetClusterOptionsClusterOption[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly type: string;
}
/**
 * ClusterOptions data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleClusterOptions = astronomer.getClusterOptions({
 *     type: "HYBRID",
 * });
 * const exampleClusterOptionsFilterByProvider = astronomer.getClusterOptions({
 *     type: "HYBRID",
 *     cloudProvider: "AWS",
 * });
 * export const clusterOptions = exampleClusterOptions;
 * ```
 */
export function getClusterOptionsOutput(args: GetClusterOptionsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClusterOptionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("astronomer:index/getClusterOptions:getClusterOptions", {
        "cloudProvider": args.cloudProvider,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getClusterOptions.
 */
export interface GetClusterOptionsOutputArgs {
    cloudProvider?: pulumi.Input<string>;
    type: pulumi.Input<string>;
}
