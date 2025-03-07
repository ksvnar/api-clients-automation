# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import dumps
from typing import Dict, Optional, Self, Union

from pydantic import BaseModel, ValidationError, model_serializer

from algoliasearch.ingestion.models.on_demand_trigger import OnDemandTrigger
from algoliasearch.ingestion.models.schedule_trigger import ScheduleTrigger
from algoliasearch.ingestion.models.streaming_trigger import StreamingTrigger
from algoliasearch.ingestion.models.subscription_trigger import SubscriptionTrigger


class Trigger(BaseModel):
    """
    Trigger that runs the task.
    """

    oneof_schema_1_validator: Optional[OnDemandTrigger] = None
    oneof_schema_2_validator: Optional[ScheduleTrigger] = None
    oneof_schema_3_validator: Optional[SubscriptionTrigger] = None
    oneof_schema_4_validator: Optional[StreamingTrigger] = None
    actual_instance: Optional[
        Union[OnDemandTrigger, ScheduleTrigger, StreamingTrigger, SubscriptionTrigger]
    ] = None

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
    def unwrap_actual_instance(
        self,
    ) -> Optional[
        Union[OnDemandTrigger, ScheduleTrigger, StreamingTrigger, SubscriptionTrigger]
    ]:
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
            instance.actual_instance = OnDemandTrigger.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.actual_instance = ScheduleTrigger.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.actual_instance = SubscriptionTrigger.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        try:
            instance.actual_instance = StreamingTrigger.from_json(json_str)

            return instance
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))

        raise ValueError(
            "No match found when deserializing the JSON string into Trigger with oneOf schemas: OnDemandTrigger, ScheduleTrigger, StreamingTrigger, SubscriptionTrigger. Details: "
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
