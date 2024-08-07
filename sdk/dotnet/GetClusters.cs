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
    public static class GetClusters
    {
        /// <summary>
        /// Clusters data source
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Astronomer = Pulumi.Astronomer;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleClusters = Astronomer.GetClusters.Invoke();
        /// 
        ///     var exampleClustersFilterByNames = Astronomer.GetClusters.Invoke(new()
        ///     {
        ///         Names = new[]
        ///         {
        ///             "my cluster",
        ///         },
        ///     });
        /// 
        ///     var exampleClustersFilterByCloudProvider = Astronomer.GetClusters.Invoke(new()
        ///     {
        ///         CloudProvider = "AWS",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetClustersResult> InvokeAsync(GetClustersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClustersResult>("astronomer:index/getClusters:getClusters", args ?? new GetClustersArgs(), options.WithDefaults());

        /// <summary>
        /// Clusters data source
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Astronomer = Pulumi.Astronomer;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleClusters = Astronomer.GetClusters.Invoke();
        /// 
        ///     var exampleClustersFilterByNames = Astronomer.GetClusters.Invoke(new()
        ///     {
        ///         Names = new[]
        ///         {
        ///             "my cluster",
        ///         },
        ///     });
        /// 
        ///     var exampleClustersFilterByCloudProvider = Astronomer.GetClusters.Invoke(new()
        ///     {
        ///         CloudProvider = "AWS",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetClustersResult> Invoke(GetClustersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClustersResult>("astronomer:index/getClusters:getClusters", args ?? new GetClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClustersArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloudProvider")]
        public string? CloudProvider { get; set; }

        [Input("names")]
        private List<string>? _names;
        public List<string> Names
        {
            get => _names ?? (_names = new List<string>());
            set => _names = value;
        }

        public GetClustersArgs()
        {
        }
        public static new GetClustersArgs Empty => new GetClustersArgs();
    }

    public sealed class GetClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cloudProvider")]
        public Input<string>? CloudProvider { get; set; }

        [Input("names")]
        private InputList<string>? _names;
        public InputList<string> Names
        {
            get => _names ?? (_names = new InputList<string>());
            set => _names = value;
        }

        public GetClustersInvokeArgs()
        {
        }
        public static new GetClustersInvokeArgs Empty => new GetClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetClustersResult
    {
        public readonly string? CloudProvider;
        public readonly ImmutableArray<Outputs.GetClustersClusterResult> Clusters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Names;

        [OutputConstructor]
        private GetClustersResult(
            string? cloudProvider,

            ImmutableArray<Outputs.GetClustersClusterResult> clusters,

            string id,

            ImmutableArray<string> names)
        {
            CloudProvider = cloudProvider;
            Clusters = clusters;
            Id = id;
            Names = names;
        }
    }
}