// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace RyanPip.Astronomer.Inputs
{

    public sealed class TeamRolesWorkspaceRoleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The role to assign to the workspace
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The ID of the workspace to assign the role to
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        public TeamRolesWorkspaceRoleGetArgs()
        {
        }
        public static new TeamRolesWorkspaceRoleGetArgs Empty => new TeamRolesWorkspaceRoleGetArgs();
    }
}
