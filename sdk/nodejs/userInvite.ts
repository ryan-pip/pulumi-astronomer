// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * User Invite resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@ryan-pip/pulumi_astronomer";
 *
 * const userInvite = new astronomer.UserInvite("userInvite", {
 *     email: "email@organization.com",
 *     role: "ORGANIZATION_MEMBER",
 * });
 * ```
 */
export class UserInvite extends pulumi.CustomResource {
    /**
     * Get an existing UserInvite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserInviteState, opts?: pulumi.CustomResourceOptions): UserInvite {
        return new UserInvite(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'astronomer:index/userInvite:UserInvite';

    /**
     * Returns true if the given object is an instance of UserInvite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserInvite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserInvite.__pulumiType;
    }

    /**
     * The email address of the user being invited
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The expiration date of the invite
     */
    public /*out*/ readonly expiresAt!: pulumi.Output<string>;
    /**
     * The ID of the invite
     */
    public /*out*/ readonly inviteId!: pulumi.Output<string>;
    /**
     * The profile of the invitee
     */
    public /*out*/ readonly invitee!: pulumi.Output<outputs.UserInviteInvitee>;
    /**
     * The profile of the inviter
     */
    public /*out*/ readonly inviter!: pulumi.Output<outputs.UserInviteInviter>;
    /**
     * The Organization role to assign to the user
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The ID of the user
     */
    public /*out*/ readonly userId!: pulumi.Output<string>;

    /**
     * Create a UserInvite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserInviteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserInviteArgs | UserInviteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserInviteState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["inviteId"] = state ? state.inviteId : undefined;
            resourceInputs["invitee"] = state ? state.invitee : undefined;
            resourceInputs["inviter"] = state ? state.inviter : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserInviteArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["inviteId"] = undefined /*out*/;
            resourceInputs["invitee"] = undefined /*out*/;
            resourceInputs["inviter"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserInvite.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserInvite resources.
 */
export interface UserInviteState {
    /**
     * The email address of the user being invited
     */
    email?: pulumi.Input<string>;
    /**
     * The expiration date of the invite
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID of the invite
     */
    inviteId?: pulumi.Input<string>;
    /**
     * The profile of the invitee
     */
    invitee?: pulumi.Input<inputs.UserInviteInvitee>;
    /**
     * The profile of the inviter
     */
    inviter?: pulumi.Input<inputs.UserInviteInviter>;
    /**
     * The Organization role to assign to the user
     */
    role?: pulumi.Input<string>;
    /**
     * The ID of the user
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserInvite resource.
 */
export interface UserInviteArgs {
    /**
     * The email address of the user being invited
     */
    email: pulumi.Input<string>;
    /**
     * The Organization role to assign to the user
     */
    role: pulumi.Input<string>;
}
