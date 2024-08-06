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
    public sealed class GetOrganizationCreatedByResult
    {
        public readonly string ApiTokenName;
        public readonly string AvatarUrl;
        public readonly string FullName;
        public readonly string Id;
        public readonly string SubjectType;
        public readonly string Username;

        [OutputConstructor]
        private GetOrganizationCreatedByResult(
            string apiTokenName,

            string avatarUrl,

            string fullName,

            string id,

            string subjectType,

            string username)
        {
            ApiTokenName = apiTokenName;
            AvatarUrl = avatarUrl;
            FullName = fullName;
            Id = id;
            SubjectType = subjectType;
            Username = username;
        }
    }
}
