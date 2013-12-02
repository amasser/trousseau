package trousseau

const TROUSSEAU_VERSION = "0.1.4"

const STORE_FILENAME = ".trousseau"

const (
	CONFIG_KEY_RECIPIENTS = "recipients"
	CONFIG_KEY_PASSWORD   = "password"
)

const (
	ENV_PASSPHRASE_KEY      = "TROUSSEAU_PASSPHRASE"
	ENV_KEYRING_SERVICE_KEY = "TROUSSEAU_KEYRING_SERVICE"
	ENV_KEYRING_USER_KEY    = "USER"
	ENV_MASTER_GPG_ID_KEY   = "TROUSSEAU_MASTER_GPG_ID"
	ENV_SSH_PRIVATE_KEY     = "TROUSSEAU_PRIVATE_KEY"
)
