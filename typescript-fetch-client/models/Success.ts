/* tslint:disable */
/* eslint-disable */
/**
 * FirstQuadrant API
 * The FirstQuadrant API is used to interact with FirstQuadrant programmatically. We also have SDKs available (coming soon).
 *
 * The version of the OpenAPI document: 0.0.0
 * Contact: inquiry@firstquadrant.ai
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface Success
 */
export interface Success {
    /**
     * 
     * @type {any}
     * @memberof Success
     */
    success: SuccessSuccessEnum;
}


/**
 * @export
 */
export const SuccessSuccessEnum = {
    True: 'true'
} as const;
export type SuccessSuccessEnum = typeof SuccessSuccessEnum[keyof typeof SuccessSuccessEnum];


/**
 * Check if a given object implements the Success interface.
 */
export function instanceOfSuccess(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "success" in value;

    return isInstance;
}

export function SuccessFromJSON(json: any): Success {
    return SuccessFromJSONTyped(json, false);
}

export function SuccessFromJSONTyped(json: any, ignoreDiscriminator: boolean): Success {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'success': json['success'],
    };
}

export function SuccessToJSON(value?: Success | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'success': value.success,
    };
}

