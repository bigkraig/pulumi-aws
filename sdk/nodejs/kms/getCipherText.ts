// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The KMS ciphertext data source allows you to encrypt plaintext into ciphertext
 * by using an AWS KMS customer master key. The value returned by this data source
 * changes every apply. For a stable ciphertext value, see the [`aws.kms.Ciphertext`
 * resource](https://www.terraform.io/docs/providers/aws/r/kms_ciphertext.html).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const oauthConfig = new aws.kms.Key("oauthConfig", {
 *     description: "oauth config",
 *     isEnabled: true,
 * });
 * const oauth = oauthConfig.keyId.apply(keyId => aws.kms.getCipherText({
 *     keyId: keyId,
 *     plaintext: `{
 *   "clientId": "e587dbae22222f55da22",
 *   "clientSecret": "8289575d00000ace55e1815ec13673955721b8a5"
 * }
 * `,
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/kms_ciphertext.html.markdown.
 */
export function getCipherText(args: GetCipherTextArgs, opts?: pulumi.InvokeOptions): Promise<GetCipherTextResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:kms/getCipherText:getCipherText", {
        "context": args.context,
        "keyId": args.keyId,
        "plaintext": args.plaintext,
    }, opts);
}

/**
 * A collection of arguments for invoking getCipherText.
 */
export interface GetCipherTextArgs {
    /**
     * An optional mapping that makes up the encryption context.
     */
    readonly context?: {[key: string]: string};
    /**
     * Globally unique key ID for the customer master key.
     */
    readonly keyId: string;
    /**
     * Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
     */
    readonly plaintext: string;
}

/**
 * A collection of values returned by getCipherText.
 */
export interface GetCipherTextResult {
    /**
     * Base64 encoded ciphertext
     */
    readonly ciphertextBlob: string;
    readonly context?: {[key: string]: string};
    readonly keyId: string;
    readonly plaintext: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
