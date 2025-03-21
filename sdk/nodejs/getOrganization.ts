// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Organization data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleOrganization = astronomer.getOrganization({});
 * export const organization = exampleOrganization;
 * ```
 */
export function getOrganization(opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("astronomer:index/getOrganization:getOrganization", {
    }, opts);
}

/**
 * A collection of values returned by getOrganization.
 */
export interface GetOrganizationResult {
    /**
     * Organization billing email
     */
    readonly billingEmail: string;
    /**
     * Organization creation timestamp
     */
    readonly createdAt: string;
    /**
     * Organization creator
     */
    readonly createdBy: outputs.GetOrganizationCreatedBy;
    /**
     * Organization identifier
     */
    readonly id: string;
    /**
     * Whether SCIM is enabled for the organization
     */
    readonly isScimEnabled: boolean;
    /**
     * Organization name
     */
    readonly name: string;
    /**
     * Organization payment method
     */
    readonly paymentMethod: string;
    /**
     * Organization product type
     */
    readonly product: string;
    /**
     * Organization status
     */
    readonly status: string;
    /**
     * Organization support plan
     */
    readonly supportPlan: string;
    /**
     * Organization trial expiration timestamp
     */
    readonly trialExpiresAt: string;
    /**
     * Organization last updated timestamp
     */
    readonly updatedAt: string;
    /**
     * Organization updater
     */
    readonly updatedBy: outputs.GetOrganizationUpdatedBy;
}
/**
 * Organization data source
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as astronomer from "@pulumi/astronomer";
 *
 * const exampleOrganization = astronomer.getOrganization({});
 * export const organization = exampleOrganization;
 * ```
 */
export function getOrganizationOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOrganizationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("astronomer:index/getOrganization:getOrganization", {
    }, opts);
}
