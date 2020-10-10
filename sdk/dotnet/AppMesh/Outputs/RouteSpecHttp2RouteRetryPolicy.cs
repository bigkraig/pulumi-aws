// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh.Outputs
{

    [OutputType]
    public sealed class RouteSpecHttp2RouteRetryPolicy
    {
        /// <summary>
        /// List of HTTP retry events.
        /// Valid values: `client-error` (HTTP status code 409), `gateway-error` (HTTP status codes 502, 503, and 504), `server-error` (HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511), `stream-error` (retry on refused stream).
        /// Valid values: `client-error` (HTTP status code 409), `gateway-error` (HTTP status codes 502, 503, and 504), `server-error` (HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511), `stream-error` (retry on refused stream).
        /// </summary>
        public readonly ImmutableArray<string> HttpRetryEvents;
        /// <summary>
        /// The maximum number of retries.
        /// </summary>
        public readonly int MaxRetries;
        /// <summary>
        /// The per-retry timeout.
        /// </summary>
        public readonly Outputs.RouteSpecHttp2RouteRetryPolicyPerRetryTimeout PerRetryTimeout;
        /// <summary>
        /// List of TCP retry events. The only valid value is `connection-error`.
        /// </summary>
        public readonly ImmutableArray<string> TcpRetryEvents;

        [OutputConstructor]
        private RouteSpecHttp2RouteRetryPolicy(
            ImmutableArray<string> httpRetryEvents,

            int maxRetries,

            Outputs.RouteSpecHttp2RouteRetryPolicyPerRetryTimeout perRetryTimeout,

            ImmutableArray<string> tcpRetryEvents)
        {
            HttpRetryEvents = httpRetryEvents;
            MaxRetries = maxRetries;
            PerRetryTimeout = perRetryTimeout;
            TcpRetryEvents = tcpRetryEvents;
        }
    }
}