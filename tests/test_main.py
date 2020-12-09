from logic import main, request_connector
import logic
from connector import Connector
from unittest.mock import patch
import unittest


class IntegrationTest(unittest.TestCase):

    def setUp(self):
        self.response_body = {"product":
                            {
                            "Id":id,
                            "Price": 12.34,
                            "Name":"Playdebiel"
                            }
                        }

    def tearDown(self):
        pass

    @patch('connector.requests.post')
    def test_request_connector_post(self, mocked_post):
        mocked_post.return_value.status_code = 200
        mocked_post.return_value.ok = True

        response = request_connector('coolbere', 5, req='post')

        mocked_post.assert_called_with('https://coolbere.com/cadeaus/5', 
                                        data={'action': 'order'})
        self.assertEqual(response, 'Buy order has succesfully been placed at coolbere')
    
    @patch('connector.requests.get')
    def test_request_connector_get(self, mocked_get):
        mocked_get.return_value.status_code = 200
        mocked_get.return_value.ok = True
        mocked_get.return_value.body = self.response_body

        response = request_connector('coolbere', 5)

        mocked_get.assert_called_with('https://coolbere.com/cadeaus/5')
        self.assertEqual(response, 12.34)

    @patch('logic.request_connector') # mocking this method to assert that it's called (with correct parameters)
    @patch('logic.get_prices')
    def test_main(self, mocked_prices, mocked_request_connector):
        mocked_prices.return_value = {'bollie':12.34, 'coolbere':2.34, 'aliblabla':34.56}
        
        results = main(['bollie', 'coolbere', 'aliblabla'], 5)

        # print(mocked_request_connector.mock_calls)
        mocked_request_connector.assert_called_with(5, 'coolbere', req='post')

if __name__ == '__main__':
    unittest.main()