// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get the pricing information of all products in AWS.
 * This data source is only available in a us-east-1 or ap-south-1 provider.
 */
export function getProduct(args: GetProductArgs, opts?: pulumi.InvokeOptions): Promise<GetProductResult> {
    return pulumi.runtime.invoke("aws:pricing/getProduct:getProduct", {
        "filters": args.filters,
        "serviceCode": args.serviceCode,
    }, opts);
}

/**
 * A collection of arguments for invoking getProduct.
 */
export interface GetProductArgs {
    /**
     * A list of filters. Passed directly to the API (see GetProducts API reference). These filters must describe a single product, this resource will fail if more than one product is returned by the API.
     */
    readonly filters: { field: string, value: string }[];
    /**
     * The code of the service. Available service codes can be fetched using the DescribeServices pricing API call.
     */
    readonly serviceCode: string;
}

/**
 * A collection of values returned by getProduct.
 */
export interface GetProductResult {
    /**
     * Set to the product returned from the API.
     */
    readonly result: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
