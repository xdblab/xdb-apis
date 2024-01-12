# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.3
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import Any, ClassVar, Dict, List, Optional
from pydantic import BaseModel, StrictStr
from pydantic import Field
from xcherryapi.models.async_state_config import AsyncStateConfig
from xcherryapi.models.encoded_object import EncodedObject
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class StateMovement(BaseModel):
    """
    StateMovement
    """ # noqa: E501
    state_id: StrictStr = Field(alias="stateId")
    state_input: Optional[EncodedObject] = Field(default=None, alias="stateInput")
    state_config: Optional[AsyncStateConfig] = Field(default=None, alias="stateConfig")
    __properties: ClassVar[List[str]] = ["stateId", "stateInput", "stateConfig"]

    model_config = {
        "populate_by_name": True,
        "validate_assignment": True,
        "protected_namespaces": (),
    }


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of StateMovement from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={
            },
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of state_input
        if self.state_input:
            _dict['stateInput'] = self.state_input.to_dict()
        # override the default output from pydantic by calling `to_dict()` of state_config
        if self.state_config:
            _dict['stateConfig'] = self.state_config.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of StateMovement from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "stateId": obj.get("stateId"),
            "stateInput": EncodedObject.from_dict(obj.get("stateInput")) if obj.get("stateInput") is not None else None,
            "stateConfig": AsyncStateConfig.from_dict(obj.get("stateConfig")) if obj.get("stateConfig") is not None else None
        })
        return _obj


