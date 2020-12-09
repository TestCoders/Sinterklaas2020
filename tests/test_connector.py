from unittest.mock import patch
from connector import Connector
import unittest


class TestBase(unittest.TestCase):
    def __init__(self, methodName='runTest', param=None):
        super(TestBase, self).__init__(methodName)
        self.vendor = param
    
    @staticmethod
    def parametrize(testcase_class, param=None):
        testloader = unittest.TestLoader()
        testnames = testloader.getTestCaseNames(testcase_class)
        suite = unittest.TestSuite()
        for name in testnames:
            suite.addTest(testcase_class(name, param=param))
        return suite


class ConnectorTest(TestBase):
    def setUp(self):
        self.connector = Connector(self.vendor, f'https://{self.vendor}.com')
        self.response_get = {"product":
                                {
                                "Id":5,
                                "Price":4.89,
                                "Name":"Playdebiel"
                                }
                            }
        self.payload = {"action":"order"}

    def tearDown(self):
        print(f'Test for {self.vendor} is done!')

    @patch('connector.requests.get')
    def test_get(self, mocked_get):
        mocked_get.return_value.status_code = 200
        mocked_get.return_value.ok = True
        mocked_get.return_value.body = self.response_get
                
        response = self.connector.get(5)

        mocked_get.assert_called_with(f'https://{self.vendor}.com/cadeaus/5')
        self.assertEqual(response.ok, True)
        self.assertEqual(response.status_code, 200)
        self.assertIsNotNone(response.body)
        self.assertEqual(response.body, self.response_get)

    @patch('connector.requests.post')
    def test_post(self, mocked_post):
        mocked_post.return_value.status_code = 200
        mocked_post.return_value.ok = True
                
        response = self.connector.post(5, self.payload)

        mocked_post.assert_called_with(f'https://{self.vendor}.com/cadeaus/5', 
                                        data=self.payload)
        self.assertEqual(response.status_code, 200)


if __name__ == '__main__':
    suite = unittest.TestSuite()
    for vendor in ['bollie', 'aliblabla', 'coolbere']:
        suite.addTest(TestBase.parametrize(ConnectorTest, param=vendor))
    unittest.TextTestRunner(verbosity=2).run(suite)
