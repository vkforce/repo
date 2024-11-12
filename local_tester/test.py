import http.client
import json

conn = http.client.HTTPSConnection("portal-new.insights.dev.forcepointone.com")
payload = json.dumps({
   "catalogueIds": [
      "9756b411-20f5-47e5-9a82-2b8b8c2d5fb6"
   ]
})
headers = {
   'User-Agent': 'Apidog/1.0.0 (https://apidog.com)',
   'Content-Type': 'application/json',
   'Authorization': 'Bearer fp_oidc_at_vsbkdaEZIVisv7ib_bo8DS_Cz78pbJoB2O69FruOacQ.BDvWWaM1yFEPEjEcUuBFJW_YwEhE2ANm5PRaLxG8R0s'
}
conn.request("POST", "/api/catalogue/export", payload, headers)
res = conn.getresponse()
data = res.read()
print(data.decode("utf-8"))