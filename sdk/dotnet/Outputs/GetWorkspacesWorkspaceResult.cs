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
    public sealed class GetWorkspacesWorkspaceResult
    {
        /// <summary>
        /// Whether new Deployments enforce CI/CD deploys by default
        /// </summary>
        public readonly bool CicdEnforcedDefault;
        /// <summary>
        /// Workspace creation timestamp
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Workspace creator
        /// </summary>
        public readonly Outputs.GetWorkspacesWorkspaceCreatedByResult CreatedBy;
        /// <summary>
        /// Workspace description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Workspace identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Workspace name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Workspace last updated timestamp
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// Workspace updater
        /// </summary>
        public readonly Outputs.GetWorkspacesWorkspaceUpdatedByResult UpdatedBy;

        [OutputConstructor]
        private GetWorkspacesWorkspaceResult(
            bool cicdEnforcedDefault,

            string createdAt,

            Outputs.GetWorkspacesWorkspaceCreatedByResult createdBy,

            string description,

            string id,

            string name,

            string updatedAt,

            Outputs.GetWorkspacesWorkspaceUpdatedByResult updatedBy)
        {
            CicdEnforcedDefault = cicdEnforcedDefault;
            CreatedAt = createdAt;
            CreatedBy = createdBy;
            Description = description;
            Id = id;
            Name = name;
            UpdatedAt = updatedAt;
            UpdatedBy = updatedBy;
        }
    }
}