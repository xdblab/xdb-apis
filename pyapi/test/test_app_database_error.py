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

from xcherryapi.models.app_database_error import AppDatabaseError

class TestAppDatabaseError(unittest.TestCase):
    """AppDatabaseError unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AppDatabaseError:
        """Test AppDatabaseError
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AppDatabaseError`
        """
        model = AppDatabaseError()
        if include_optional:
            return AppDatabaseError(
                app_db_error_type = 'UNCATEGORIZED_ERROR',
                app_db_error_code = '',
                app_db_error_message = '',
                app_db_error_table_name = ''
            )
        else:
            return AppDatabaseError(
                app_db_error_type = 'UNCATEGORIZED_ERROR',
                app_db_error_code = '',
        )
        """

    def testAppDatabaseError(self):
        """Test AppDatabaseError"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()