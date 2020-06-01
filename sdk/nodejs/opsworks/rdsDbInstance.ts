// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an OpsWorks RDS DB Instance resource.
 *
 * > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myInstance = new aws.opsworks.RdsDbInstance("myInstance", {
 *     dbPassword: "somePass",
 *     dbUser: "someUser",
 *     rdsDbInstanceArn: aws_db_instance_my_instance.arn,
 *     stackId: aws_opsworks_stack_my_stack.id,
 * });
 * ```
 */
export class RdsDbInstance extends pulumi.CustomResource {
    /**
     * Get an existing RdsDbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RdsDbInstanceState, opts?: pulumi.CustomResourceOptions): RdsDbInstance {
        return new RdsDbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opsworks/rdsDbInstance:RdsDbInstance';

    /**
     * Returns true if the given object is an instance of RdsDbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RdsDbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RdsDbInstance.__pulumiType;
    }

    /**
     * A db password
     */
    public readonly dbPassword!: pulumi.Output<string>;
    /**
     * A db username
     */
    public readonly dbUser!: pulumi.Output<string>;
    /**
     * The db instance to register for this stack. Changing this will force a new resource.
     */
    public readonly rdsDbInstanceArn!: pulumi.Output<string>;
    /**
     * The stack to register a db instance for. Changing this will force a new resource.
     */
    public readonly stackId!: pulumi.Output<string>;

    /**
     * Create a RdsDbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RdsDbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RdsDbInstanceArgs | RdsDbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RdsDbInstanceState | undefined;
            inputs["dbPassword"] = state ? state.dbPassword : undefined;
            inputs["dbUser"] = state ? state.dbUser : undefined;
            inputs["rdsDbInstanceArn"] = state ? state.rdsDbInstanceArn : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
        } else {
            const args = argsOrState as RdsDbInstanceArgs | undefined;
            if (!args || args.dbPassword === undefined) {
                throw new Error("Missing required property 'dbPassword'");
            }
            if (!args || args.dbUser === undefined) {
                throw new Error("Missing required property 'dbUser'");
            }
            if (!args || args.rdsDbInstanceArn === undefined) {
                throw new Error("Missing required property 'rdsDbInstanceArn'");
            }
            if (!args || args.stackId === undefined) {
                throw new Error("Missing required property 'stackId'");
            }
            inputs["dbPassword"] = args ? args.dbPassword : undefined;
            inputs["dbUser"] = args ? args.dbUser : undefined;
            inputs["rdsDbInstanceArn"] = args ? args.rdsDbInstanceArn : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RdsDbInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RdsDbInstance resources.
 */
export interface RdsDbInstanceState {
    /**
     * A db password
     */
    readonly dbPassword?: pulumi.Input<string>;
    /**
     * A db username
     */
    readonly dbUser?: pulumi.Input<string>;
    /**
     * The db instance to register for this stack. Changing this will force a new resource.
     */
    readonly rdsDbInstanceArn?: pulumi.Input<string>;
    /**
     * The stack to register a db instance for. Changing this will force a new resource.
     */
    readonly stackId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RdsDbInstance resource.
 */
export interface RdsDbInstanceArgs {
    /**
     * A db password
     */
    readonly dbPassword: pulumi.Input<string>;
    /**
     * A db username
     */
    readonly dbUser: pulumi.Input<string>;
    /**
     * The db instance to register for this stack. Changing this will force a new resource.
     */
    readonly rdsDbInstanceArn: pulumi.Input<string>;
    /**
     * The stack to register a db instance for. Changing this will force a new resource.
     */
    readonly stackId: pulumi.Input<string>;
}
