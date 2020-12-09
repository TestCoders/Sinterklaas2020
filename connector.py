import requests


class Connector:
    
    def __init__(self, name, url):
        self.name = name
        self.url = url
    
    def get(self, id):
        r = requests.get(f'{self.url}/cadeaus/{id}')
        if r.ok:
            return r
        else:
            return 'Bad response!'

    def post(self, id, payload):
        r = requests.post(f'{self.url}/cadeaus/{id}', data=payload)
        if r.ok:
            return r
        else:
            return 'Bad response!'
