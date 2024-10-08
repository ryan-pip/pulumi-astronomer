// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the astro package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'astronomer';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * API host to use for the provider. Default is `https://api.astronomer.io`
     */
    public readonly host!: pulumi.Output<string | undefined>;
    /**
     * Organization ID this provider will operate on.
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * Astro API Token. Can be set with an `ASTRO_API_TOKEN` env var.
     */
    public readonly token!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["host"] = args ? args.host : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["token"] = (args?.token ? pulumi.secret(args.token) : undefined) ?? utilities.getEnv("ASTRO_API_TOKEN");
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["token"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * API host to use for the provider. Default is `https://api.astronomer.io`
     */
    host?: pulumi.Input<string>;
    /**
     * Organization ID this provider will operate on.
     */
    organizationId: pulumi.Input<string>;
    /**
     * Astro API Token. Can be set with an `ASTRO_API_TOKEN` env var.
     */
    token?: pulumi.Input<string>;
}
