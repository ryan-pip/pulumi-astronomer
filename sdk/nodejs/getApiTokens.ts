// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * API Tokens data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleApiTokens = astronomer.getApiTokens({});
 * const organizationOnlyExample = astronomer.getApiTokens({
 *     includeOnlyOrganizationTokens: true,
 * });
 * const workspaceExample = astronomer.getApiTokens({
 *     workspaceId: "clx42sxw501gl01o0gjenthnh",
 * });
 * const deploymentExample = astronomer.getApiTokens({
 *     deploymentId: "clx44jyu001m201m5dzsbexqr",
 * });
 * export const apiTokens = exampleApiTokens;
 * ```
 */
export function getApiTokens(args?: GetApiTokensArgs, opts?: pulumi.InvokeOptions): Promise<GetApiTokensResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("astronomer:index/getApiTokens:getApiTokens", {
        "deploymentId": args.deploymentId,
        "includeOnlyOrganizationTokens": args.includeOnlyOrganizationTokens,
        "workspaceId": args.workspaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiTokens.
 */
export interface GetApiTokensArgs {
    deploymentId?: string;
    includeOnlyOrganizationTokens?: boolean;
    workspaceId?: string;
}

/**
 * A collection of values returned by getApiTokens.
 */
export interface GetApiTokensResult {
    readonly apiTokens: outputs.GetApiTokensApiToken[];
    readonly deploymentId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeOnlyOrganizationTokens?: boolean;
    readonly workspaceId?: string;
}
/**
 * API Tokens data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleApiTokens = astronomer.getApiTokens({});
 * const organizationOnlyExample = astronomer.getApiTokens({
 *     includeOnlyOrganizationTokens: true,
 * });
 * const workspaceExample = astronomer.getApiTokens({
 *     workspaceId: "clx42sxw501gl01o0gjenthnh",
 * });
 * const deploymentExample = astronomer.getApiTokens({
 *     deploymentId: "clx44jyu001m201m5dzsbexqr",
 * });
 * export const apiTokens = exampleApiTokens;
 * ```
 */
export function getApiTokensOutput(args?: GetApiTokensOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiTokensResult> {
    return pulumi.output(args).apply((a: any) => getApiTokens(a, opts))
}

/**
 * A collection of arguments for invoking getApiTokens.
 */
export interface GetApiTokensOutputArgs {
    deploymentId?: pulumi.Input<string>;
    includeOnlyOrganizationTokens?: pulumi.Input<boolean>;
    workspaceId?: pulumi.Input<string>;
}
