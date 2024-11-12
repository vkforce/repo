import requests
import yaml
from APILocalSDK.oidcToken import GenerateToken
import logging


class BaseCollectionHandler:
    def __init__(self, env):
        self.env = env
        self.config = self.load_env_from_yaml('C:/Users/vedhas.kharche/Music/Codes/APILocalSDK/config.yaml')
        self._base_url = self.config.get(env.upper()).get('INSIGHTS_HOST', '')
        token_object = GenerateToken(env=env)
        self.oidc_token = token_object.handler()

    def load_env_from_yaml(self, yaml_file):
        """Loads environment variables from a YAML file.

      Args:
        yaml_file: Path to the YAML file.

      Returns:
        A dictionary of environment variables.
      """
        with open(yaml_file, 'r') as f:
            config = yaml.safe_load(f)
            return config

    @property
    def base_url(self):
        return self._base_url

    @base_url.setter
    def base_url(self, base_url):
        self._base_url = base_url


class GetCollectionHandler(BaseCollectionHandler):
    def __init__(self, env):
        super().__init__(env)
        self._url = None
        self._collectionName = None
        self.ppcode = None

    @property
    def url(self):
        return self._url

    @url.setter
    def url(self, url):
        self._url = url + '/api/collections/' + self.name + '/' + self.ppcode

    @property
    def name(self):
        return self._collectionName

    @name.setter
    def name(self, name):
        self._collectionName = name

    @property
    def ppcode(self):
        return self.ppcode

    @ppcode.setter
    def ppcode(self, ppcode):
        self.ppcode = ppcode

    def handler(self, collection_name, ppcode):
        self.name = collection_name
        self.ppcode = ppcode
        self.url = self.base_url
        response = requests.get(self.url, headers={"Authorization": f"Bearer {self.oidc_token}"})
        if response.status_code == 200:
            logging.info("Successfully Get collection!")
            return response.json()
        else:
            logging.error("Failed to Get a collection!")
            return response.json()


class GetAllCollectionHandler(BaseCollectionHandler):
    def __init__(self, env):
        super().__init__(env)
        self._url = None

    @property
    def url(self):
        return self._url

    @url.setter
    def url(self, url):
        self._url = url + '/api/collections'

    def handler(self):
        self.url = self.base_url
        response = requests.get(self.url, headers={"Authorization": f"Bearer {self.oidc_token}"})
        if response.status_code == 200:
            logging.info("Successfully Get Collections!")
            return response.json()
        else:
            logging.error("Failed to Get all Collections!")
            return response.json()


class CreateCollectionHandler(BaseCollectionHandler):
    def __init__(self, env):
        super().__init__(env)
        self.new_collection_name = None
        self.new_ppcode = None
        self._url = None

    @property
    def url(self):
        return self._url

    @url.setter
    def url(self, url):
        self._url = url + '/api/collections'

    def handler(self, new_collection_name, new_ppcode, payload):
        self.url = self.base_url
        payload['name'] = new_collection_name
        payload['metadata']['pointProductCode'] = new_ppcode
        response = requests.post(self.url, data=payload, headers={'Content-Type': 'application/json',
                                                                  'Accept': 'application/json',
                                                                  "Authorization": f"Bearer {self.oidc_token}"})
        if response.status_code == 200:
            logging.info("Successfully Created collection!")
            return response.json()
        else:
            logging.error("Failed to Create a New collection!")
            return response.json()
