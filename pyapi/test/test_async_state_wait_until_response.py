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

from xcherryapi.models.async_state_wait_until_response import AsyncStateWaitUntilResponse

class TestAsyncStateWaitUntilResponse(unittest.TestCase):
    """AsyncStateWaitUntilResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AsyncStateWaitUntilResponse:
        """Test AsyncStateWaitUntilResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AsyncStateWaitUntilResponse`
        """
        model = AsyncStateWaitUntilResponse()
        if include_optional:
            return AsyncStateWaitUntilResponse(
                command_request = xcherryapi.models.command_request.CommandRequest(
                    waiting_type = 'EmptyCommand', 
                    timer_commands = [
                        xcherryapi.models.timer_command.TimerCommand(
                            delay_in_seconds = 56, )
                        ], 
                    local_queue_commands = [
                        xcherryapi.models.local_queue_command.LocalQueueCommand(
                            queue_name = '', 
                            count = 56, )
                        ], ),
                publish_to_local_queue = [
                    xcherryapi.models.local_queue_message.LocalQueueMessage(
                        queue_name = '', 
                        dedup_id = '', 
                        payload = xcherryapi.models.encoded_object.EncodedObject(
                            encoding = '', 
                            data = '', ), )
                    ]
            )
        else:
            return AsyncStateWaitUntilResponse(
                command_request = xcherryapi.models.command_request.CommandRequest(
                    waiting_type = 'EmptyCommand', 
                    timer_commands = [
                        xcherryapi.models.timer_command.TimerCommand(
                            delay_in_seconds = 56, )
                        ], 
                    local_queue_commands = [
                        xcherryapi.models.local_queue_command.LocalQueueCommand(
                            queue_name = '', 
                            count = 56, )
                        ], ),
        )
        """

    def testAsyncStateWaitUntilResponse(self):
        """Test AsyncStateWaitUntilResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()