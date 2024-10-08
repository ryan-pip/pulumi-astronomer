// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Hybrid cluster workspace authorization resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@ryan-pip/pulumi_astronomer";
 *
 * const example = new astronomer.HybridClusterWorkspaceAuthorization("example", {
 *     clusterId: "clk8h0fv1006801j8yysfybbt",
 *     workspaceIds: [
 *         "cl70oe7cu445571iynrkthtybl",
 *         "cl70oe7cu445571iynrkthacsd",
 *     ],
 * });
 * const importedClusterWorkspaceAuthorization = new astronomer.HybridClusterWorkspaceAuthorization("importedClusterWorkspaceAuthorization", {
 *     clusterId: "clk8h0fv1006801j8yysfybbt",
 *     workspaceIds: ["cl70oe7cu445571iynrkthtybl"],
 * });
 * ```
 */
export class HybridClusterWorkspaceAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing HybridClusterWorkspaceAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HybridClusterWorkspaceAuthorizationState, opts?: pulumi.CustomResourceOptions): HybridClusterWorkspaceAuthorization {
        return new HybridClusterWorkspaceAuthorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'astronomer:index/hybridClusterWorkspaceAuthorization:HybridClusterWorkspaceAuthorization';

    /**
     * Returns true if the given object is an instance of HybridClusterWorkspaceAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HybridClusterWorkspaceAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HybridClusterWorkspaceAuthorization.__pulumiType;
    }

    /**
     * The ID of the hybrid cluster to set authorizations for
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The IDs of the workspaces to authorize for the hybrid cluster
     */
    public readonly workspaceIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a HybridClusterWorkspaceAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HybridClusterWorkspaceAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HybridClusterWorkspaceAuthorizationArgs | HybridClusterWorkspaceAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HybridClusterWorkspaceAuthorizationState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["workspaceIds"] = state ? state.workspaceIds : undefined;
        } else {
            const args = argsOrState as HybridClusterWorkspaceAuthorizationArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["workspaceIds"] = args ? args.workspaceIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HybridClusterWorkspaceAuthorization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HybridClusterWorkspaceAuthorization resources.
 */
export interface HybridClusterWorkspaceAuthorizationState {
    /**
     * The ID of the hybrid cluster to set authorizations for
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The IDs of the workspaces to authorize for the hybrid cluster
     */
    workspaceIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a HybridClusterWorkspaceAuthorization resource.
 */
export interface HybridClusterWorkspaceAuthorizationArgs {
    /**
     * The ID of the hybrid cluster to set authorizations for
     */
    clusterId: pulumi.Input<string>;
    /**
     * The IDs of the workspaces to authorize for the hybrid cluster
     */
    workspaceIds?: pulumi.Input<pulumi.Input<string>[]>;
}
