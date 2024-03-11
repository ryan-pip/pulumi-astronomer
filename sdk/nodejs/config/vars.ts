// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("astronomer");

/**
 * Organization id this provider will operate on.
 */
export declare const organizationId: string | undefined;
Object.defineProperty(exports, "organizationId", {
    get() {
        return __config.get("organizationId");
    },
    enumerable: true,
});

/**
 * Astronomer API Token. Can be set with an `ASTRONOMER_API_TOKEN` env var.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

