import time
from insights.config import Config
from insights import client

line = {
    "id": "ba939b7d-3dcf-4e69-b837-20e5d930bebd",
    "tenantId": "d517e0ae-2519-4969-8933-ec191b7faa59",
    "tenantName": "fppunjab",
    "eventTime": "2024-07-30 05:00:05",
    "userID": "ed535f08-263a-4a9f-a000-28d5bd65ebe9",
    "userFirstName": "admin",
    "userLastName": "",
    "userName": "admin@vishalp.com",
    "renderingAction": 2,
    "url": "https://www.google.com",
    "host": "www.google.com",
    "urlCategories": "Search Engines and Portals",
    "malicious": False,
    "rbiPolicyID": "00000000-0000-0000-0000-000000000000",
    "rbiPolicyName": "",
    "applianceId": "6d963e85-e4be-479c-bcc9-5e4d2f21145b",
    "nodeId": "6b62e3df-830a-481c-a822-d2c679d16e68",
    "title": "Google",
    "tabId": "00000000-0000-0000-0000-000000000000",
    "deviceId": "5bab815f-8977-46c2-8539-a7080a94d519",
    "sessionId": "5e327e8e-a52f-4870-9a96-3a21d2bb93a7",
    "displayName": "",
    "ipAddress": "192.168.122.34",
    "clientType": "viewer",
    "clientVersion": "6.0.0",
    "os": "Windows",
    "osVersion": "NA",
    "browserType": "Edge",
    "threatScore": 10,
    "tags": ["test"],
    "country": ["india"],
    "thumbnail": ["test"]
}

lines = []
for i in range(0,1):
    lines.append(line)

config = Config(
    insights_host="portal-new.insights.dev.forcepointone.com",
    platform_host="portal.dev.forcepointone.com",
    client_id="01ec7292-f7fb-4f6f-94eb-89f54a5d7184",
    client_secret="snJQuiCeiVXIPVZhEBgg_ojAXr3zhcbLxjEaYi1HyPE=",
    # proxy_host="54.242.73.41",
    # proxy_port=3128,
    # proxy_scheme="http",
    # proxy_username="morpheus",
    # proxy_password="squid123"
)

ingestion_client = client('Data',config)
while True:
    result = ingestion_client.ingest("RBI" , "SiteVisitArray", lines,send_as_json=True)
    if result.request_delivered == True:
        print("correlationID " +  result.correlation_id , result.message , result.error)

    else:
        #in this case you have to buffer data for retrying later
        print("Error: " +  result.error +", need to buffer the data" ,  result.correlation_id , result.message  )

    #time.sleep(5)