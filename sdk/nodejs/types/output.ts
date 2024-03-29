// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ClusterK8sTag {
    /**
     * The tag's key.
     */
    key: string;
    /**
     * The tag's value.
     */
    value: string;
}

export interface ClusterMetadata {
    externalIps?: string[];
    oidcIssuerUrl?: string;
}

export interface ClusterNodePool {
    /**
     * Whether the node pool is the default node pool of the cluster.
     */
    isDefault?: boolean;
    /**
     * The maximum number of nodes that can be created in the node pool.
     */
    maxNodeCount: number;
    /**
     * The name of the node pool.
     */
    name: string;
    /**
     * The type of node instance that is used for the node pool.
     */
    nodeInstanceType: string;
}

export interface DeploymentEnvironmentVariable {
    /**
     * Whether the environment variable is a secret.
     */
    isSecret: boolean;
    /**
     * The environment variable key, used to call the value in code.
     */
    key: string;
    /**
     * The environment variable value.
     */
    value?: string;
}

export interface DeploymentWorkerQueue {
    astroMachine: string;
    id: string;
    isDefault: boolean;
    maxWorkerCount: number;
    minWorkerCount: number;
    name: string;
    workerConcurrency: number;
}

export interface GetClusterK8sTag {
    /**
     * The tag's key.
     */
    key: string;
    /**
     * The tag's value.
     */
    value: string;
}

export interface GetClusterMetadata {
    externalIps?: string[];
    oidcIssuerUrl?: string;
}

export interface GetClusterNodePool {
    /**
     * Whether the node pool is the default node pool of the cluster.
     */
    isDefault?: boolean;
    /**
     * The maximum number of nodes that can be created in the node pool.
     */
    maxNodeCount: number;
    /**
     * The name of the node pool.
     */
    name: string;
    /**
     * The type of node instance that is used for the node pool.
     */
    nodeInstanceType: string;
}

export interface GetOrganizationManagedDomain {
    createdAt: string;
    enforcedLogins: string[];
    id: boolean;
    name: string;
    status: string;
    updatedAt: string;
}

