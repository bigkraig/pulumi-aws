// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage the accepter's side of a Direct Connect hosted public virtual interface.
 * This resource accepts ownership of a public virtual interface created by another AWS account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const accepter = new aws.Provider("accepter", {});
 * // Accepter's credentials.
 * const accepterCallerIdentity = aws.getCallerIdentity({});
 * // Creator's side of the VIF
 * const creator = new aws.directconnect.HostedPublicVirtualInterface("creator", {
 *     connectionId: "dxcon-zzzzzzzz",
 *     ownerAccountId: accepterCallerIdentity.then(accepterCallerIdentity => accepterCallerIdentity.accountId),
 *     vlan: 4094,
 *     addressFamily: "ipv4",
 *     bgpAsn: 65352,
 *     customerAddress: "175.45.176.1/30",
 *     amazonAddress: "175.45.176.2/30",
 *     routeFilterPrefixes: [
 *         "210.52.109.0/24",
 *         "175.45.176.0/22",
 *     ],
 * });
 * // Accepter's side of the VIF.
 * const accepterHostedPublicVirtualInterfaceAccepter = new aws.directconnect.HostedPublicVirtualInterfaceAccepter("accepterHostedPublicVirtualInterfaceAccepter", {
 *     virtualInterfaceId: creator.id,
 *     tags: {
 *         Side: "Accepter",
 *     },
 * }, {
 *     provider: aws.accepter,
 * });
 * ```
 */
export class HostedPublicVirtualInterfaceAccepter extends pulumi.CustomResource {
    /**
     * Get an existing HostedPublicVirtualInterfaceAccepter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostedPublicVirtualInterfaceAccepterState, opts?: pulumi.CustomResourceOptions): HostedPublicVirtualInterfaceAccepter {
        return new HostedPublicVirtualInterfaceAccepter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter';

    /**
     * Returns true if the given object is an instance of HostedPublicVirtualInterfaceAccepter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostedPublicVirtualInterfaceAccepter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostedPublicVirtualInterfaceAccepter.__pulumiType;
    }

    /**
     * The ARN of the virtual interface.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    public readonly virtualInterfaceId!: pulumi.Output<string>;

    /**
     * Create a HostedPublicVirtualInterfaceAccepter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostedPublicVirtualInterfaceAccepterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostedPublicVirtualInterfaceAccepterArgs | HostedPublicVirtualInterfaceAccepterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HostedPublicVirtualInterfaceAccepterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualInterfaceId"] = state ? state.virtualInterfaceId : undefined;
        } else {
            const args = argsOrState as HostedPublicVirtualInterfaceAccepterArgs | undefined;
            if (!args || args.virtualInterfaceId === undefined) {
                throw new Error("Missing required property 'virtualInterfaceId'");
            }
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualInterfaceId"] = args ? args.virtualInterfaceId : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HostedPublicVirtualInterfaceAccepter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostedPublicVirtualInterfaceAccepter resources.
 */
export interface HostedPublicVirtualInterfaceAccepterState {
    /**
     * The ARN of the virtual interface.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    readonly virtualInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HostedPublicVirtualInterfaceAccepter resource.
 */
export interface HostedPublicVirtualInterfaceAccepterArgs {
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Direct Connect virtual interface to accept.
     */
    readonly virtualInterfaceId: pulumi.Input<string>;
}
