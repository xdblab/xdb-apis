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

from xcherryapi.models.state_movement import StateMovement

class TestStateMovement(unittest.TestCase):
    """StateMovement unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> StateMovement:
        """Test StateMovement
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `StateMovement`
        """
        model = StateMovement()
        if include_optional:
            return StateMovement(
                state_id = '',
                state_input = xcherryapi.models.encoded_object.EncodedObject(
                    encoding = '', 
                    data = '', ),
                state_config = xcherryapi.models.async_state_config.AsyncStateConfig(
                    skip_wait_until = True, 
                    wait_until_api_timeout_seconds = 56, 
                    execute_api_timeout_seconds = 56, 
                    wait_until_api_retry_policy = xcherryapi.models.retry_policy.RetryPolicy(
                        initial_interval_seconds = 56, 
                        backoff_coefficient = 1.337, 
                        maximum_interval_seconds = 56, 
                        maximum_attempts = 56, 
                        maximum_attempts_duration_seconds = 56, ), 
                    execute_api_retry_policy = xcherryapi.models.retry_policy.RetryPolicy(
                        initial_interval_seconds = 56, 
                        backoff_coefficient = 1.337, 
                        maximum_interval_seconds = 56, 
                        maximum_attempts = 56, 
                        maximum_attempts_duration_seconds = 56, ), 
                    state_failure_recovery_options = xcherryapi.models.state_failure_recovery_options.StateFailureRecoveryOptions(
                        policy = 'FAIL_PROCESS_ON_STATE_FAILURE', 
                        state_failure_proceed_state_id = '', 
                        state_failure_proceed_state_config = xcherryapi.models.async_state_config.AsyncStateConfig(
                            skip_wait_until = True, 
                            wait_until_api_timeout_seconds = 56, 
                            execute_api_timeout_seconds = 56, 
                            app_database_read_request = xcherryapi.models.app_database_read_request.AppDatabaseReadRequest(
                                tables = [
                                    xcherryapi.models.app_database_table_read_request.AppDatabaseTableReadRequest(
                                        table_name = '', 
                                        lock_type = 'NO_LOCKING', 
                                        columns = [
                                            ''
                                            ], )
                                    ], ), 
                            load_local_attributes_request = xcherryapi.models.load_local_attributes_request.LoadLocalAttributesRequest(
                                keys_to_load_no_lock = [
                                    ''
                                    ], 
                                keys_to_load_with_lock = [
                                    ''
                                    ], ), ), ), 
                    app_database_read_request = xcherryapi.models.app_database_read_request.AppDatabaseReadRequest(), 
                    load_local_attributes_request = xcherryapi.models.load_local_attributes_request.LoadLocalAttributesRequest(), )
            )
        else:
            return StateMovement(
                state_id = '',
        )
        """

    def testStateMovement(self):
        """Test StateMovement"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
