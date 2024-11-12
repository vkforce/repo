import configparser
import requests
from urllib.parse import urlencode
import yaml


def load_env_from_yaml(yaml_file):
    """Loads environment variables from a YAML file.

  Args:
    yaml_file: Path to the YAML file.

  Returns:
    A dictionary of environment variables.
  """
    with open(yaml_file, 'r') as f:
        config = yaml.safe_load(f)
        return config


class GenerateToken:

    def __init__(self, env):
        self.env = env
        self._url = None
        self._client_secret = None
        self._client_id = None
        self._audience = None
        self._scope = None
        self._grant_type = None
        self.data = {}

    @property
    def url(self):
        return self._url

    @property
    def scope(self):
        return self._scope

    @property
    def grant_type(self):
        return self._grant_type

    @property
    def client_id(self):
        return self._client_id

    @property
    def client_secret(self):
        return self._client_secret

    @property
    def audience(self):
        return self._audience

    @audience.setter
    def audience(self, audience):
        self._audience = audience

    @client_id.setter
    def client_id(self, value):
        self._client_id = value

    @client_secret.setter
    def client_secret(self, value):
        self._client_secret = value

    @url.setter
    def url(self, value):
        self._url = value + "/oidc/token"

    @scope.setter
    def scope(self, value):
        self._scope = value

    @grant_type.setter
    def grant_type(self, value):
        self._grant_type = value

    def configure(self, env):
        config = load_env_from_yaml('C:/Users/vedhas.kharche/Music/Codes/APILocalSDK/config.yaml')
        self.grant_type = config.get('Global').get('GRANT_TYPE')
        self.scope = config.get('Global').get('SCOPE')
        self.audience = config.get('Global').get('AUDIENCE')
        self.client_id = config.get(env.upper()).get('CLIENT_ID', '')
        self.client_secret = config.get(env.upper()).get('CLIENT_SECRET', '')
        self.url = config.get(env.upper()).get('HOST', '')
        self.data = {
            'grant_type': self._grant_type,
            'client_id': self.client_id,
            'client_secret': self.client_secret,
            'audience': self.audience,
            'scope': self.scope
        }

    def handler(self):
        self.configure(self.env)
        payload = urlencode(self.data)
        response = requests.post(self.url, data=payload,
                                 headers={'Content-Type': 'application/x-www-form-urlencoded'})
        return response.json().get('access_token')
