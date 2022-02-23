.PHONY: production logs

A0= $(subst .,,$(suffix $@))
A1=$(basename $@)

ENVS = production
DOMAINS = user

npmci:
	npm ci

$(foreach x,$(DOMAINS),$(addsuffix .$x,$(ENVS))): npmci
	@mkdir -p $(LOG_DIR)
	@echo $(A0) start
	@$(MAKE) -j2 -C $(A0) $(A1) > $(LOG_DIR)/deploy-$(A0).log 2>&1 
	@echo $(A0) ok

production: $(addprefix production.,$(DOMAINS))
	@echo "deployment finish"

logs: $(addprefix logs.,$(DOMAINS))
