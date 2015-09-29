package randomorg

const (
    URL = "https://api.random.org/json-rpc/1/invoke"
    ContentType = "application/json-rpc"
    JSONRPC = "2.0"
)

// Methods

const (
    generateIntegersMethod = "generateIntegers"
    generateDecimalFractionsMethod = "generateDecimalFractions"
    generateGaussiansMethod = "generateGaussians"
    generateStringsMethod = "generateStrings"
    generateUUIDsMethod = "generateUUIDs"
    generateBlobsMethod = "generateBlobs"
    getUsageMethod = "getUsage"
)