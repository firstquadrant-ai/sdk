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
 * @interface Team
 */
export interface Team {
    /**
     * 
     * @type {any}
     * @memberof Team
     */
    id: any | null;
    /**
     * 
     * @type {any}
     * @memberof Team
     */
    name: any | null;
    /**
     * 
     * @type {any}
     * @memberof Team
     */
    createdAt: any | null;
    /**
     * 
     * @type {any}
     * @memberof Team
     */
    updatedAt: any | null;
    /**
     * 
     * @type {any}
     * @memberof Team
     */
    timeZone: any | null;
}

/**
 * Check if a given object implements the Team interface.
 */
export function instanceOfTeam(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "createdAt" in value;
    isInstance = isInstance && "updatedAt" in value;
    isInstance = isInstance && "timeZone" in value;

    return isInstance;
}

export function TeamFromJSON(json: any): Team {
    return TeamFromJSONTyped(json, false);
}

export function TeamFromJSONTyped(json: any, ignoreDiscriminator: boolean): Team {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'createdAt': json['createdAt'],
        'updatedAt': json['updatedAt'],
        'timeZone': json['timeZone'],
    };
}

export function TeamToJSON(value?: Team | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'createdAt': value.createdAt,
        'updatedAt': value.updatedAt,
        'timeZone': value.timeZone,
    };
}
