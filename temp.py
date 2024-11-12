from APILocalSDK.oidcToken import GenerateToken
from APILocalSDK.lib.exceptions import NewMethod


token = GenerateToken(env='dev')

print(token
      )

print(NewMethod())