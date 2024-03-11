// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Astronomer Workspace Resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@ryan-pip/pulumi_astronomer";
 *
 * const completeSetup = new astronomer.Workspace("completeSetup", {
 *     cicdEnforcedDefault: true,
 *     description: "Testing Workspace",
 * });
 * ```
 */
export class Workspace extends pulumi.CustomResource {
    /**
     * Get an existing Workspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkspaceState, opts?: pulumi.CustomResourceOptions): Workspace {
        return new Workspace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'astronomer:index/workspace:Workspace';

    /**
     * Returns true if the given object is an instance of Workspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workspace.__pulumiType;
    }

    /**
     * Whether new Deployments enforce CI/CD deploys by default.
     */
    public readonly cicdEnforcedDefault!: pulumi.Output<boolean | undefined>;
    /**
     * The Workspace's description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Workspace's name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkspaceArgs | WorkspaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkspaceState | undefined;
            resourceInputs["cicdEnforcedDefault"] = state ? state.cicdEnforcedDefault : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as WorkspaceArgs | undefined;
            resourceInputs["cicdEnforcedDefault"] = args ? args.cicdEnforcedDefault : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Workspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workspace resources.
 */
export interface WorkspaceState {
    /**
     * Whether new Deployments enforce CI/CD deploys by default.
     */
    cicdEnforcedDefault?: pulumi.Input<boolean>;
    /**
     * The Workspace's description.
     */
    description?: pulumi.Input<string>;
    /**
     * The Workspace's name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * Whether new Deployments enforce CI/CD deploys by default.
     */
    cicdEnforcedDefault?: pulumi.Input<boolean>;
    /**
     * The Workspace's description.
     */
    description?: pulumi.Input<string>;
    /**
     * The Workspace's name.
     */
    name?: pulumi.Input<string>;
}