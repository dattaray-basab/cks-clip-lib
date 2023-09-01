package globals

// const LOG_DIR_PATH = "logs"
// const LOG_FILE_NAME = "/app.log"

const RecipePathKey = "CKS_ENV_RECIPE_PATH"
const RecipeDirectory = "__DEV_3"
const TEMPLATES_DIRNAME = "_templates"

const RECIPE = "recipe"
const PHASE = "phase"
const ALTER = "alter"
const SPECIAL_DIR_PREFIX_ = "__"
const RECIPE_ROOT_DIR_ = SPECIAL_DIR_PREFIX_ + RECIPE
const ALTER_ROOT_DIR_ = SPECIAL_DIR_PREFIX_ + ALTER

const BLUEPRINTS = "BLUEPRINTS"
const MISC = "MISC"
const PHASES = "PHASES"
const CODE = "CODE"
const TOKENS = "TOKENS"
const DEPENDS_ON = "DEPENDS_ON"
const CODE_BLOCK = "CODE_BLOCK"
const QUERY = "QUERY"

const BLUEPRINTS_DIRNAME = SPECIAL_DIR_PREFIX_ + BLUEPRINTS
const MISC_DIRNAME = SPECIAL_DIR_PREFIX_ + MISC
const PHASES_DIRNAME = SPECIAL_DIR_PREFIX_ + PHASES
const CODE_BLOCK_ROOT = SPECIAL_DIR_PREFIX_ + CODE
const TOKENS_DIRNAME = SPECIAL_DIR_PREFIX_ + TOKENS
const DEPENDS_ON_DIRNAME = SPECIAL_DIR_PREFIX_ + DEPENDS_ON
const CODE_BLOCK_ = SPECIAL_DIR_PREFIX_ + CODE_BLOCK
const QUERY_DIRNAME = SPECIAL_DIR_PREFIX_ + QUERY

const RUN_PY = "run.py"
const JSON_EXT = ".json"

const DIRECTIVES_lc = "directives"
const DIRECTIVES_JSON = DIRECTIVES_lc + JSON_EXT

const RECIPE_CONFIG_ = "__RECIPE_CONFIG.json"

const KEY_TARGET = "{{target}}"
const KEY_PHASE_NAME = "{{phase-name}}"
const KEY_CODE_BLOCK_NAME = "{{code-block-name}}"
const KEY_MOVE_ITEMS = "{{move-items}}"

const KEY_RECIPE_PATH = "{{recipe-path}}"
const KEY_LAST_PHASE = "{{last-phase}}"
const KEY_DEPENDS_ON_PHASE = "{{depends-on-phase}}"
const KEY_ALTER_NAME = "{{alter-name}}"
const KEY_ALTER_DIR_PATH = "{{alter-dir-path}}"

const KEY_FORCE = "{{force}}"

const KEY_ALTER_PATH = "{{alter-path}}"
const KEY_BLUEPRINTS_PATH = "{{blueprints-path}}"
const KEY_CODE_BLOCK_ROOT_PATH = "{{code-block-root-path}}"
const KEY_CODE_BLOCK_PATH = "{{code-block-path}}"
const KEY_PHASES_PATH = "{{phases-path}}"
const KEY_CODE_BLOCK_NAME_WITH_QUOTES = "{{code-block-name-with-quotes}}"
const KEY_ALTER_PATH_WITH_QUOTES = "{{alter-path-with-quotes}}"
const KEY_DEPENDS_ON_PHASE_WITH_QUOTES = "{{depends-on-phase-with-quotes}}"

const KEY_STORE_DIR_PATH = "{{store-dir-path}}"
const KEY_CONTROL_JSON_PATH = "{{control-json-path}}"

const CONTROL_JSON_FILE = "control.json"
const STORE_lc = "store"
const STORE_DIRNAME = SPECIAL_DIR_PREFIX_ + STORE_lc
