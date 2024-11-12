import logging
from APILocalSDK.oidcToken import GenerateToken
import yaml
import requests


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


class Ingestion:
    def __init__(self, env, payload):
        self.env = env
        self._url = None
        self.payload = payload

    @property
    def url(self):
        return self._url

    @url.setter
    def url(self, url):
        self._url = url + '/api/data'

    def handler(self):
        token_object = GenerateToken(env=self.env)
        oidc_token = token_object.handler()
        config = load_env_from_yaml('C:/Users/vedhas.kharche/Music/Codes/APILocalSDK/config.yaml')
        self.url = config.get(self.env.upper()).get('INSIGHTS_HOST', '')
        response = requests.post(url=self.url, data=self.payload,
                                 headers={'Content-Type': 'application/json', "Authorization": f"Bearer {oidc_token}"})
        if response.status_code == 202:
            logging.info("Successfully Ingested!")
            return response.json()
        else:
            logging.error("Failed to Ingest!")
            return response.json()
