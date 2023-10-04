// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { Discount } from './discount';
import type { Price } from './price';

export type ObjectDataAfterSearch = {
  /**
   * ID of the query that this specific record is attributable to. Used to track purchase events with multiple items originating from different searches.
   */
  queryID?: string;

  price?: Price;

  /**
   * The quantity of the purchased or added-to-cart item. The total value of a purchase is the sum of `quantity` multiplied with the `price` for each purchased item.
   */
  quantity?: number;

  discount?: Discount;
};
