# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.3
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest
import datetime

from xcherryapi.models.load_local_attributes_response import LoadLocalAttributesResponse

class TestLoadLocalAttributesResponse(unittest.TestCase):
    """LoadLocalAttributesResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> LoadLocalAttributesResponse:
        """Test LoadLocalAttributesResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `LoadLocalAttributesResponse`
        """
        model = LoadLocalAttributesResponse()
        if include_optional:
            return LoadLocalAttributesResponse(
                attributes = [
                    xcherryapi.models.key_value.KeyValue(
                        key = '', 
                        value = xcherryapi.models.encoded_object.EncodedObject(
                            encoding = '', 
                            data = '', ), )
                    ]
            )
        else:
            return LoadLocalAttributesResponse(
        )
        """

    def testLoadLocalAttributesResponse(self):
        """Test LoadLocalAttributesResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()