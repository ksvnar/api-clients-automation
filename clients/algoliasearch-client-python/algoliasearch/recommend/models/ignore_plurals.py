# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import dumps, loads
from typing import Dict, List, Optional, Self, Union

from pydantic import BaseModel, Field, StrictBool, ValidationError, model_serializer

from algoliasearch.recommend.models.supported_language import SupportedLanguage


class IgnorePlurals(BaseModel):
    """
    Treat singular, plurals, and other forms of declensions as equivalent. You should only use this feature for the languages used in your index.
    """

    oneof_schema_1_validator: Optional[List[SupportedLanguage]] = Field(
        default=None,
        description="ISO code for languages for which this feature should be active. This overrides languages you set with `queryLanguages`. ",
    )
    oneof_schema_2_validator: Optional[StrictBool] = Field(
        default=False,
        description="If true, `ignorePlurals` is active for all languages included in `queryLanguages`, or for all supported languages, if `queryLanguges` is empty. If false, singulars, plurals, and other declensions won't be considered equivalent. ",
    )
    actual_instance: Optional[Union[List[SupportedLanguage], bool]] = None

    def __init__(self, *args, **kwargs) -> None:
        if args:
            if len(args) > 1:
                raise ValueError(
                    "If a position argument is used, only 1 is allowed to set `actual_instance`"
                )
            if kwargs:
                raise ValueError(
                    "If a position argument is used, keyword arguments cannot be used."
                )
            super().__init__(actual_instance=args[0])
        else:
            super().__init__(**kwargs)

    @model_serializer
    def unwrap_actual_instance(self) -> Optional[Union[List[SupportedLanguage], bool]]:
        """
        Unwraps the `actual_instance` when calling the `to_json` method.
        """
        return self.actual_instance

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        return cls.from_json(dumps(obj))

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Returns the object represented by the json string"""
        instance = cls.model_construct()
        error_messages = []

        try:
            instance.oneof_schema_1_validator = loads(json_str)
            instance.actual_instance = instance.oneof_schema_1_validator

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.oneof_schema_2_validator = loads(json_str)
            instance.actual_instance = instance.oneof_schema_2_validator

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))

        raise ValueError(
            "No match found when deserializing the JSON string into IgnorePlurals with oneOf schemas: List[SupportedLanguage], bool. Details: "
            + ", ".join(error_messages)
        )

    def to_json(self) -> str:
        """Returns the JSON representation of the actual instance"""
        if self.actual_instance is None:
            return "null"

        to_json = getattr(self.actual_instance, "to_json", None)
        if callable(to_json):
            return self.actual_instance.to_json()
        else:
            return dumps(self.actual_instance)

    def to_dict(self) -> Dict:
        """Returns the dict representation of the actual instance"""
        if self.actual_instance is None:
            return None

        to_dict = getattr(self.actual_instance, "to_dict", None)
        if callable(to_dict):
            return self.actual_instance.to_dict()
        else:
            return self.actual_instance
