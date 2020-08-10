// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Service Discovery Private DNS Namespace resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
 * const examplePrivateDnsNamespace = new aws.servicediscovery.PrivateDnsNamespace("examplePrivateDnsNamespace", {
 *     description: "example",
 *     vpc: exampleVpc.id,
 * });
 * ```
 */
export class PrivateDnsNamespace extends pulumi.CustomResource {
    /**
     * Get an existing PrivateDnsNamespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateDnsNamespaceState, opts?: pulumi.CustomResourceOptions): PrivateDnsNamespace {
        return new PrivateDnsNamespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace';

    /**
     * Returns true if the given object is an instance of PrivateDnsNamespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateDnsNamespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateDnsNamespace.__pulumiType;
    }

    /**
     * The ARN that Amazon Route 53 assigns to the namespace when you create it.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description that you specify for the namespace when you create it.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
     */
    public /*out*/ readonly hostedZone!: pulumi.Output<string>;
    /**
     * The name of the namespace.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the namespace.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of VPC that you want to associate the namespace with.
     */
    public readonly vpc!: pulumi.Output<string>;

    /**
     * Create a PrivateDnsNamespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateDnsNamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateDnsNamespaceArgs | PrivateDnsNamespaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PrivateDnsNamespaceState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["hostedZone"] = state ? state.hostedZone : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpc"] = state ? state.vpc : undefined;
        } else {
            const args = argsOrState as PrivateDnsNamespaceArgs | undefined;
            if (!args || args.vpc === undefined) {
                throw new Error("Missing required property 'vpc'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpc"] = args ? args.vpc : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["hostedZone"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PrivateDnsNamespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateDnsNamespace resources.
 */
export interface PrivateDnsNamespaceState {
    /**
     * The ARN that Amazon Route 53 assigns to the namespace when you create it.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The description that you specify for the namespace when you create it.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
     */
    readonly hostedZone?: pulumi.Input<string>;
    /**
     * The name of the namespace.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the namespace.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of VPC that you want to associate the namespace with.
     */
    readonly vpc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateDnsNamespace resource.
 */
export interface PrivateDnsNamespaceArgs {
    /**
     * The description that you specify for the namespace when you create it.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the namespace.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the namespace.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of VPC that you want to associate the namespace with.
     */
    readonly vpc: pulumi.Input<string>;
}
