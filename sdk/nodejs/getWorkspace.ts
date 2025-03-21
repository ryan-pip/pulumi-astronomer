// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Workspace data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleWorkspace = astronomer.getWorkspace({
 *     id: "clozc036j01to01jrlgvueo8t",
 * });
 * export const workspace = exampleWorkspace;
 * ```
 */
export function getWorkspace(args: GetWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("astronomer:index/getWorkspace:getWorkspace", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getWorkspace.
 */
export interface GetWorkspaceArgs {
    /**
     * Workspace identifier
     */
    id: string;
}

/**
 * A collection of values returned by getWorkspace.
 */
export interface GetWorkspaceResult {
    /**
     * Whether new Deployments enforce CI/CD deploys by default
     */
    readonly cicdEnforcedDefault: boolean;
    /**
     * Workspace creation timestamp
     */
    readonly createdAt: string;
    /**
     * Workspace creator
     */
    readonly createdBy: outputs.GetWorkspaceCreatedBy;
    /**
     * Workspace description
     */
    readonly description: string;
    /**
     * Workspace identifier
     */
    readonly id: string;
    /**
     * Workspace name
     */
    readonly name: string;
    /**
     * Workspace last updated timestamp
     */
    readonly updatedAt: string;
    /**
     * Workspace updater
     */
    readonly updatedBy: outputs.GetWorkspaceUpdatedBy;
}
/**
 * Workspace data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleWorkspace = astronomer.getWorkspace({
 *     id: "clozc036j01to01jrlgvueo8t",
 * });
 * export const workspace = exampleWorkspace;
 * ```
 */
export function getWorkspaceOutput(args: GetWorkspaceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWorkspaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("astronomer:index/getWorkspace:getWorkspace", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getWorkspace.
 */
export interface GetWorkspaceOutputArgs {
    /**
     * Workspace identifier
     */
    id: pulumi.Input<string>;
}
