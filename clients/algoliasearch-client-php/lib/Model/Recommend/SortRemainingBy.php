<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Recommend;

/**
 * SortRemainingBy Class Doc Comment.
 *
 * @category Class
 * @description Order of facet values that aren&#39;t explicitly positioned with the &#x60;order&#x60; setting.  - &#x60;count&#x60;.   Order remaining facet values by decreasing count.   The count is the number of matching records containing this facet value.  - &#x60;alpha&#x60;.   Sort facet values alphabetically.  - &#x60;hidden&#x60;.   Don&#39;t show facet values that aren&#39;t explicitly positioned.
 */
class SortRemainingBy
{
    /**
     * Possible values of this enum.
     */
    public const COUNT = 'count';

    public const ALPHA = 'alpha';

    public const HIDDEN = 'hidden';

    /**
     * Gets allowable values of the enum.
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::COUNT,
            self::ALPHA,
            self::HIDDEN,
        ];
    }
}
