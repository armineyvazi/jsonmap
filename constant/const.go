package constant

const GptToken = "sk-proj-q-fWioWjw-f4i9L7wGMX0pAaWAhcV5UBBR1Jc7vNWMu8CemDl" +
	"Q80IEw0up4-AY1P6KXSMCM71RT3BlbkFJXxgEIfwcqmn5JCU5QlmAlA" +
	"9Hfi0P7z8v93KzpfDhKfvOoywmHLQW3NurYWxWxSbKKznT8P2TkA"

const Regex = `(?s)\[.*\]`

const PromptToJSON = "\"Convert the provided data into JSON according to the specified struct. If certain fields (e.g., `ramType`) are missing, use common knowledge or search online to fill them in. For any fields related to the battery or power source (e.g., `BatteryStatus`, `Battery`, `BatteryLife`, `BatteryIncluded`, `BatteryCondition`, `PowerSource`, `ChargerStatus`, or similar terms), apply the following rules:\n\n- If terms like 'missing battery', 'battery removed', or 'damaged battery' are found, set the corresponding `BatteryStatus` to 'no'.\n- If no battery-related terms are mentioned, set `BatteryStatus` to 'yes' by default.\n- If terms like 'missing charger' are present, set `ChargerStatus` to 'no'.\n- If no mention of the charger is made, set `ChargerStatus` to 'yes' by default.\""
const BindToStruct = "type LaptopDetail struct {\n\tBrand           string `json:\"brand\"`\n\tModel           string `json:\"model\"`\n\tProcessor       string `json:\"processor\"`\n\tRamCapacity     string `json:\"ram_capacity\"`\n\tRamType         string `json:\"ram_type\"`\n\tStorageCapacity string `json:\"storage_capacity\"`\n\tBatteryStatus   string `json:\"battery_status\"`\n}"

const DirPath = "./config.yml"
