// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace RyanPip.Astronomer.Outputs
{

    [OutputType]
    public sealed class GetClusterMetadataResult
    {
        /// <summary>
        /// Cluster external IPs
        /// </summary>
        public readonly ImmutableArray<string> ExternalIps;
        /// <summary>
        /// Cluster OIDC issuer URL
        /// </summary>
        public readonly string OidcIssuerUrl;

        [OutputConstructor]
        private GetClusterMetadataResult(
            ImmutableArray<string> externalIps,

            string oidcIssuerUrl)
        {
            ExternalIps = externalIps;
            OidcIssuerUrl = oidcIssuerUrl;
        }
    }
}
