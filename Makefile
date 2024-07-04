.PHONY: v1 api

v1:
	kiota generate -l go -d https://api.infusionsoft.com/info-service/crm/docs/rest/V1 -c KeapClientV1 -n github.com/0x090909/keap_api -o ./
api:
	kiota generate -l go -d https://api.infusionsoft.com/info-service/crm/docs/rest/V2 -c KeapClientV2 -n github.com/0x090909/keap_api -o ./