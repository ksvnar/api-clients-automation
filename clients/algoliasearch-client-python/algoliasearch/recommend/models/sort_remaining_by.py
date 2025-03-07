# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from enum import Enum
from json import loads
from typing import Self


class SortRemainingBy(str, Enum):
    """
    Order of facet values that aren't explicitly positioned with the `order` setting.  - `count`.   Order remaining facet values by decreasing count.   The count is the number of matching records containing this facet value.  - `alpha`.   Sort facet values alphabetically.  - `hidden`.   Don't show facet values that aren't explicitly positioned.
    """

    """
    allowed enum values
    """
    COUNT = "count"
    ALPHA = "alpha"
    HIDDEN = "hidden"

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of SortRemainingBy from a JSON string"""
        return cls(loads(json_str))
