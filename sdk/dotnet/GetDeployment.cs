// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace RyanPip.Astronomer
{
    public static class GetDeployment
    {
        /// <summary>
        /// Astronomer Deployment Resource
        /// </summary>
        public static Task<GetDeploymentResult> InvokeAsync(GetDeploymentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentResult>("astronomer:index/getDeployment:getDeployment", args ?? new GetDeploymentArgs(), options.WithDefaults());

        /// <summary>
        /// Astronomer Deployment Resource
        /// </summary>
        public static Output<GetDeploymentResult> Invoke(GetDeploymentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentResult>("astronomer:index/getDeployment:getDeployment", args ?? new GetDeploymentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeploymentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Deployment's Identifier
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDeploymentArgs()
        {
        }
        public static new GetDeploymentArgs Empty => new GetDeploymentArgs();
    }

    public sealed class GetDeploymentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Deployment's Identifier
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDeploymentInvokeArgs()
        {
        }
        public static new GetDeploymentInvokeArgs Empty => new GetDeploymentInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeploymentResult
    {
        /// <summary>
        /// The Deployment's Astro Runtime version.
        /// </summary>
        public readonly string AirflowVersion;
        /// <summary>
        /// The cloud provider for the Deployment's cluster. Optional if `ClusterId` is specified.
        /// </summary>
        public readonly string CloudProvider;
        /// <summary>
        /// The ID of the cluster to which the Deployment will be created in. Optional if cloud provider and region is specified.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Cluster Name
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// The Deployment's description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Deployment's Identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the Deployment requires that all deploys are made through CI/CD.
        /// </summary>
        public readonly bool IsCicdEnforced;
        /// <summary>
        /// The Deployment's name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetDeploymentResult(
            string airflowVersion,

            string cloudProvider,

            string clusterId,

            string clusterName,

            string description,

            string id,

            bool isCicdEnforced,

            string name)
        {
            AirflowVersion = airflowVersion;
            CloudProvider = cloudProvider;
            ClusterId = clusterId;
            ClusterName = clusterName;
            Description = description;
            Id = id;
            IsCicdEnforced = isCicdEnforced;
            Name = name;
        }
    }
}
