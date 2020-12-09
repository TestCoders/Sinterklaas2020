from connector import Connector
from unittest import mock


def request_connector(shop, id, req='get'):
    """This method get's executed if logic states that a shop is the cheapest."""
    connector = Connector(shop, f'https://{shop}.com')
    if req == 'get':
        response = connector.get(id)
        return response.body['product']['Price']
    elif req == 'post':
        connector.post(id, {"action":"order"})
        return f'Buy order has succesfully been placed at {shop}'
    else: 
        raise TypeError("keyword argument 'req' only accepts 'get' and 'post' as its value!")

def get_prices(providers, id):
    return {shop:request_connector(shop, id) for shop in providers}

def main(providers, id):
    """The main method evaluates connector results 
    followed by placing an order at the cheapest shop."""
    
    prices = get_prices(providers, id)
    
    request_connector(id, min(prices, key=prices.get), req='post')

if __name__ == '__main__':
    main(['bollie', 'coolbere', 'aliblabla'], 5)