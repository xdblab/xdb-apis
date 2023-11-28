# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.3
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import json
import pprint
import re  # noqa: F401
from enum import Enum



try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class CommandWaitingType(str, Enum):
    """
    CommandWaitingType
    """

    """
    allowed enum values
    """
    EMPTYCOMMAND = 'EmptyCommand'
    ANYOFCOMPLETION = 'AnyOfCompletion'
    ALLOFCOMPLETION = 'AllOfCompletion'

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of CommandWaitingType from a JSON string"""
        return cls(json.loads(json_str))


