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
    public static class GetWorkspace
    {
        /// <summary>
        /// Astronomer Workspace Resource
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Astronomer = Pulumi.Astronomer;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var importedWorkspace = Astronomer.GetWorkspace.Invoke(new()
        ///     {
        ///         Id = "cabcabcabcabcabcabcabcabcabc",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("astronomer:index/getWorkspace:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithDefaults());

        /// <summary>
        /// Astronomer Workspace Resource
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Astronomer = Pulumi.Astronomer;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var importedWorkspace = Astronomer.GetWorkspace.Invoke(new()
        ///     {
        ///         Id = "cabcabcabcabcabcabcabcabcabc",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetWorkspaceResult> Invoke(GetWorkspaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkspaceResult>("astronomer:index/getWorkspace:getWorkspace", args ?? new GetWorkspaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkspaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Workspace's identifier.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
        public static new GetWorkspaceArgs Empty => new GetWorkspaceArgs();
    }

    public sealed class GetWorkspaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Workspace's identifier.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetWorkspaceInvokeArgs()
        {
        }
        public static new GetWorkspaceInvokeArgs Empty => new GetWorkspaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// Whether new Deployments enforce CI/CD deploys by default.
        /// </summary>
        public readonly bool CicdEnforcedDefault;
        /// <summary>
        /// The Workspace's description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Workspace's identifier.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Workspace's name
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetWorkspaceResult(
            bool cicdEnforcedDefault,

            string description,

            string id,

            string name)
        {
            CicdEnforcedDefault = cicdEnforcedDefault;
            Description = description;
            Id = id;
            Name = name;
        }
    }
}
