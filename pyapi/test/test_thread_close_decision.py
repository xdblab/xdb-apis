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

from xcherryapi.models.thread_close_decision import ThreadCloseDecision

class TestThreadCloseDecision(unittest.TestCase):
    """ThreadCloseDecision unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ThreadCloseDecision:
        """Test ThreadCloseDecision
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ThreadCloseDecision`
        """
        model = ThreadCloseDecision()
        if include_optional:
            return ThreadCloseDecision(
                close_type = 'FORCE_COMPLETE_PROCESS',
                close_input = xcherryapi.models.encoded_object.EncodedObject(
                    encoding = '', 
                    data = '', )
            )
        else:
            return ThreadCloseDecision(
                close_type = 'FORCE_COMPLETE_PROCESS',
        )
        """

    def testThreadCloseDecision(self):
        """Test ThreadCloseDecision"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
