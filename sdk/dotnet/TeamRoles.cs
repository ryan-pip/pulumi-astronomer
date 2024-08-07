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
    /// <summary>
    /// Team Roles resource
    /// </summary>
    [AstronomerResourceType("astronomer:index/teamRoles:TeamRoles")]
    public partial class TeamRoles : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The roles to assign to the deployments
        /// </summary>
        [Output("deploymentRoles")]
        public Output<ImmutableArray<Outputs.TeamRolesDeploymentRole>> DeploymentRoles { get; private set; } = null!;

        /// <summary>
        /// The role to assign to the organization
        /// </summary>
        [Output("organizationRole")]
        public Output<string> OrganizationRole { get; private set; } = null!;

        /// <summary>
        /// The ID of the team to assign the roles to
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;

        /// <summary>
        /// The roles to assign to the workspaces
        /// </summary>
        [Output("workspaceRoles")]
        public Output<ImmutableArray<Outputs.TeamRolesWorkspaceRole>> WorkspaceRoles { get; private set; } = null!;


        /// <summary>
        /// Create a TeamRoles resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TeamRoles(string name, TeamRolesArgs args, CustomResourceOptions? options = null)
            : base("astronomer:index/teamRoles:TeamRoles", name, args ?? new TeamRolesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TeamRoles(string name, Input<string> id, TeamRolesState? state = null, CustomResourceOptions? options = null)
            : base("astronomer:index/teamRoles:TeamRoles", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ryan-pip/pulumi-astronomer",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TeamRoles resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TeamRoles Get(string name, Input<string> id, TeamRolesState? state = null, CustomResourceOptions? options = null)
        {
            return new TeamRoles(name, id, state, options);
        }
    }

    public sealed class TeamRolesArgs : global::Pulumi.ResourceArgs
    {
        [Input("deploymentRoles")]
        private InputList<Inputs.TeamRolesDeploymentRoleArgs>? _deploymentRoles;

        /// <summary>
        /// The roles to assign to the deployments
        /// </summary>
        public InputList<Inputs.TeamRolesDeploymentRoleArgs> DeploymentRoles
        {
            get => _deploymentRoles ?? (_deploymentRoles = new InputList<Inputs.TeamRolesDeploymentRoleArgs>());
            set => _deploymentRoles = value;
        }

        /// <summary>
        /// The role to assign to the organization
        /// </summary>
        [Input("organizationRole", required: true)]
        public Input<string> OrganizationRole { get; set; } = null!;

        /// <summary>
        /// The ID of the team to assign the roles to
        /// </summary>
        [Input("teamId", required: true)]
        public Input<string> TeamId { get; set; } = null!;

        [Input("workspaceRoles")]
        private InputList<Inputs.TeamRolesWorkspaceRoleArgs>? _workspaceRoles;

        /// <summary>
        /// The roles to assign to the workspaces
        /// </summary>
        public InputList<Inputs.TeamRolesWorkspaceRoleArgs> WorkspaceRoles
        {
            get => _workspaceRoles ?? (_workspaceRoles = new InputList<Inputs.TeamRolesWorkspaceRoleArgs>());
            set => _workspaceRoles = value;
        }

        public TeamRolesArgs()
        {
        }
        public static new TeamRolesArgs Empty => new TeamRolesArgs();
    }

    public sealed class TeamRolesState : global::Pulumi.ResourceArgs
    {
        [Input("deploymentRoles")]
        private InputList<Inputs.TeamRolesDeploymentRoleGetArgs>? _deploymentRoles;

        /// <summary>
        /// The roles to assign to the deployments
        /// </summary>
        public InputList<Inputs.TeamRolesDeploymentRoleGetArgs> DeploymentRoles
        {
            get => _deploymentRoles ?? (_deploymentRoles = new InputList<Inputs.TeamRolesDeploymentRoleGetArgs>());
            set => _deploymentRoles = value;
        }

        /// <summary>
        /// The role to assign to the organization
        /// </summary>
        [Input("organizationRole")]
        public Input<string>? OrganizationRole { get; set; }

        /// <summary>
        /// The ID of the team to assign the roles to
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        [Input("workspaceRoles")]
        private InputList<Inputs.TeamRolesWorkspaceRoleGetArgs>? _workspaceRoles;

        /// <summary>
        /// The roles to assign to the workspaces
        /// </summary>
        public InputList<Inputs.TeamRolesWorkspaceRoleGetArgs> WorkspaceRoles
        {
            get => _workspaceRoles ?? (_workspaceRoles = new InputList<Inputs.TeamRolesWorkspaceRoleGetArgs>());
            set => _workspaceRoles = value;
        }

        public TeamRolesState()
        {
        }
        public static new TeamRolesState Empty => new TeamRolesState();
    }
}
