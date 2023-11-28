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

from xcherryapi.models.process_start_config import ProcessStartConfig

class TestProcessStartConfig(unittest.TestCase):
    """ProcessStartConfig unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ProcessStartConfig:
        """Test ProcessStartConfig
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ProcessStartConfig`
        """
        model = ProcessStartConfig()
        if include_optional:
            return ProcessStartConfig(
                timeout_seconds = 56,
                id_reuse_policy = 'ALLOW_IF_PREVIOUS_EXIT_ABNORMALLY',
                app_database_config = xcherryapi.models.app_database_config.AppDatabaseConfig(
                    tables = [
                        xcherryapi.models.app_database_table_config.AppDatabaseTableConfig(
                            table_name = '', 
                            rows = [
                                xcherryapi.models.app_database_table_row_selector.AppDatabaseTableRowSelector(
                                    primary_key = [
                                        xcherryapi.models.app_database_column_value.AppDatabaseColumnValue(
                                            column = '', 
                                            query_value = '', )
                                        ], 
                                    initial_write = [
                                        xcherryapi.models.app_database_column_value.AppDatabaseColumnValue(
                                            column = '', 
                                            query_value = '', )
                                        ], 
                                    conflict_mode = 'RETURN_ERROR_ON_CONFLICT', )
                                ], )
                        ], ),
                local_attribute_config = xcherryapi.models.local_attribute_config.LocalAttributeConfig(
                    initial_write = [
                        xcherryapi.models.key_value.KeyValue(
                            key = '', 
                            value = xcherryapi.models.encoded_object.EncodedObject(
                                encoding = '', 
                                data = '', ), )
                        ], )
            )
        else:
            return ProcessStartConfig(
        )
        """

    def testProcessStartConfig(self):
        """Test ProcessStartConfig"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
