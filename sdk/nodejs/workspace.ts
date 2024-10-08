// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Workspace resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@ryan-pip/pulumi_astronomer";
 *
 * const example = new astronomer.Workspace("example", {
 *     description: "my first workspace",
 *     cicdEnforcedDefault: true,
 * });
 * const importedWorkspace = new astronomer.Workspace("importedWorkspace", {
 *     description: "an existing workspace",
 *     cicdEnforcedDefault: true,
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
     * Whether new Deployments enforce CI/CD deploys by default
     */
    public readonly cicdEnforcedDefault!: pulumi.Output<boolean>;
    /**
     * Workspace creation timestamp
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Workspace creator
     */
    public /*out*/ readonly createdBy!: pulumi.Output<outputs.WorkspaceCreatedBy>;
    /**
     * Workspace description
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Workspace name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Workspace last updated timestamp
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Workspace updater
     */
    public /*out*/ readonly updatedBy!: pulumi.Output<outputs.WorkspaceUpdatedBy>;

    /**
     * Create a Workspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkspaceArgs | WorkspaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkspaceState | undefined;
            resourceInputs["cicdEnforcedDefault"] = state ? state.cicdEnforcedDefault : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["createdBy"] = state ? state.createdBy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["updatedBy"] = state ? state.updatedBy : undefined;
        } else {
            const args = argsOrState as WorkspaceArgs | undefined;
            if ((!args || args.cicdEnforcedDefault === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cicdEnforcedDefault'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            resourceInputs["cicdEnforcedDefault"] = args ? args.cicdEnforcedDefault : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["updatedBy"] = undefined /*out*/;
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
     * Whether new Deployments enforce CI/CD deploys by default
     */
    cicdEnforcedDefault?: pulumi.Input<boolean>;
    /**
     * Workspace creation timestamp
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Workspace creator
     */
    createdBy?: pulumi.Input<inputs.WorkspaceCreatedBy>;
    /**
     * Workspace description
     */
    description?: pulumi.Input<string>;
    /**
     * Workspace name
     */
    name?: pulumi.Input<string>;
    /**
     * Workspace last updated timestamp
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Workspace updater
     */
    updatedBy?: pulumi.Input<inputs.WorkspaceUpdatedBy>;
}

/**
 * The set of arguments for constructing a Workspace resource.
 */
export interface WorkspaceArgs {
    /**
     * Whether new Deployments enforce CI/CD deploys by default
     */
    cicdEnforcedDefault: pulumi.Input<boolean>;
    /**
     * Workspace description
     */
    description: pulumi.Input<string>;
    /**
     * Workspace name
     */
    name?: pulumi.Input<string>;
}
